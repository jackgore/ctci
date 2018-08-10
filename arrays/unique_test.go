package arrays

import (
	"testing"
)

var isUniqueTests = []struct {
	s      string
	unique bool
}{
	{"jack", true},
	{"jjack", false},
	{"", true},
	{"a", true},
	{"abcdefghijklmnopqrstuvwyxz", true},
	{"aa", false},
	{"abcdefghijklmnopqrstuvwyxza", false},
}

func TestIsUnique(t *testing.T) {
	for _, test := range isUniqueTests {
		if unique := isUniqueLoops(test.s); unique != test.unique {
			t.Fail()
		}
		if unique := isUniqueSort(test.s); unique != test.unique {
			t.Fail()
		}
		if unique := isUniqueMap(test.s); unique != test.unique {
			t.Fail()
		}
	}
}
