package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temHead := &ListNode{
		Val:  0,
		Next: head,
	}
	point1 := head
	length := 0
	for point1 != nil {
		length++
		point1 = point1.Next
	}
	if n == length {
		return head.Next
	} else {
		for i := length - n; i >= 1; i-- {
			if i == 1 {
				temHead.Next.Next = temHead.Next.Next.Next
			} else {
				temHead.Next = temHead.Next.Next
			}
		}
	}
	return head
}
