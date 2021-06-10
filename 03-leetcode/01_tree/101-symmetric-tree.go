package tree

func isSymmetric(root *TreeNode) bool {
	current := []*TreeNode{root}
	for len(current) > 0 {
		next := []*TreeNode{}
		for i := 0; i < len(current); i++ {
			node := current[i]
			other := current[len(current)-i-1]

			if node == nil && other == nil {
				continue
			}

			if node == nil || other == nil {
				return false
			}

			if node.Val != other.Val {
				return false
			}

			next = append(next, node.Left, node.Right)
		}
		current = next
	}
	return true
}
