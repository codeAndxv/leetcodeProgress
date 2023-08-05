package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head1, head2 *ListNode
	var tem1, tem2 *ListNode
	if list1.Val > list2.Val {
		head1 = list2
		head2 = list1
	} else {
		head1 = list1
		head2 = list2
	}
	tem1 = head1
	tem2 = head2
	for tem2 != nil {
		for tem1.Val <= tem2.Val {
			if tem1.Next == nil || tem1.Next.Val > tem2.Val {
				break
			}
			tem1 = tem1.Next
		}
		tem3 := tem2
		tem2 = tem2.Next
		tem3.Next = tem1.Next
		tem1.Next = tem3
	}
	return head1
}
