package tree

type inorderTraversalStack struct {
	data []*TreeNode
	tail int
}

func (s *inorderTraversalStack) push(n *TreeNode) {
	if s.tail >= len(s.data) {
		s.data = append(s.data, make([]*TreeNode, len(s.data))...)
	}
	s.data[s.tail] = n
	s.tail++
}

func (s *inorderTraversalStack) pop() *TreeNode {
	if s.empty() {
		panic("empty stack")
	}
	s.tail--
	return s.data[s.tail]
}

func (s *inorderTraversalStack) empty() bool {
	return s.tail == 0
}
func inorderTraversal(root *TreeNode) []int {
	stack := &inorderTraversalStack{data: make([]*TreeNode, 100)}
	n := root
	var res []int
	for n != nil || !stack.empty() {
		for n != nil {
			stack.push(n)
			n = n.Left
		}
		n = stack.pop()
		res = append(res, n.Val)
		n = n.Right
	}
	return res
}
