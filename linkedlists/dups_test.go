package linkedlists

import (
	"testing"
)

var dupsTests = []struct {
	l      *ListNode
	result *ListNode
}{
	{
		&ListNode{1, nil},
		&ListNode{1, nil},
	},
	{
		&ListNode{1, &ListNode{1, nil}},
		&ListNode{1, nil},
	},
	{
		&ListNode{1, &ListNode{1, &ListNode{8, &ListNode{1, nil}}}},
		&ListNode{1, &ListNode{8, nil}},
	},
}

func TestDups(t *testing.T) {
	for _, test := range dupsTests {
		if result := deletedups(test.l); !ListNodeEquals(result, test.result) {
			t.Errorf("error")
		}

		if result := nobufferdups(test.l); !ListNodeEquals(result, test.result) {
			t.Errorf("error: received %v, expect: %v", result, test.result)
		}
	}
}
