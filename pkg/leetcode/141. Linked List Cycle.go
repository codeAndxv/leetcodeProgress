package leetcode

func hasCycle(head *ListNode) bool {
	nodes := make(map[*ListNode]bool)
	point1 := head
	for point1 != nil {
		if nodes[point1] {
			return true
		}
		nodes[point1] = true
		point1 = point1.Next
	}
	return false
}
