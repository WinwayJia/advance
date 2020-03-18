package sort

import "testing"

func TestPartion(t *testing.T) {
	data := []int{4, 0, 5, 3}
	pos := Partion(data)
	t.Logf("%+v %+v", data, pos)
}

func TestPartion2(t *testing.T) {
	data := []int{4, 0, 5, 3}
	pos := Partion(data)
	t.Logf("%+v %+v", data, pos)
}

func TestQuickSort(t *testing.T) {
	data := []int{4, 0, 5, 3}
	QuickSort(data)
	t.Logf("%+v", data)
}

func TestSelectSort(t *testing.T) {
	data := []int{4, 0, 5, 3}
	SelectSort(data)
	t.Logf("%+v", data)
}

func TestBubbleSort(t *testing.T) {
	data := []int{4, 0, 5, 3}
	BubbleSort(data)
	t.Logf("%+v", data)
}
