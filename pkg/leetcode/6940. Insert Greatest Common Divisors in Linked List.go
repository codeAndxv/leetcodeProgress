package leetcode

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	point1 := head
	for point1.Next != nil {

		tem := &ListNode{
			Val:  gcd(point1.Val, point1.Next.Val),
			Next: nil,
		}
		tem2 := point1.Next
		point1.Next = tem
		tem.Next = tem2
		point1 = point1.Next.Next
	}
	return head
}

// 求最大公约数的函数
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
