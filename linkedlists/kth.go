package linkedlists

func kth(l *ListNode, k int) *ListNode {
	length := ListNodeLength(l)

	i := length - k - 1

	if i < 0 {
		return nil
	}

	cur := l
	curIndex := 0

	for cur != nil {
		if curIndex == i {
			return &ListNode{cur.value, nil}
		}
		curIndex++
		cur = cur.next
	}

	return nil
}

func kthhelper(l *ListNode, i, target int) *ListNode {
	if l == nil {
		return nil
	}

	if target == i {
		return &ListNode{l.value, nil}
	}

	return kthhelper(l.next, i+1, target)
}

func kthrec(l *ListNode, k int) *ListNode {
	length := ListNodeLength(l)

	i := length - k - 1

	if i < 0 {
		return nil
	}
	return kthhelper(l, 0, i)
}

func kth3(l *ListNode, k int) *ListNode {
	n, _ := kth3helper(l, k)
	return n
}

func kth3helper(l *ListNode, k int) (*ListNode, int) {
	if l == nil {
		return nil, 0
	}

	node, counter := kth3helper(l.next, k)
	if node != nil {
		return node, counter
	} else if counter == k {
		return &ListNode{l.value, nil}, counter
	}

	return nil, counter + 1
}
