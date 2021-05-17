package tree

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

var null int = -1e9

// 兼容
func NewTreeNode(data []int, i int) *TreeNode {
	return NewTree(data)
}

func NewTree(data []int) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	root := NewTreeNodeVal(data[0])
	current := []*TreeNode{root}

	for i := 1; len(current) > 0 && i < len(data); {
		var next []*TreeNode
		for _, n := range current {
			if n == nil {
				continue
			}
			n.Left = NewTreeNodeVal(data[i])
			n.Right = NewTreeNodeVal(data[i+1])
			next = append(next, n.Left, n.Right)
			i += 2
		}
		current = next
	}
	return root
}

func NewTreeNodeVal(v int) *TreeNode {
	if v == null {
		return nil
	}
	return &TreeNode{Val: v}
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

		allNil := true
		for _, n := range next {
			if n != nil {
				allNil = false
				break
			}
		}
		if allNil {
			next = []*TreeNode{}
		}
		current = next
	}
	return res
}
