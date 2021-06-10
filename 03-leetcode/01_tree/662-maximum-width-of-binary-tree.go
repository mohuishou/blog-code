package tree

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	root.Val = 0
	current := []*TreeNode{root}
	max := 0
	for len(current) > 0 {
		var next []*TreeNode

		for _, n := range current {
			if n.Left != nil {
				n.Left.Val = 2 * n.Val
				next = append(next, n.Left)
			}
			if n.Right != nil {
				n.Right.Val = 2*n.Val + 1
				next = append(next, n.Right)
			}
		}
		val := current[len(current)-1].Val + 1 - current[0].Val
		if val > max {
			max = val
		}
		current = next
	}
	return max
}
