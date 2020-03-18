package stack

type Stack struct {
	top  int
	data []interface{}
}

func NewStack(size int) *Stack {
	s := &Stack{
		top:  0,
		data: make([]interface{}, 0, size),
	}
	return s
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
	s.top++
}

func (s *Stack) Pop() interface{} {
	ret := s.data[s.top-1]
	s.data = s.data[0:s.top]
	s.top--
	return ret
}

func (s *Stack) Empty() bool {
	return s.top <= 0
}
