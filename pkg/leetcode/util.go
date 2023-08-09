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
