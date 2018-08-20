package linkedlists

import (
	"testing"
)

var partitionTests = []struct {
	l      *ListNode
	result *ListNode
	x      int
}{
	{
		&ListNode{1, nil},
		&ListNode{1, nil},
		0,
	},
	{
		&ListNode{1, &ListNode{1, nil}},
		&ListNode{1, &ListNode{1, nil}},
		1,
	},
}

func TestPartition(t *testing.T) {
	for _, test := range partitionTests {
		if result := partition(test.l, test.x); !ListNodeEquals(result, test.result) {
			t.Errorf("error")
		}
	}
}
