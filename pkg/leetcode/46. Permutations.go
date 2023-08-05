package leetcode

func permute(nums []int) [][]int {
	traversal := make([]int, 0)
	result := make([][]int, 0)
	flag := make([]int, 21)
	for _, value := range nums {
		flag[value+10] = 1
	}
	return permuteBase(nums, traversal, flag, result)
}

func permuteBase(nums []int, traversal []int, flag []int, result [][]int) [][]int {
	if len(traversal) == len(nums) {
		return append(result, traversal)
	}
	for index, value := range flag {
		if value > 0 {
			flag[index] -= 1
			result = permuteBase(nums, append(traversal, index-10), flag, result)
			flag[index] += 1
		}
	}
	return result
}
