package common

type Stack struct {
	c []*Node
}

func NewStack() *Stack {
	return &Stack{[]*Node{}}
}

// Push push
func (s *Stack) Push(r *Node) {
	s.c = append(s.c, r)
}

// Pop pop
func (s *Stack) Pop() *Node {

	top := len(s.c) - 1
	r := s.c[top]
	s.c = s.c[:top]
	return r
}

// Size size
func (s *Stack) Size() int {
	return len(s.c)
}
