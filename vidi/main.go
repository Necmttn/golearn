package main

import (
	"encoding/xml"
	"fmt"
	"github.com/Necmttn/golearn/vidi/parser"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	return
}

func crawl(url string, ch chan string, chFinished chan bool) {
	resp, err := http.Get(url)

	// notify when it's finish with this channel
	defer func() {
		chFinished <- true
	}()

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return
	}

	b := resp.Body
	defer b.Close() // close Body when the function returns
	z := html.NewTokenizer(b)
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document we're done.
			return
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "a"

			// filter anchors.
			if !isAnchor {
				continue
			}
			ok, url := getHref(t)
			if !ok {
				continue
			}
			// Make sure Url starts with http**

			hasProto := strings.Index(url, "/watch") == 0
			if hasProto {
				// ch <- getSub(url)
				getSub(url)
			}
		}
	}
}

func getVideoId(path string) string {
	u, err := url.Parse(path)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	return q.Get("v")
}

func getSubUrl(id string) string {
	s := "https://www.youtube.com/api/timedtext?v=k__Et4Y798Q&key=yttt1&asr_langs=it,es,ru,fr,ko,de,pt,nl,ja,en&signature=2910070525A69A926C9C8328EF89B1CB3C6FFF5B.4197307189657C64E6512021DFEFF252A48DBBFA&sparams=asr_langs,caps,v,expire&hl=tr_TR&caps=asr&lang=zh-TW&fmt=srv1&tlang=zh-Hans"
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("v", id)
	u.RawQuery = q.Encode()
	return u.String()
}

type Transcript struct {
	XMLName xml.Name `xml:"transcript"`
	Text    []text   `xml:"text"`
}

type text struct {
	XMLName  xml.Name `xml:"text"`
	Start    string   `xml:"start,attr"`
	Dur      string   `xml:"dur,attr"`
	Sentence string   `xml:",innerxml"`
}

func getSub(path string) {
	id := getVideoId(path)
	url := getSubUrl(id)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// f, err := os.Create("./subs/" + id + ".xml")

	// defer f.Close()
	// resp.Write(f)
	b := resp.Body
	defer b.Close()
	v := Transcript{}
	xml.NewDecoder(b).Decode(&v)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Printf("result")
	fmt.Println(v)
	// return transcript
	return
}

func main() {
	foundUrls := make(map[string]bool)
	seedUrls := os.Args[1:]

	// Channels
	chUrls := make(chan string)
	chFinished := make(chan bool)

	// Kick off the crawl process (concurently)

	for _, url := range seedUrls {
		go crawl(url, chUrls, chFinished)
	}

	// Subscribe to both channels
	for c := 0; c < len(seedUrls); {
		select {
		case url := <-chUrls:
			foundUrls[url] = true
		case <-chFinished:
			c++
		}
	}
	// we're done.

	fmt.Println("\nFound", len(foundUrls), "unique urls:\n")
	for url, _ := range foundUrls {
		fmt.Println(" - " + url)
	}

	close(chUrls)
	// Subscribe for subtitles
}
