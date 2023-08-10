package leetcode

import "math"

func coinChange(coins []int, amount int) int {
	records := make([]int, amount+1)
	records[0] = 0
	for i := 1; i <= amount; i++ {
		records[i] = math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && records[i-coins[j]] != math.MaxInt32 {
				records[i] = min(records[i], records[i-coins[j]]+1)
			}
		}
	}
	if records[amount] == math.MaxInt32 {
		return -1
	}
	return records[amount]
}
