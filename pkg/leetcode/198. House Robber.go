package leetcode

import "math"

func rob(nums []int) int {
	records := make([]int, len(nums)+3)
	for index, _ := range records {
		records[index] = math.MaxInt32
	}
	return robBase(nums, 0, records)

}

func robBase(nums []int, index int, records []int) int {
	if index >= len(nums) {
		return 0
	}
	if index == len(nums)-1 {
		records[index] = nums[index]
		return nums[index]
	}
	if records[index+2] == math.MaxInt32 {
		records[index+2] = robBase(nums, index+2, records)
	}
	if records[index+1] == math.MaxInt32 {
		records[index+1] = robBase(nums, index+1, records)
	}
	return max(records[index+2]+nums[index], records[index+1])
}
