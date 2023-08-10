package leetcode

func detectCycle(head *ListNode) *ListNode {
	nodes := make(map[*ListNode]bool)
	point1 := head
	for point1 != nil {
		if nodes[point1] {
			return point1
		}
		nodes[point1] = true
		point1 = point1.Next
	}
	return nil
}

//  https://leetcode.cn/problems/linked-list-cycle-ii/solutions/12616/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/?envType=study-plan-v2&envId=top-100-liked
