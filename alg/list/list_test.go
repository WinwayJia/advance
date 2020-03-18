package list

import "testing"

func TestGenSList(t *testing.T) {

	l := GenSList([]int{1, 2, 3, 4, 5})
	for l != nil {
		t.Logf("val: %d", l.val)
		l = l.next
	}
}
