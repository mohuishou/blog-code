package list

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 1. 第一步找到两个链表的结尾
	a, b := headA, headB
	stackA, stackB := []*ListNode{}, []*ListNode{}
	for a.Next != nil {
		stackA = append(stackA, a)
		a = a.Next
	}

	for b.Next != nil {
		stackB = append(stackB, b)
		b = b.Next
	}

	if a != b {
		return nil
	}

	for i := 1; i <= len(stackA) && i <= len(stackB); i++ {
		if stackA[len(stackA)-i] != stackB[len(stackB)-i] {
			break
		}
		a = stackA[len(stackA)-i]
	}

	return a
}
