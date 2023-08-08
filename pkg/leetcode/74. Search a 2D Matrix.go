package leetcode

func searchMatrix2(matrix [][]int, target int) bool {
	low := 0
	high := len(matrix)*len(matrix[0]) - 1
	for low < high {
		mid := (low + high) / 2
		a, b := parseIndex(mid, len(matrix[0]))
		if matrix[a][b] == target {
			return true
		} else if matrix[a][b] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	a, b := parseIndex(low, len(matrix[0]))
	return matrix[a][b] == target
}

func parseIndex(index int, len int) (int, int) {
	return index / len, index % len
}
