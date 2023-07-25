package leetcode

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	var maxValue int
	lastAccum := 0
	for index, value := range nums {
		if index == 0 {
			maxValue = value
			lastAccum = value
			continue
		}
		if lastAccum >= 0 {
			lastAccum += value
		} else {
			lastAccum = value
		}
		maxValue = max(lastAccum, maxValue)
	}
	return maxValue
}

func main() {
	bills := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(bills))
}
