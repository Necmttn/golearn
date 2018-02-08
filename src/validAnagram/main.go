package main

import (
	"fmt"
)

var p = fmt.Println

func main() {
	result := isAnagram("at", "act")
	p(result)
	fastSimple := fast("ata", "act")
	p(fastSimple)

	ascii := isAnagram("cat$%^", "cat$%^")
	p(ascii)
	asciiWithFast := fast("cat$%^", "cat$%^")
	p(asciiWithFast)

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

// returns about 24ms

// this is smart answer because in the question it has note that.
// > You may assume the string contains only lowercase alphabets.

func fast(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// biggest lowercase ascii char num is 'z' with value of 122
	// lowest lowercase ascii char num is 'a' with value of	97
	// 122 - 97 = 25 + 1 for index 0;
	// so we don't need a map for dict for only lovercase chars.
	// it's a good solution for asci chars
	// if you run this code with ascii it will panic.
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		m[t[i]-'a']--
		if m[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

// returns about 3ms
