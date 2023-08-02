package leetcode

func missingNumber(nums []int) int {
	tem := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tem[i] = i + 1
	}
	for _, value := range nums {
		if value > 0 {
			tem[value-1] = 0
		}
	}
	for i := 0; i < len(nums); i++ {
		if tem[i] != 0 {
			return tem[i]
		}
	}
	return 0
}
