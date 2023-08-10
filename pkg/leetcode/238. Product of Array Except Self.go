package leetcode

func ProductExceptSelf(nums []int) []int {
	multip1 := make([]int, len(nums))
	multip2 := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			multip1[i] = nums[i]
		} else {
			multip1[i] = multip1[i-1] * nums[i]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			multip2[i] = nums[i]
		} else {
			multip2[i] = multip2[i+1] * nums[i]
		}
	}
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tem := 1
		if i > 0 {
			tem = tem * multip1[i-1]
		}
		if i < len(nums)-1 {
			tem = tem * multip2[i+1]
		}
		result[i] = tem
	}
	return result
}
