package list

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	stack := make([]*ListNode, right-left+1)
	node := &ListNode{Next: head}
	prev := node
	for i := 1; i <= right; i++ {
		if i == left-1 {
			prev = head
		}
		if i >= left {
			stack[i-left] = head
		}
		head = head.Next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		prev.Next = stack[i]
		prev = stack[i]
	}
	prev.Next = head

	return node.Next
}
