package arrays

import (
	"testing"
)

var palindromeTests = []struct {
	s          string
	palindrome bool
}{
	{"jack", false},
	{"Tact Coa", true},
	{" r a c e c a r", true},
}

func TestPalindrome(t *testing.T) {
	for _, test := range palindromeTests {
		if palindrome := palindrome(test.s); palindrome != test.palindrome {
			t.Fail()
		}
	}
}
