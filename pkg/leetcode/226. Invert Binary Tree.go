package leetcode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	tem := root.Right
	root.Right = root.Left
	root.Left = tem
	return root
}
