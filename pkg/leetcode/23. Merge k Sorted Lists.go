package leetcode

import (
	"fmt"
	"strconv"
)

func mergeKLists(lists []*ListNode) *ListNode {
	heapNodes := make([]*ListNode, 0)
	for _, v := range lists {
		if v != nil {
			heapNodes = append(heapNodes, v)
		}
	}
	heapSortListNode(heapNodes)
	var head, tem *ListNode
	for len(heapNodes) > 0 {
		n := len(heapNodes)
		heapifyListNode(heapNodes, n, 0)
		logHeap(heapNodes)
		minNode := heapNodes[0]
		if head == nil {
			head = minNode
			tem = minNode
		} else {
			tem.Next = minNode
			tem = tem.Next
		}
		if minNode.Next != nil {
			heapNodes[0] = minNode.Next
		} else {
			heapNodes[0], heapNodes[n-1] = heapNodes[n-1], heapNodes[0]
			heapNodes = heapNodes[:len(heapNodes)-1]
		}
	}
	return head
}

func logHeap(heapNodes []*ListNode) {
	for i, v := range heapNodes {
		fmt.Println(strconv.Itoa(i) + `_` + strconv.Itoa(v.Val))
	}
}

// 调整堆，使其满足堆的性质
func heapifyListNode(arr []*ListNode, n int, i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left].Val < arr[smallest].Val {
		smallest = left
	}

	if right < n && arr[right].Val < arr[smallest].Val {
		smallest = right
	}

	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		heapifyListNode(arr, n, smallest)
	}
}

func heapSortListNode(arr []*ListNode) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapifyListNode(arr, n, i)
	}
	/*
		// 一个个从堆中取出元素
		for i := n - 1; i >= n-k; i-- {
			// 将当前堆顶元素（最大值）与未排序部分的最后一个元素交换
			arr[0], arr[n - 1] = arr[n - 1], arr[0]

			// 重新构建最大堆
			heapify(arr, i, 0)
		}*/
}

/**
var head, tem *ListNode
flag := true
for flag {
	minValue := math.MaxInt32
	var minIndex int
	flag = false
	for i, v := range lists {
		if v != nil {
			flag = true
			if minValue > v.Val {
				minValue = v.Val
				minIndex = i
			}
		}
	}
	if flag {
		if head == nil {
			head = lists[minIndex]
			tem = head

		} else {
			tem.Next = lists[minIndex]
			tem = tem.Next
		}
		lists[minIndex] = lists[minIndex].Next
	}
}
return head
*/

// 时间复杂度  O(n*logk + k)
