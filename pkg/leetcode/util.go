package leetcode

import "math"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func pow(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sortN2(source []int) []int {
	if len(source) < 2 {
		return source
	}
	for i := 0; i < len(source); i++ {
		for j := i + 1; j < len(source); j++ {
			if source[i] > source[j] {
				source[i], source[j] = source[j], source[i]
			}
		}
	}
	return source
}

func quickSort(arr []int) {
	quickSortBase(arr, 0, len(arr))
}

func quickSortBase(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSortBase(arr, low, pivot-1)
		quickSortBase(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// 调整堆，使其满足堆的性质
func heapify(arr []int, n int, i int) {
	largest := i     // 初始化父节点为最大值
	left := 2*i + 1  // 左子节点索引
	right := 2*i + 2 // 右子节点索引

	// 如果左子节点比父节点大，则更新最大值索引
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 如果右子节点比当前最大值大，则更新最大值索引
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是父节点，则交换父节点和最大值位置，并递归调整交换后的子树
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
	}
}

// 堆排序主函数
func heapSort(arr []int, k int) {
	n := len(arr)

	// 从最后一个非叶子节点开始构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 一个个从堆中取出元素
	for i := n - 1; i >= n-k; i-- {
		// 将当前堆顶元素（最大值）与未排序部分的最后一个元素交换
		arr[0], arr[i] = arr[i], arr[0]

		// 重新构建最大堆
		heapify(arr, i, 0)
	}
}
