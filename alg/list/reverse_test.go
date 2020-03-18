package list

import "testing"

func TestReverseWithStack(t *testing.T) {
	l := GenSList([]int{1, 2, 3, 4, 5})
	l = ReverseWithStack(l)
	t.Logf("2233333333\n")
	l.ForEach(func(node *SNode) {
		t.Logf("reverse %d", node.val)
	})
}
