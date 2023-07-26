package leetcode

import "math"

func rob213(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	records := make([]int, len(nums)+3)
	for index, _ := range records {
		records[index] = math.MaxInt32
	}
	tem1 := robBase(nums[:len(nums)-1], 0, records)
	for index, _ := range records {
		records[index] = math.MaxInt32
	}
	tem2 := robBase(nums[1:], 0, records)
	return max(tem1, tem2)
}
