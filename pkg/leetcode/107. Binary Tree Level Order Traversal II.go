package leetcode

func levelOrderBottom(root *TreeNode) [][]int {
	result := levelOrder(root)
	low := 0
	high := len(result) - 1
	for low < high {
		tem := result[low]
		result[low] = result[high]
		result[high] = tem
		low++
		high--
	}
	return result
}
