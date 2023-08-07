package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	buildTreeBase(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func buildTreeBase(preorder []int, inorder []int,
	preorderStart int, preorderEnd int, inorderStart int, inorderEnd int) *TreeNode {
	if pre
	root := TreeNode{
		Val:   preorder[preorderStart],
		Left:  nil,
		Right: nil,
	}

	if preorderEnd > preorderStart {
		root.Left = buildTreeBase(preorder, inorder, inorderStart, )
	}
	return &root
}

func searchIndex(nums []int, target int) int {
	for index, value := range nums {
		if value == target {
			return index
		}
	}
	return -1
}
