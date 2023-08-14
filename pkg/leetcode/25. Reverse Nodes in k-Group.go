package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	nodes := make([]*ListNode, 0)
	point1 := head
	count := 0
	start := 0
	for point1 != nil {
		count++
		nodes = append(nodes, point1)
		if count == k {
			for i := 0; i < k/2; i++ {
				nodes[start+i], nodes[start+k-1-i] = nodes[start+k-1-i], nodes[i+start]
			}
			count = 0
			start = start + k
		}
		point1 = point1.Next
	}
	var temHead, result *ListNode
	for i, v := range nodes {
		if temHead == nil {
			temHead = v
			result = v
			continue
		}
		temHead.Next = v
		temHead = temHead.Next
		if i == len(nodes)-1 {
			temHead.Next = nil
		}
	}
	return result
}
