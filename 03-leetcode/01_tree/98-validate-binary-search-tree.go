package tree

func isValidBST(root *TreeNode) bool {
	_, _, valid := isValidBSTSearch(root)
	return valid
}

func isValidBSTSearch(root *TreeNode) (min, max int, valid bool) {
	min, max = root.Val, root.Val

	if root.Left != nil {
		lmin, lmax, valid := isValidBSTSearch(root.Left)
		if !valid {
			return lmin, lmax, valid
		}

		if lmax >= root.Val {
			return min, max, false
		}
		min = lmin
	}

	if root.Right != nil {
		rmin, rmax, valid := isValidBSTSearch(root.Right)
		if !valid {
			return min, max, valid
		}

		if rmin <= root.Val {
			return min, max, false
		}
		max = rmax
	}

	return min, max, true
}
