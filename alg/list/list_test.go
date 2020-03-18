package list

import "testing"

func TestGenSList(t *testing.T) {

	l := GenSList([]int{1, 2, 3, 4, 5})
	for l != nil {
		t.Logf("val: %d", l.val)
		l = l.next
	}
}

func TestSNode_ForEach(t *testing.T) {
	l := GenSList([]int{1, 2, 3, 4, 5})
	l.ForEach(func(node *SNode) {
		t.Logf("ForEach: %d", node.val)
	})
}
