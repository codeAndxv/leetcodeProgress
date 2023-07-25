package leetcode

/*
*
  - Definition for a binary tree node.
*/
/*func rob(root *TreeNode) int {
	return robBase(root)
}

func robBase(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	score1 := 0
	if root.Left != nil {
		score1 += robBase(root.Left)
	}
	if root.Right != nil {
		score1 += robBase(root.Right)
	}
	score2 := root.Val
	if root.Left != nil {
		if root.Left.Left != nil {
			score2 += robBase(root.Left.Left)
		}
		if root.Left.Right != nil {
			score2 += robBase(root.Left.Right)
		}
	}
	if root.Right != nil {
		if root.Right.Left != nil {
			score2 += robBase(root.Right.Left)
		}
		if root.Right.Right != nil {
			score2 += robBase(root.Right.Right)
		}
	}
	return max(score1, score2)
}
*/
/*
func rob(root *TreeNode) int {
	scoreMap := make(map[*TreeNode]int)
	return robBase(root, scoreMap)
}

func robBase(root *TreeNode, scoreMap map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if value, exist := scoreMap[root]; exist {
		return value
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	score1 := 0
	if root.Left != nil {
		if _, exist := scoreMap[root.Left]; !exist {
			scoreMap[root.Left] = robBase(root.Left, scoreMap)
		}
		score1 += scoreMap[root.Left]
	}
	if root.Right != nil {
		if _, exist := scoreMap[root.Right]; !exist {
			scoreMap[root.Right] = robBase(root.Right, scoreMap)
		}
		score1 += scoreMap[root.Right]
	}
	score2 := root.Val
	if root.Left != nil {
		if root.Left.Left != nil {
			if _, exist := scoreMap[root.Left.Left]; !exist {
				scoreMap[root.Left.Left] = robBase(root.Left.Left, scoreMap)
			}
			score2 += scoreMap[root.Left.Left]
		}

		if root.Left.Right != nil {
			if _, exist := scoreMap[root.Left.Right]; !exist {
				scoreMap[root.Left.Right] = robBase(root.Left.Right, scoreMap)
			}
			score2 += scoreMap[root.Left.Right]
		}

	}
	if root.Right != nil {
		if root.Right.Left != nil {
			if _, exist := scoreMap[root.Right.Left]; !exist {
				scoreMap[root.Right.Left] = robBase(root.Right.Left, scoreMap)
			}
			score2 += scoreMap[root.Right.Left]
		}

		if root.Right.Right != nil {
			if _, exist := scoreMap[root.Right.Right]; !exist {
				scoreMap[root.Right.Right] = robBase(root.Right.Right, scoreMap)
			}
			score2 += scoreMap[root.Right.Right]
		}

	}

	scoreMap[root] = max(score1, score2)
	return scoreMap[root]
}*/

/**
func rob(root *TreeNode) int {
    return max(robNode(root))
}

func robNode(n *TreeNode) (rob, skip int) {
    if n == nil { return 0, 0 }

    lRob, lSkip := robNode(n.Left)
    rRob, rSkip := robNode(n.Right)

    rob = n.Val + lSkip + rSkip
    skip = max(lRob, lSkip) + max(rRob, rSkip)
    return
}

func max(N ...int) int {
    r := N[0]
    for _, n := range N {
        if n > r { r = n }
    }
    return r
}
*/
