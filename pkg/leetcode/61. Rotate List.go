package leetcode

import "fmt"

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
	k = (length - k%length) % length
	fmt.Println(k)
	if k == 0 || k == length {
		return head
	}
	fmt.Println(k)
	tem2 := head
	for i := 0; i < length; i++ {
		if i == length-1 {
			tem2.Next = head
			break
		}
		tem2 = tem2.Next
	}
	tem := &ListNode{
		Val:  0,
		Next: head,
	}
	for i := 0; i < k; i++ {
		if i == k-1 {
			tem1 := tem.Next.Next
			tem.Next.Next = nil
			tem.Next = tem1
			break
		}
		tem.Next = tem.Next.Next
	}
	return tem.Next
}
