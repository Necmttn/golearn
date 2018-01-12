package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type transcript struct {
	XMLName xml.Name `xml:"transcript"`
	text    []text   `xml:"text"`
}

type text struct {
	XMLName xml.Name `xml:"text"`
	start   string   `xml:"start, attr"`
	dur     string   `xml:"dur, attr"`
}

func main() {
	file, err := os.open("../test.xml")
	if err != nil {
		fmt.printf("error: %v", err)
		return
	}
	defer file.close()
	data, err := ioutil.readall(file)

	if err != nil {
		fmt.printf("error: %v", err)
		return
	}
	v := recurlyservers{}

	err = xml.unmarshal(data, &v)

	if err != nil {
		fmt.printf("error: %v", err)
		return
	}

	fmt.println(v)
}
