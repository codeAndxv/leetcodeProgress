package leetcode

import "math"

func minAbsoluteDifference(nums []int, x int) int {
	minValue := math.MaxInt32
	for i := x; i < len(nums); i++ {
		for j := 0; j <= i-x; j++ {
			minValue = min(minValue, abs(nums[i]-nums[j]))
		}
	}
	return minValue
}
