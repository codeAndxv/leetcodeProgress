package leetcode

func diameterOfBinaryTree(root *TreeNode) int {
	var maxValue int
	diameterOfBinaryTreeBase(root, &maxValue)
	return maxValue
}

func diameterOfBinaryTreeBase(root *TreeNode, maxValue *int) int {
	if root == nil {
		return 0
	}
	left := diameterOfBinaryTreeBase(root.Left, maxValue)
	right := diameterOfBinaryTreeBase(root.Right, maxValue)
	*maxValue = max(*maxValue, left+right)
	return max(left, right)
}

/**
func diameterOfBinaryTree(root *TreeNode) int {
	_, tem2 := diameterOfBinaryTreeBase(root, 0)
	return tem2
}

func diameterOfBinaryTreeBase(root *TreeNode, maxLen int) (int, int) {
	if root == nil {
		return 0, maxLen
	}
	left, maxLenL := diameterOfBinaryTreeBase(root.Left, maxLen)
	right, maxLenR := diameterOfBinaryTreeBase(root.Right, maxLen)
	return max(left, right) + 1, max(max(max(left+right+1, maxLen), maxLenL), maxLenR)
}

*/
