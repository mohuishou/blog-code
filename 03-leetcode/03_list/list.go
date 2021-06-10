package list

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}

	head := &ListNode{Val: data[0]}
	n := head
	for i := 1; i < len(data); i++ {
		n.Next = &ListNode{Val: data[i]}
		n = n.Next
	}

	return head
}

func (n *ListNode) array() []int {
	var res []int
	for n != nil {
		res = append(res, n.Val)
		n = n.Next
	}
	return res
}
