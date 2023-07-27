package leetcode

func findKthLargest(nums []int, k int) int {
	return 0
}

6, 7 ,10 , 223, 31, 20, 12, 15, 90


/**
var num1, num2, num3 int
var nums1, nums3 []int
for _, value := range nums {
	if value == nums[0] {
		num2++
		continue
	}
	if value > nums[0] {
		num3++
		nums3 = append(nums3, value)
		continue
	}
	if value < nums[0] {
		num1++
		nums1 = append(nums1, value)
		continue
	}
}
if num3 >= k {
	return FindKthLargest(nums3, k)
}
if num3+num2 >= k {
	return nums[0]
}
return FindKthLargest(nums1, k-num3-num2)
*/
