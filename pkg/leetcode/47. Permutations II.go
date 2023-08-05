package leetcode

func permuteUnique(nums []int) [][]int {

}

func permuteUniqueBase(nums []int, traversal []int, flag []int, result [][]int) [][]int {
	if len(traversal) == len(nums) {
		return append(result, traversal)
	}
	for _, value := range nums {
		if !existNum(traversal, value) {
			result = permuteUniqueBase(nums, append(traversal, value), result)
		}
	}
	return result
}
