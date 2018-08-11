package linkedlists

type ListNode struct {
	value int
	next  *ListNode
}

func ListNodeLength(l *ListNode) int {
	count := 0

	for l != nil {
		count++
		l = l.next
	}

	return count
}

func ListNodeEquals(l1, l2 *ListNode) bool {
	len1 := ListNodeLength(l1)
	len2 := ListNodeLength(l2)

	if len1 != len2 {
		return false
	}

	for l1 != nil {
		if l1.value != l2.value {
			return false
		}

		l1 = l1.next
		l2 = l2.next
	}

	return true
}
