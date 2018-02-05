package main

import (
	"fmt"
	"strconv"
	s "strings"
)

var p = fmt.Println

func main() {
	reverse(-123)
}

func reverse(x int) int {
	t := strconv.Itoa(x)
	a := s.Split(t, "")
	isMinus := false
	if a[0] == "-" {
		a = append(a[:0], a[1:]...)
		isMinus = true
	} else {
	}
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	if isMinus {
		p("MINUSSS")
		a = append([]string{"-"}, a...)
	}
	str := s.Join(a, "")
	p(str)
	answer := strconv.Atoi
	p(answer)
	return answer
}
