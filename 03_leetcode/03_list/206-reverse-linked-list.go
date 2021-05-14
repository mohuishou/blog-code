package list

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	next := current.Next
	current.Next = nil
	for next != nil {
		n := next.Next
		next.Next = current
		current = next
		next = n
	}
	return current
}
