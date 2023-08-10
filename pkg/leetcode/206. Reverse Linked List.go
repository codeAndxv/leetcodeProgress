package leetcode

func reverseList(head *ListNode) *ListNode {
	nodeLists := make([]*ListNode, 0)
	point1 := head
	for point1 != nil {
		nodeLists = append(nodeLists, point1)
		point1 = point1.Next
	}

	var revHead, tem *ListNode
	for j := len(nodeLists) - 1; j >= 0; j-- {
		if tem == nil {
			tem = nodeLists[j]
			revHead = nodeLists[j]
			continue
		}

		tem.Next = nodeLists[j]
		tem = tem.Next
		if j == 0 {
			tem.Next = nil
		}
	}
	return revHead
}
