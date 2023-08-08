package leetcode

func sortColors(nums []int) {
	counts := make([]int, 3)
	for _, v := range nums {
		counts[v] += 1
	}
	for i := 0; i < len(nums); i++ {
		if i < counts[0] {
			nums[i] = 0
		} else if i < counts[0]+counts[1] {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
