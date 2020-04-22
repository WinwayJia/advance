package topk

import "testing"

func TestSwap(t *testing.T) {
	arr := []int{4, 0}
	swap(arr, 0, 1)
	if arr[0] != 0 || arr[1] != 4 {
		t.FailNow()
	}
}

func TestTopK(t *testing.T) {
	k := 3
	arr := []int{4, 0, 1, 2, 3, 5, -1, 7, 9}
	ans := TopK(arr, 4)
	t.Logf("top %d of %+v is %+v", k, arr, ans)
}
