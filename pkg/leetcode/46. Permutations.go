package leetcode

func permute(nums []int) [][]int {
	traversal := make([]int, 0)
	result := make([][]int, 0)
	return permuteBase(nums, traversal, result)
}

func permuteBase(nums []int, traversal []int, result [][]int) [][]int {
	if len(traversal) == len(nums) {
		return append(result, traversal)
	}
	for _, value := range nums {
		if !existNum(traversal, value) {
			result = permuteBase(nums, append(traversal, value), result)
		}
	}
	return result
}

func existNum(tem []int, target int) bool {
	for _, value := range tem {
		if value == target {
			return true
		}
	}
	return false
}
