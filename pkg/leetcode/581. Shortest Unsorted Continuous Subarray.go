package leetcode

func FindUnsortedSubarray(nums []int) int {
	var index1, index2, index3 int
	
	for index:= 0; index < len(nums)-1; index++ {
		if nums[index] < nums[]
	}
	for ; index1 < len(nums)-1; index1++ {
		if nums[index1] > nums[index1+1] {
			break
		}
	}
	for index2 = len(nums) - 1; index2 > 0; index2-- {
		if nums[index2] < nums[index2-1] {
			break
		}
	}
	if index2 > index1 {
		return index2 - index1 + 1
	}
	return 0
}
