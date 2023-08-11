package leetcode

func kthSmallest(root *TreeNode, k int) int {
	var target int
	count := 0
	kthSmallestBase(root, &count, k, &target)
	return target
}

func kthSmallestBase(root *TreeNode, count *int, k int, target *int) {
	if root == nil {
		return
	}
	kthSmallestBase(root.Left, count, k, target)
	*count = *count + 1
	if *count == k {
		*target = root.Val
		return
	}
	kthSmallestBase(root.Right, count, k, target)
}
