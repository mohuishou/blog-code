package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	current := []*TreeNode{root}

	for len(current) > 0 {
		res = append(res, current[len(current)-1].Val)
		var next []*TreeNode

		for _, n := range current {
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		current = next
	}
	return res
}
