package list

import "math"

func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{}
	n := res
	for len(lists) > 0 {
		min := &ListNode{Val: math.MaxInt32}
		minIndex := -1
		for i, l := range lists {
			if l == nil {
				continue
			}
			if l.Val < min.Val {
				min = l
				minIndex = i
			}
		}
		if minIndex == -1 {
			break
		}
		lists[minIndex] = lists[minIndex].Next
		n.Next = min
		n = min
	}
	return res.Next
}
