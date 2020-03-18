package list

// single direction
type SNode struct {
	val  int
	prev *SNode
	next *SNode
}

// double direction
type DNode struct {
	val  int
	prev *DNode
	next *DNode
}

func NewSNode(val int) *SNode {
	return &SNode{
		val: val,
	}
}

func GenSList(vs []int) *SNode {
	head := NewSNode(0)
	curr := head
	for _, v := range vs {
		curr.next = NewSNode(v)
		curr = curr.next
	}

	return head.next
}

func (l *SNode) ForEach(f func(node *SNode)) {
	for l != nil {
		f(l)
		l = l.next
	}
}
