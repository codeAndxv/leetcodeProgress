package leetcode

import (
	"math"
)

func goodNodes(root *TreeNode) int {
	return goodNodesBase(root, math.MinInt32)
}

func goodNodesBase(root *TreeNode, maxValue int) int {
	if root == nil {
		return 0
	}

	current := 0
	if root.Val >= maxValue {
		current = 1
	}
	left := goodNodesBase(root.Left, max(root.Val, maxValue))
	right := goodNodesBase(root.Right, max(root.Val, maxValue))
	return current + left + right
}
