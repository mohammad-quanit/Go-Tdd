package stack

type Stack struct {
	isEmpty bool
}

func NewStack() *Stack {
	return &Stack{isEmpty: true}
}

func (s *Stack) Empty() bool {
	return s.isEmpty
}

func (s *Stack) Add(value string) {
	s.isEmpty = false
}
