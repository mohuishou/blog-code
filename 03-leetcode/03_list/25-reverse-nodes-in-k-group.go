package list

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

// k 是一个正整数，它的值小于或等于链表的长度。

// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

// 进阶：

// 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil {
		return head
	}

	n := head
	head = nil
	stack := []*ListNode{}
	var prev *ListNode
	for n != nil {
		stack = append(stack, n)
		if len(stack) < k {
			n = n.Next
			continue
		}

		current := n
		next := n.Next

		if head == nil {
			head = current
		}

		if prev != nil {
			prev.Next = current
		}

		for i := len(stack) - 2; i >= 0; i-- {
			current.Next = stack[i]
			current = stack[i]
		}

		stack = stack[:0]
		current.Next = next
		prev = current
		n = next
	}

	return head
}
