package arrays

import (
	"testing"
)

var permutationTests = []struct {
	s1   string
	s2   string
	perm bool
}{
	{"jack", "jack", true},
	{"jjack", "jack", false},
	{"", "", true},
	{"jack", "jcak", true},
	{"racecar", "racecar", true},
	{"once upon a time", "upon a time once", true},
}

func TestPermutation(t *testing.T) {
	for _, test := range permutationTests {
		if perm := permutationSort(test.s1, test.s2); perm != test.perm {
			t.Fail()
		}
		if perm := permutationMap(test.s1, test.s2); perm != test.perm {
			t.Fail()
		}
	}
}
