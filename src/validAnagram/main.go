package main

import (
	"fmt"
)

var p = fmt.Println

func main (
	isAnagram("anagram", "")
)


func isAnagram(s string, t string) bool {
	var charCnt [256]int
	for _, c := range s {
		p("c", c)
		charCnt[int(c)]++
	}
	for _, c := range t {
		p("t", c)
		charCnt[int(c)]--
	}
	for i, c :+ range s {
		if (charCnt[int(c)]) >= 0) {
			return true
		} else {
			return false
		}
	}
}
