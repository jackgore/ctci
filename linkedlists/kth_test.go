package linkedlists

import (
	"testing"
)

var kthTests = []struct {
	l      *ListNode
	result *ListNode
	k      int
}{
	{
		&ListNode{1, nil},
		&ListNode{1, nil},
		0,
	},
	{
		&ListNode{1, &ListNode{1, nil}},
		&ListNode{1, nil},
		1,
	},
	{
		&ListNode{1, &ListNode{1, &ListNode{8, &ListNode{1, nil}}}},
		&ListNode{8, nil},
		1,
	},
	{
		&ListNode{1, &ListNode{21, &ListNode{8, &ListNode{9, nil}}}},
		&ListNode{1, nil},
		3,
	},
}

func TestKth(t *testing.T) {
	for _, test := range kthTests {
		if result := kth(test.l, test.k); !ListNodeEquals(result, test.result) {
			t.Errorf("error")
		}

		if result := kth3(test.l, test.k); !ListNodeEquals(result, test.result) {
			t.Errorf("error with kth3")
		}

		if result := kthrec(test.l, test.k); !ListNodeEquals(result, test.result) {
			t.Errorf("error: received %v, expect: %v", result, test.result)
		}
	}
}
