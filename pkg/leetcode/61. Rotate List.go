package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length := 0
	point1 := head
	for point1 != nil {
		length += 1
		point1 = point1.Next
	}
	if k%length == 0 {
		return head
	}
	point1 = head
	for point1.Next != nil {
		point1 = point1.Next
	}
	point1.Next = head
	k = length - k%length
	point1 = head
	for i := k - 1; i > 0; i-- {
		if i == 1 {
			tem := point1
			point1 = point1.Next
			tem.Next = nil
			continue
		}
		point1 = point1.Next
	}
	return head
}
