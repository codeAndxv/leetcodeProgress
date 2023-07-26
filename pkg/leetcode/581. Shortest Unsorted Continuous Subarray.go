package leetcode

import "math"

func FindUnsortedSubarray(nums []int) int {
	minValue := math.MaxInt32
	maxValue := math.MinInt32
	for index := 0; index < len(nums)-1; index++ {
		if nums[index] > nums[index+1] {
			minValue = min(minValue, nums[index+1])
			maxValue = max(maxValue, nums[index])
		}
	}
	minLeft := len(nums)
	maxRight := 0
	for index := 0; index < len(nums)-1; index++ {
		if nums[index] > minValue {
			minLeft = index
			break
		}

	}
	for index := len(nums) - 1; index >= 0; index-- {
		if nums[index] < maxValue {
			maxRight = index
			break
		}
	}
	return max(0, maxRight-minLeft+1)
}

/**
leftInit := false
	leftIndex := len(nums) - 1
	rightInit := false
	rightIndex := 0
	for index := 0; index < len(nums)-1; index++ {
		if nums[index] > nums[index+1] {
			if !leftInit {
				leftIndex = min(leftIndex, index)
				leftInit = true
			}
			index1 := leftIndex - 1
			for ; index1 >= 0; index1-- {
				if nums[index1] <= nums[index+1] {
					leftIndex = index1 + 1
					break
				}
			}
			if index1 < 0 {
				if nums[index1+1] > nums[index+1] {
					leftIndex = min(leftIndex, 0)
					break
				}
			}
		}
	}
	for index := len(nums) - 1; index > 0; index-- {
		if nums[index] < nums[index-1] {
			if !rightInit {
				rightIndex = max(rightIndex, index)
				rightInit = true
			}

			index1 := rightIndex + 1
			for ; index1 < len(nums); index1++ {
				if nums[index1] >= nums[index-1] {
					rightIndex = index1
					break
				}
			}
			if index1 == len(nums) {
				if nums[index1-1] < nums[index-1] {
					rightIndex = max(rightIndex, index1)
					break
				}
			}
		}
	}
	if rightIndex > leftIndex {
		return rightIndex - leftIndex
	}
	return 0
*/
