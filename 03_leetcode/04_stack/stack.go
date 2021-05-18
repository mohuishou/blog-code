package stack

type Stack struct {
	data []int
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() int {
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}
