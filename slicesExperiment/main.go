package main

import (
	"fmt"
)

var p = fmt.Println

func main() {
	result := isAnagram("at", "act")
	p(result)
}

func isAnagram(s string, t string) bool {
	// Quick check for non-valid question.
	if len(t) > len(s) {
		return false
	}

	hash := make(map[rune]int)

	// built with first one.
	for _, c := range t {
		if _, ok := hash[c]; ok {
			hash[c]++
		} else {
			hash[c] = 1
		}
	}
	// remove on the second
	for _, c := range s {
		if _, ok := hash[c]; ok && hash[c] > 0 {
			hash[c]--
		} else {
			return false
		}
	}
	return true
}
