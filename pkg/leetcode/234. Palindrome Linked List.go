package leetcode

func isPalindrome(head *ListNode) bool {
	nodes := make([]*ListNode, 0)
	point1 := head
	for point1 != nil {
		nodes = append(nodes, point1)
		point1 = point1.Next
	}
	for i := 0; i < len(nodes)/2; i++ {
		if nodes[i].Val != nodes[len(nodes)-1-i].Val {
			return false
		}
	}
	return true
}
