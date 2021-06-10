package list

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	n := res
	for l1 != nil && l2 != nil {

		if l1.Val > l2.Val {
			n.Next = l2
			l2 = l2.Next
		} else {
			n.Next = l1
			l1 = l1.Next
		}
		n = n.Next
	}
	if l1 == nil {
		n.Next = l2
	}
	if l2 == nil {
		n.Next = l1
	}
	return res.Next
}
