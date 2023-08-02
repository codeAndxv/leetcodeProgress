package leetcode

/**
 * Definition for singly-linked list.

 */
func oddEvenList(head *ListNode) *ListNode {
	tem := head
	point1 := head
	point2 := head.Next
	num := 0
	for point2 != nil && tem.Next != nil {
		if num%2 == 0 {
			tem.Next =
		}
		num++
		tem = tem.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
