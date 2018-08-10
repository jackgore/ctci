package arrays

import (
	"testing"
)

var compressionTests = []struct {
	s           string
	compression string
}{
	{"jack", "jack"},
	{"jac", "jac"},
	{"jjjjjjjjjjjjjjjjac", "j16a1c1"},
	{"jjjaaaccckkk", "j3a3c3k3"},
	{"aabcccccaaa", "a2b1c5a3"},
}

func TestCompressino(t *testing.T) {
	for _, test := range compressionTests {
		if compression := compression(test.s); compression != test.compression {

			t.Errorf("for string %v expected %v but received %v", test.s, test.compression, compression)
		}
	}
}
