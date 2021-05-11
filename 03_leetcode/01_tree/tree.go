package tree

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

var null int = -1e9

func NewTreeNode(data []int, i int) *TreeNode {
	if len(data) <= i || data[i] == null {
		return nil
	}
	n := &TreeNode{Val: data[i]}
	n.Left = NewTreeNode(data, 2*i+1)
	n.Right = NewTreeNode(data, 2*i+2)
	return n
}

func (n *TreeNode) array() []int {
	var res []int
	if n == nil {
		return res
	}

	current := []*TreeNode{n}

	for len(current) > 0 {
		var next []*TreeNode

		for _, n := range current {
			if n == nil {
				res = append(res, null)
				continue
			}
			res = append(res, n.Val)
			next = append(next, n.Left)
			next = append(next, n.Right)
		}
		current = next
	}
	return res
}
