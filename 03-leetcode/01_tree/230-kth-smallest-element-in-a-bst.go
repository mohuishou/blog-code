package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var kthSmallestIndex int

func kthSmallest(root *TreeNode, k int) int {
	kthSmallestIndex = 0
	n := kthSmallestSearch(root, k)
	if n == nil {
		return -1
	}
	return n.Val
}

func kthSmallestSearch(root *TreeNode, k int) *TreeNode {
	if root == nil {
		return nil
	}
	if data := kthSmallestSearch(root.Left, k); data != nil {
		return data
	}
	kthSmallestIndex++
	if kthSmallestIndex == k {
		return root
	}
	if data := kthSmallestSearch(root.Right, k); data != nil {
		return data
	}
	return nil
}
