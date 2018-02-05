package main

import (
	"fmt"
	"math"
	"strconv"
	s "strings"
)

var p = fmt.Println

func main() {
	reverseFast(-123)
}

/**
* this one is about 20ms which is way slow compare to next one
 */
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
	answer, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}
	return int(answer)
}

// Faster
func reverseFast(x int) int {
	var ret int
	var sign = x < 0 //  let's check the number given is negative or not.

	if sign {
		x = x * -1 // if it's negative, let's make it positive.
	}

	// for loop every number of the interger
	for x != 0 {
		// multiply current reverse step with 10 and add what remains when we divede x with 10.
		p("original", x)
		ret = ret*10 + x%10
		// divede number with 10
		p("reversed", ret)
		x /= 10
	}

	if sign {
		ret *= -1
	}

	if math.MaxInt32 < ret || math.MinInt32 > ret {
		ret = 0
	}

	return int(ret)
}

/**
which console logs following

original 123
reversed 3
original 12
reversed 32
original 1
reversed 321
*/
