package leetcode

func isSymmetric(root *TreeNode) bool {
	return isSymmetricBase(root.Left, root.Right)
}

func isSymmetricBase(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	return left.Val == right.Val && isSymmetricBase(left.Left, right.Right) && isSymmetricBase(left.Right, right.Left)
}
