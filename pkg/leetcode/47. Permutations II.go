package leetcode

func permuteUnique(nums []int) [][]int {
	traversal := make([]int, 0)
	result := make([][]int, 0)
	flag := make([]int, 21)
	for _, value := range nums {
		flag[value+10] += 1
	}
	return permuteUniqueBase(nums, traversal, flag, result)
}

func permuteUniqueBase(nums []int, traversal []int, flag []int, result [][]int) [][]int {
	if len(traversal) == len(nums) {
		return append(result, traversal)
	}
	for index, value := range flag {
		if value > 0 {
			flag[index] -= 1
			result = permuteUniqueBase(nums, append(traversal, index-10), flag, result)
			flag[index] += 1
		}
	}
	return result
}
