package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	// 找到交界位置
	i := 0
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}

	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}
