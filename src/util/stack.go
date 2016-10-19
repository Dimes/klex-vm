package util

type Stack struct {
	s []interface{}
}

func NewStack() *Stack {
	return &Stack {make([]interface{}, 0)}
}

func CopyStack(stack *Stack) *Stack {
	clone := make([]interface{}, len(stack.s))
	copy(clone, stack.s)
	return &Stack {clone}
}

func (s *Stack) At(index int) interface{} {
	return s.s[index]
}

func (s *Stack) Push(v interface{}) {
	s.s = append(s.s, v)
}

func (s *Stack) Peek() interface{} {
	return s.s[len(s.s) - 1]
}

func (s *Stack) Pop() interface{} {
	l := len(s.s)
	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}
