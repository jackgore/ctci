package linkedlists

func deletedups(l *ListNode) *ListNode {
	vals := make(map[int]bool)
	var prev *ListNode
	cur := l

	for cur != nil {
		val := cur.value
		if _, ok := vals[val]; ok {
			prev.next = cur.next
			cur = cur.next
		} else {
			vals[val] = true
			prev = cur
			cur = cur.next
		}
	}

	return l
}

func checkExists(l *ListNode, v int) bool {
	cur := l
	for cur != nil {
		if cur.value == v {
			return true
		}

		cur = cur.next
	}

	return false
}

func nobufferdups(l *ListNode) *ListNode {
	var prev *ListNode
	cur := l
	head := l

	for cur != nil {
		if checkExists(cur.next, cur.value) {
			// The value exists so delete the node

			// Two case first is prev is null i.e. deleting head
			if prev == nil {
				cur = cur.next
				head = cur
			} else {
				prev.next = cur.next
				cur = cur.next
			}
			continue
		}

		prev = cur
		cur = cur.next
	}

	return head
}
