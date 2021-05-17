package tree

func isCompleteTree(root *TreeNode) bool {
	current := []*TreeNode{root}
	isOver := true
	for len(current) > 0 {
		var next []*TreeNode
		hasNil := false
		for _, n := range current {
			if n == nil {
				hasNil = true
				continue
			}
			next = append(next, n.Left, n.Right)
			if hasNil {
				return false
			}
		}
		if !isOver && len(next) > 0 {
			return false
		}
		isOver = !hasNil
		current = next
	}
	return true
}
