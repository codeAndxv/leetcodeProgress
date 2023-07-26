package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	var temSlice []int
	result := &[][]int{}
	pathSumBase(root, 0, temSlice, targetSum, result)
	return *result
}

func pathSumBase(root *TreeNode, currentSum int, currentPath []int, targetSum int, result *[][]int) {
	if root == nil {
		return
	}
	if root != nil && root.Left == nil && root.Right == nil && currentSum+root.Val == targetSum {
		*result = append(*result, append(currentPath, root.Val))
	}
	if root.Left != nil {
		pathSumBase(root.Left, currentSum+root.Val, append(currentPath, root.Val), targetSum, result)
	}

	if root.Right != nil {
		pathSumBase(root.Right, currentSum+root.Val, append(currentPath, root.Val), targetSum, result)
	}
}

/*type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
*/
