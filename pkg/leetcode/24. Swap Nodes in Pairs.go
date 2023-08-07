package leetcode

func swapPairs(head *ListNode) *ListNode {
	trueHead := &ListNode{
		Val:  0,
		Next: head,
	}
	temHead := trueHead
	for temHead != nil && temHead.Next != nil && temHead.Next.Next != nil {
		tem := temHead.Next.Next
		temHead.Next.Next = temHead.Next.Next.Next
		tem.Next = temHead.Next
		temHead.Next = tem
		temHead = temHead.Next.Next
	}
	return trueHead.Next
}
