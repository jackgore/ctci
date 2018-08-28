package linkedlists

//import "fmt"

func sumForward(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry == 0 {
			return nil
		}

		return &ListNode{carry, nil}
	}

	var v1, v2 int
	var l1next, l2next *ListNode

	if l1 == nil {
		v2 = l2.value
		l2next = l2.next
	} else if l2 == nil {
		v1 = l1.value
		l1next = l1.next
	} else {
		v1 = l1.value
		v2 = l2.value
		l1next = l1.next
		l2next = l2.next
	}

	temp := v1 + v2 + carry

	if temp > 10 {
		temp -= 10
		carry = 1
	} else {
		carry = 0
	}

	head := &ListNode{temp, sumForward(l1next, l2next, carry)}

	return head
}

/*
func sumForward(l1, l2 *ListNode) *ListNode {
	var sum *ListNode
	var cur *ListNode
	carry := 0

	for l1 != nil && l2 != nil {
		temp := l1.value + l2.value + carry

		if temp >= 10 {
			temp -= 10
		} else {
			carry = 0
		}

		// Now insert into sum list checking if cur is nil
		n := &ListNode{temp, nil}
		if cur == nil {
			cur = n
			sum = n
		} else {
			cur.next = n
		}

		l1 = l1.next
		l2 = l2.next
	}

	if l1 != nil {
		cur.next, carry = carryOver(l1, carry)
		cur = cur.next
	} else if l2 != nil {
		cur.next, carry = carryOver(l2, carry)
		cur = cur.next
	}

	// Last step is to check the final carry
	if carry > 0 {
		cur.next = &ListNode{1, nil}
	}

	return sum
}

func carryOver(l *ListNode, carry int) (*ListNode, int) {
	head := l
	for l != nil {
		temp := l.value + carry
		if temp > 10 {
			temp -= 10
			l.value = temp
		} else {
			carry = 0
		}

		l = l.next
	}

	return head, carry
} */
