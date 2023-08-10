package leetcode

func rotate2(nums []int, k int) {
	tem := make([]int, 0)
	k = k % len(nums)
	tem = append(tem, nums[len(nums)-k:]...)
	tem = append(tem, nums[:len(nums)-k]...)
	for i := 0; i < len(nums); i++ {
		nums[i] = tem[i]
	}
}
