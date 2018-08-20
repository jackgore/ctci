package linkedlists

func insert(l *ListNode, c *ListNode) *ListNode {
	c.next = l
	return c
}

func partition(l *ListNode, x int) *ListNode {
	var left *ListNode
	var right *ListNode

	cur := l

	for cur != nil {
		next := cur.next
		if cur.value < x {
			left = insert(left, cur)
		} else {
			right = insert(right, cur)
		}

		cur = next
	}

	// Case where there is nothing in left partition
	if left == nil {
		return right
	} else if right == nil {
		return left // vice versa
	}

	var prev *ListNode
	cur = left

	for cur != nil {
		prev = cur
		cur = cur.next
	}

	prev.next = right
	return left
}
