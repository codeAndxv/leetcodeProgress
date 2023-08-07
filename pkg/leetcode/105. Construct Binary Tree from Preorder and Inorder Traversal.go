package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeBase(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func buildTreeBase(preorder []int, inorder []int,
	preorderStart int, preorderEnd int, inorderStart int, inorderEnd int) *TreeNode {
	root := TreeNode{
		Val:   preorder[preorderStart],
		Left:  nil,
		Right: nil,
	}
	index := searchIndex(inorder, preorder[preorderStart])
	leftLen := index - inorderStart
	rightLen := inorderEnd - index
	if leftLen > 0 {
		root.Left = buildTreeBase(preorder, inorder, preorderStart+1, preorderStart+leftLen, inorderStart, index-1)
	}
	if rightLen > 0 {
		root.Right = buildTreeBase(preorder, inorder, preorderStart+leftLen+1, preorderEnd, index+1, inorderEnd)
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
