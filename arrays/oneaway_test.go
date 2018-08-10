package arrays

import (
	"testing"
)

var oneawayTests = []struct {
	s1      string
	s2      string
	oneaway bool
}{
	{"jack", "jack", true},
	{"jac", "jack", true},
	{"pale", "ple", true},
	{"pales", "pale", true},
	{"pale", "bale", true},
	{"pale", "bake", false},
	{"jacs", "jack", true},
	{"", "k", true},
	{"jacky", "jack", true},
	{"once upon a time there was a man", "once upon a time there was a man!", true},
	{"jack", "kcaj", false},
}

func TestOneaway(t *testing.T) {
	for _, test := range oneawayTests {
		if oneaway := oneaway(test.s1, test.s2); oneaway != test.oneaway {

			t.Errorf("for strings %v and %v", test.s1, test.s2)
		}
	}
}
