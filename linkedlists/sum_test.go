package linkedlists

import (
	"testing"
)

var sumForwardTests = []struct {
	l1     *ListNode
	l2     *ListNode
	result *ListNode
}{
	/*{
		&ListNode{1, nil},
		&ListNode{1, nil},
		&ListNode{2, nil},
	},
	{
		&ListNode{1, &ListNode{1, nil}},
		&ListNode{1, &ListNode{1, nil}},
		&ListNode{2, &ListNode{2, nil}},
	},
	{
		&ListNode{1, nil},
		&ListNode{1, &ListNode{1, nil}},
		&ListNode{2, &ListNode{1, nil}},
	},*/
	{
		nil,
		&ListNode{1, &ListNode{1, nil}},
		&ListNode{1, &ListNode{1, nil}},
	}, /*
		{
			&ListNode{8, &ListNode{1, nil}},
			&ListNode{7, &ListNode{1, nil}},
			&ListNode{5, &ListNode{3, nil}},
		},*/
}

func TestSumForward(t *testing.T) {
	for _, test := range sumForwardTests {
		if result := sumForward(test.l1, test.l2, 0); !ListNodeEquals(result, test.result) {
			t.Errorf("Expect: %v\nReceived: %v\n", test.result.print(), result.print())
		}
	}
}
