package tree

// 计算节点的最大贡献值
var sum int

func maxPathSum(root *TreeNode) int {
	sum = -1e9
	maxPathSumGain(root)
	return sum
}

func maxPathSumGain(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 子节点只有大于 0 才能产生贡献
	left := max(maxPathSumGain(root.Left), 0)
	right := max(maxPathSumGain(root.Right), 0)
	sum = max(left+root.Val+right, sum)

	return max(root.Val+left, root.Val+right)
}

func max(data ...int) int {
	var res int = -1e9
	for _, v := range data {
		if v > res {
			res = v
		}
	}
	return res
}
