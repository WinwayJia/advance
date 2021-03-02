package string

// 暴力搜索
func StrStr(target, pattern string) int {
	return -1
}

// StrStrHorspool horspool 算法，bm的简化版
func GenSkipTable(pattern string) []int {
	table := make([]int, 256, 256)
	for i := 0; i < 256; i++ {
		table[i] = len(pattern)
	}

	for i := 0; i < len(pattern)-1; i++ {
		table[pattern[i]] = len(pattern) - i - 1
	}
	return table
}

func StrStrHorspool(target, pattern string) int {
	table := GenSkipTable(pattern)

	for i := 0; i <= len(target)-len(pattern); {
		j := 0
		for ; j < len(pattern); j++ {
			if target[i+j] != pattern[j] {
				// 不匹配，跳转当前比较target最后的一个字符的table值
				i += table[target[i+len(pattern)-1]]
				break
			}
		}
		if j >= len(pattern) {
			return i
		}
	}

	return -1
}
