package parser

import (
	"fmt"
	"testing"
)

func TestParseFile(t *testing.T) {
	data := ParseFile("./test.xml")
	fmt.Println("data", data.Texts[1].Context)
}
