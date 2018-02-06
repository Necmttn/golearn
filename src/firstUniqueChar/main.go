package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {
	faster("lovemetender")
}

func firstUniqChar(s string) int {
	i, length := 0, len(s)
	chars := strings.Split(s, "")
	for i < length {
		if strings.Count(s, chars[i]) == 1 {
			return i
			break
		}
		i++
	}
	return -1
}

func faster(s string) int {
	var charCnt [256]int
	for _, c := range s {
		p("c", c)
		charCnt[int(c)]++
		p("charCnt", charCnt)
	}
	for i, c := range s {
		if charCnt[int(c)] == 1 {
			return i
		}
	}
	return -1
}
