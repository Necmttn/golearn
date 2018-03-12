package parser

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	data := ParseFile("./test.xml")
	numberOfLines := len(data.Texts)
	if numberOfLines != 561 {
		t.Errorf("Total numberOfLines was inccorrect")
	}
	for i := 0; i < len(data.Texts); i++ {
		Start := data.Texts[i].Start
		Context := data.Texts[i].Context
		Dur := data.Texts[i].Dur

		if Start == "" {
			t.Errorf("No Start Valur for Text")
		}

		if Context == "" {
			t.Errorf("No Context Valur for Text")
		}

		if Dur == "" {
			t.Errorf("No Dur Valur for Text")
		}
	}
}
