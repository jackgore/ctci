package arrays

import (
	"testing"
)

var urlifyTests = []struct {
	s   string
	l   int
	url string
}{
	{"jack", 4, "jack"},
	{"   ", 1, "%20"},
	{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
}

func TestUrlify(t *testing.T) {
	for _, test := range urlifyTests {
		if url := urlify(test.s, test.l); url != test.url {
			t.Fail()
		}
	}
}
