package leetcode

func searchInsert(nums []int, target int) int {
	return searchInsertBase(nums, target, 0, len(nums)-1)
}

func searchInsertBase(nums []int, target int, left int, right int) int {
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] < target {
		if mid < right {
			return searchInsertBase(nums, target, mid+1, right)
		} else {
			return mid + 1
		}
	} else {
		if mid > left {
			return searchInsertBase(nums, target, left, mid-1)
		} else {
			return mid
		}
	}
}
