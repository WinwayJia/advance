package dp

import (
	"fmt"
	"math"
)

// CoinExchange 找零钱
// @sum 要找钱的总数
// @coins 硬币面值，硬币数量不限
func CoinExchange(sum int, coins []int) int {
	if 0 == sum || len(coins) == 0 {
		return 0
	}

	max := len(coins) - 1

	results := make([]int, sum+1, sum+1)

	// 正好是某个硬币的面值
	for i := 0; i <= max; i++ {
		results[coins[i]] = 1
	}

	for i := 1; i <= sum; i++ {
		if results[i] == 1 {
			continue
		}
		min := math.MaxInt32
		for j := 0; j <= max; j++ {
			if i > coins[j] && results[i-coins[j]] != 0 {
				if results[i-coins[j]] < min {
					min = results[i-coins[j]]
				}
			}
		}
		if min < math.MaxInt32 {
			results[i] = min + 1
		}
	}
	fmt.Println(results[1:])

	return results[sum]
}
