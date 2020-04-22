package topk

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func heapify(arr []int, idx int) {
	min := idx
	l := idx*2 + 1
	if l < len(arr) && arr[min] < arr[l] {
		min = l
	}

	r := idx*2 + 2
	if r < len(arr) && arr[min] < arr[r] {
		min = r
	}

	if min != idx {
		swap(arr, min, idx)
		// 递归，确保最小的数沉到最深层
		heapify(arr, min)
	}
}

// 前N最小的值，需要构建大顶堆
func TopK(arr []int, k int) []int {
	if k >= len(arr) {
		return arr
	}

	heap := make([]int, k, k)
	for i := 0; i < k; i++ {
		heap[i] = arr[i]
	}

	// 从后向前，配合heapify递归完成堆化
	for i := len(heap) + 1; i >= 0; i-- {
		heapify(heap, i)
	}

	for i := k; i < len(arr); i++ {
		if arr[i] < heap[0] {
			heap[0] = arr[i]
			heapify(heap, 0)
		}
	}
	return heap
}
