package sort

import "fmt"

func swap(items []int, i, j int) {
	tmp := items[i]
	items[i] = items[j]
	items[j] = tmp
}

func Partion2(items []int) int {
	idx := 0

	i := 1
	for j := 1; j < len(items); j++ {
		if items[j] < items[idx] {
			swap(items, i, j)
			i++
		}
	}
	swap(items, i-1, idx)

	return i - 1
}

func Partion(items []int) int {
	if len(items) < 2 {
		return 0
	}
	idx := 0 // select first item
	i := 1
	j := len(items) - 1

	for i <= j {
		if items[i] < items[idx] {
			i++
		} else if items[j] > items[idx] {
			j--
		} else {
			swap(items, i, j)
			i++
			j--
		}
	}

	swap(items, idx, i-1)
	return i - 1
}

func QuickSort(items []int) {
	if len(items) < 2 {
		return
	}
	fmt.Println(items)
	pos := Partion(items)

	QuickSort(items[:pos])
	QuickSort(items[pos+1:])
}

func SelectSort(items []int) {
	max := func(items []int) int {
		idx := 0
		for i, v := range items {
			if v < items[idx] {
				idx = i
			}
		}
		return idx
	}

	for i := 0; i < len(items); i++ {
		idx := max(items[i:])
		idx += i
		fmt.Println("idx: ", idx)
		if i != idx {
			swap(items, i, idx)
		}
	}
}

func BubbleSort(items []int) {
	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-i; j++ {
			if items[j] > items[j+1] {
				swap(items, j, j+1)
			}
		}
	}
}
