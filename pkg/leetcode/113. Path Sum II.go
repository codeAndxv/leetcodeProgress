package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*func pathSum(root *TreeNode, targetSum int) [][]int {
	currentPath := make([]int, 5000)
	result := &[][]int{}
	pathSumBase(root, 0, currentPath, targetSum, result)
	return *result
}

func pathSumBase(root *TreeNode, currentSum int, currentPath *[]int, targetSum int, result *[][]int) {
	if root == nil {
		return
	}
	*currentPath = append(*currentPath, root.Val)
	currentSum += root.Val
	if root.Left == nil && root.Right == nil && currentSum == targetSum {
		*result = append(*result, append(*currentPath, root.Val))
	}
	if root.Left != nil {
		pathSumBase(root.Left, currentSum, currentPath, targetSum, result)
	}
	if root.Right != nil {
		pathSumBase(root.Right, currentSum, currentPath, targetSum, result)
	}
	*currentPath = &(*currentPath[:len(*currentPath)-1])
}
*/
