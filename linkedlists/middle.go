package linkedlists

// O(N) too much work
func middle(l *ListNode) {
	cur := l
	var prev *ListNode

	for cur.next != nil {
		cur.value = cur.next.value
		prev = cur
		cur = cur.next
	}

	if prev != nil {
		prev.next = nil
	}
}

func middleConstant(l *ListNode) {
	if l.next != nil {
		l.value = l.next.value
		l.next = l.next.next
	}
}
