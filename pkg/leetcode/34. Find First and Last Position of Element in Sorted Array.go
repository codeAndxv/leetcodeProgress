package leetcode

func searchRange(nums []int, target int) []int {
	start := -1
	end := -1
	for index, value := range nums {
		if value == target {
			if start == -1 {
				start = index
				end = index
			} else {
				end = index
			}
		}
	}
	return []int{start, end}
}
