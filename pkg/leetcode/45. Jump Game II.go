package leetcode

import "math"

func Jump(nums []int) int {
	steps := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		steps[i] = math.MaxInt32
	}
	return jumpBae(nums, 0, steps)
}

func jumpBae(nums []int, index int, steps []int) int {
	if index >= len(nums)-1 {
		return 0
	}
	if nums[index] == 0 {
		return math.MaxInt32
	}
	if steps[index] != math.MaxInt32 {
		return steps[index]
	}
	step := math.MaxInt32
	for i := 1; i <= nums[index]; i++ {
		if index+i < len(nums) {
			step = min(step, jumpBae(nums, index+i, steps)+1)
		}
	}
	steps[index] = step
	return steps[index]
}
