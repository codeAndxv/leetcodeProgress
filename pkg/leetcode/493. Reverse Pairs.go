package leetcode

func reversePairs(nums []int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > 2*nums[j] {
				num += 1
			}
		}
	}
	return num
}
