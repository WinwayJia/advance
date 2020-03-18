package list

import (
	"fmt"
	"stack"
)

func Reverse(l *SNode) *SNode {
	// TODO
	return l
}

// 逆序链表借助栈
func ReverseWithStack(l *SNode) *SNode {
	s := stack.NewStack(1024)
	for l != nil {
		s.Push(l)
		l = l.next
	}

	val := s.Pop()
	ret, ok := val.(*SNode)
	if !ok {
		return nil
	}
	curr := ret
	for !s.Empty() {
		val := s.Pop()
		node, ok := val.(*SNode)
		if ok {
			curr.next = node
			node.next = nil
			curr = curr.next
		}
	}

	return ret
}
