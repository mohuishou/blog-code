package list

import "math"

// 归并，直接在链表上归并
func sortList2(head *ListNode) *ListNode {
	left, right := sortListFindMid(head)
	if right == nil {
		return left
	}

	return sortListFindMergeSorted(sortList2(left), sortList2(right))
}

// 查找一个链表的中间节点
// 快慢指针，快指针走到结束时，慢指针刚好在中间
func sortListFindMid(head *ListNode) (left, right *ListNode) {
	if head == nil {
		return nil, nil
	}

	n := &ListNode{Next: head}
	slow, fast := n, n
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 截断链表
	right = slow.Next
	slow.Next = nil
	return head, right
}

// sortListFindMergeSorted 合并两个有序链表
func sortListFindMergeSorted(left, right *ListNode) *ListNode {
	res := &ListNode{}
	n := res
	for left != nil && right != nil {
		if left.Val < right.Val {
			n.Next = left
			left = left.Next
		} else {
			n.Next = right
			right = right.Next
		}
		n = n.Next
	}
	if left != nil {
		n.Next = left
	}
	if right != nil {
		n.Next = right
	}
	return res.Next
}

// 归并，先转换成数组在排序，不用考虑指针问题
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var data []*ListNode
	for head != nil {
		n := head
		head = head.Next
		n.Next = nil
		data = append(data, n)
	}

	res := sortListMerge(data[:len(data)/2], data[len(data)/2:])
	head = res[0]
	for i := 0; i < len(res)-1; i++ {
		res[i].Next = res[i+1]
	}
	return head
}

func sortListMerge(left, right []*ListNode) []*ListNode {
	if len(left) == 0 || len(right) == 0 {
		if len(left) != 0 {
			return left
		}
		return right
	}
	sortedLeft := sortListMerge(left[:len(left)/2], left[len(left)/2:])
	sortedRight := sortListMerge(right[:len(right)/2], right[len(right)/2:])
	var res []*ListNode
	i, j := 0, 0
	for i < len(sortedLeft) && j < len(sortedRight) {
		if sortedLeft[i].Val < sortedRight[j].Val {
			res = append(res, sortedLeft[i])
			i++
			continue
		}
		res = append(res, sortedRight[j])
		j++
	}
	if i < len(sortedLeft) {
		res = append(res, sortedLeft[i:]...)
	}
	if j < len(sortedRight) {
		res = append(res, sortedRight[j:]...)
	}
	return res
}

// 插入排序，不行，会超出时间限制
func sortListInsert(head *ListNode) *ListNode {
	res := &ListNode{Val: math.MinInt64}
	for head != nil {
		prev := res
		n := res.Next
		for n != nil && head.Val > n.Val {
			prev = n
			n = n.Next
		}
		prev.Next = head
		head = head.Next
		prev.Next.Next = n
	}
	return res.Next
}
