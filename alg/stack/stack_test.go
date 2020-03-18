package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack(16)
	s.Push(1)
	s.Push(2)

	ret := s.Pop()
	v, ok := ret.(int)
	if !ok || v != 2 {
		t.Errorf("pop value wrong")
	}
	if s.Empty() {
		t.Errorf("should not be empty")
	}

	ret = s.Pop()
	v, ok = ret.(int)
	if !ok || v != 1 {
		t.Errorf("pop value wrong")
	}
	if !s.Empty() {
		t.Errorf("should be empty")
	}
}
