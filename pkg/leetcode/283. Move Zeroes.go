package leetcode

func moveZeroes(nums []int) {
	left := 0
	right := 0
	for left < len(nums) && right < len(nums) {
		if nums[left] != 0 {
			left++
			right++
			continue
		}
		if nums[right] == 0 {
			right++
			continue
		}
		tem := nums[left]
		nums[left] = nums[right]
		nums[right] = tem
	}
}
