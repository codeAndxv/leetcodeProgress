package leetcode

func doubleIt(head *ListNode) *ListNode {
	nodes := make([]*ListNode, 0)
	point1 := head
	for point1 != nil {
		nodes = append(nodes, point1)
		point1 = point1.Next
	}
	var afterHead *ListNode
	lastCarry := 0
	for i := len(nodes) - 1; i >= 0; i-- {
		tem := &ListNode{
			Val:  (nodes[i].Val*2 + lastCarry) % 10,
			Next: afterHead,
		}
		afterHead = tem
		lastCarry = (nodes[i].Val*2 + lastCarry) / 10
	}
	if lastCarry != 0 {
		tem := &ListNode{
			Val:  lastCarry,
			Next: afterHead,
		}
		afterHead = tem
	}
	return afterHead
}

/**
if head.Val == 0 {
	return head
}
point1 := head
originValue := 0
for point1 != nil {
	originValue = originValue*10 + point1.Val
	point1 = point1.Next
}
afterValue := originValue * 2
var afterHead *ListNode
for afterValue != 0 {
	tem := &ListNode{
		Val:  afterValue % 10,
		Next: afterHead,
	}
	afterHead = tem
	afterValue = afterValue / 10
}
return afterHead
*/
