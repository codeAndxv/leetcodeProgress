package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	point1 := headA
	len1 := 0
	point2 := headB
	len2 := 0
	for point1 != nil {
		len1++
		point1 = point1.Next
	}
	for point2 != nil {
		len2++
		point2 = point2.Next
	}
	point1 = headA
	point2 = headB
	if len2 > len1 {
		for i := 0; i < len2-len1; i++ {
			point2 = point2.Next
		}
	} else {
		for i := 0; i < len1-len2; i++ {
			point1 = point1.Next
		}
	}
	for point1 != nil && point2 != nil {
		if point1 == point2 {
			return point1
		}
		point1 = point1.Next
		point2 = point2.Next
	}
	return nil
}
