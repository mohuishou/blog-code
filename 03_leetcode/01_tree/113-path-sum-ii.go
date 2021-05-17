package tree

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}

	var res [][]int

	if data := pathSum(root.Left, targetSum-root.Val); data != nil {
		for _, n := range data {
			res = append(res, append([]int{root.Val}, n...))
		}
	}
	if data := pathSum(root.Right, targetSum-root.Val); data != nil {
		for _, n := range data {
			res = append(res, append([]int{root.Val}, n...))
		}
	}
	return res
}
