package leetcode

func inorderTraversal(root *TreeNode) []int {
	travResult := make([]int, 0)
	return inorderTraversalBase(root, travResult)
}

func inorderTraversalBase(root *TreeNode, travResult []int) []int {
	if root == nil {
		return travResult
	}
	if root.Left != nil {
		travResult = inorderTraversalBase(root.Left, travResult)
	}
	travResult = append(travResult, root.Val)
	if root.Right != nil {
		travResult = inorderTraversalBase(root.Right, travResult)
	}
	return travResult
}
