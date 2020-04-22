package bt

import "fmt"

func backtrack(path, items []int, visited map[int]bool, result *[][]int) {
	if len(items) == len(visited) {
		*result = append(*result, path)
		fmt.Printf("%d %d %v\n", len(items), len(visited), *result)
		return
	}

	for _, item := range items {
		// 做选择
		_, ok := visited[item]
		if ok {
			continue
		}
		visited[item] = true
		path = append(path, item)
		backtrack(path, items, visited, result)
		delete(visited, item)
		path = path[:len(path)-1]
	}
}

// FullPermutation 全排列
func FullPermutation(arr []int) [][]int {
	result := make([][]int, 0, 1024)
	path := make([]int, 0, 32)
	visited := make(map[int]bool, 64)
	backtrack(path, arr, visited, &result)
	fmt.Printf("result: %v", result)
	return result
}

func bt(path, arr []int, n int, result *[][]int) {
	if 0 == n {
		*result = append(*result, path)
		return
	}
}

// Combination 排列 C(n, m), n = len(arr)
func Combination(arr []int, m int) [][]int {
	// TODO
	path := make([]int, 0, 1024)
	for _, item := range arr {
		path = append(path, item)
	}
	return nil
}
