package dp

import "testing"

func TestCoinExchange(t *testing.T) {
	coins := []int{20, 10, 2, 5}
	count := CoinExchange(50, coins)
	t.Logf("count: %d", count)
}
