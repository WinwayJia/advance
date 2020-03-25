package bt

import "testing"

func TestPermutation(t *testing.T) {
	result := FullPermutation([]int{1, 2, 3})
	t.Logf("%+v", result)
	for idx, item := range result {
		t.Logf("idx: %d item: %+v", idx, item)
	}
}
