package leetcode

func FindMin(nums []int) int {
	return findMinBase(nums, 0, len(nums)-1)
}

func findMinBase(nums []int, left int, right int) int {
	//mid := (left + right) / 2
	//if nums[left] <= nums[mid] && nums[mid] <=
	//
	//
	//if right == left {
	//	return left
	//}
	//if right > mid && nums[mid] > nums[right] {
	//	return findMinBase(nums, mid, right)
	//}
	//if mid > left && nums[left] > nums[mid] {
	//	return findMinBase(nums, left, mid)
	//}
	//return -1
	return 1
}
