package leetcode

func FindUnsortedSubarray(nums []int) int {
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
			if nums[leftIndex] > nums[index+1] {
				for index1 := leftIndex - 1; index1 >= 0; index1-- {
					if nums[index1] < nums[index+1] {
						leftIndex = index1 + 1
						break
					}
					if index1 == 0 {
						leftIndex = 0
					}
				}
			}
		}
	}
	for index := len(nums) - 2; index >= 0; index-- {
		if nums[index+1] < nums[index] {
			if !rightInit {
				rightIndex = max(rightIndex, index)
				rightInit = true
			}
			if nums[rightIndex] < nums[index] {
				for index1 := rightIndex + 1; index1 < len(nums); index1++ {
					if nums[index1] > nums[index-1] {
						rightIndex = index1
						break
					}
					if index1 == len(nums)-1 {
						rightIndex = len(nums) - 1
					}
				}
			}
		}
	}
	if rightIndex > leftIndex {
		return rightIndex - leftIndex + 1
	}
	return 0
}
