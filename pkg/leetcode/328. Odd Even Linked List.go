package leetcode

/*
*
  - Definition for singly-linked list.
*/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tem := head
	var point1, point1tem, point2, point2tem *ListNode
	num := 0
	for tem != nil {
		if num%2 == 0 {
			if point1 == nil {
				point1 = tem
				point1tem = tem
			} else {
				point1tem.Next = tem
				point1tem = point1tem.Next
			}
		}
		if num%2 == 1 {
			if point2 == nil {
				point2 = tem
				point2tem = tem
			} else {
				point2tem.Next = tem
				point2tem = point2tem.Next
			}
		}
		num++
		tem = tem.Next
	}
	if num%2 == 1 {
		point2tem.Next = nil
	}
	point1tem.Next = point2
	return point1
}

type ListNode struct {
	Val  int
	Next *ListNode
}
