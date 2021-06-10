package tree

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var helper func(n *TreeNode) int
	helper = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		left := helper(n.Left)
		right := helper(n.Right)
		res = max(res, left+right)
		return max(left, right) + 1
	}
	helper(root)
	return res
}
