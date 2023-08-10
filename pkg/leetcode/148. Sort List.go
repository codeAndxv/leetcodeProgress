package leetcode

func sortList(head *ListNode) *ListNode {
	nodes := make([]*ListNode, 0)
	point1 := head
	for point1 != nil {
		nodes = append(nodes, point1)
		point1 = point1.Next
	}
	heapSortNode(nodes)
	var resHead, tem *ListNode
	for i := 0; i < len(nodes); i++ {
		if resHead == nil {
			resHead = nodes[i]
			tem = nodes[i]
		} else {
			tem.Next = nodes[i]
			tem = tem.Next
		}
		if i == len(nodes)-1 {
			tem.Next = nil
		}
	}
	return resHead
}

func heapifyNode(arr []*ListNode, n int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left].Val > arr[largest].Val {
		largest = left
	}

	if right < n && arr[right].Val > arr[largest].Val {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapifyNode(arr, n, largest)
	}
}

func heapSortNode(arr []*ListNode) {
	n := len(arr)

	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapifyNode(arr, n, i)
	}

	// 逐步将最大元素移到数组末尾
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 将当前最大元素移到末尾
		heapifyNode(arr, i, 0)          // 重新构建堆
	}
}
