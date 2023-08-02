package leetcode

func SearchMatrix(matrix [][]int, target int) bool {
	tem := make([]int, max(len(matrix), len(matrix[0])))
	for i, _ := range tem {
		if len(matrix) >= len(matrix[0]) {
			if i < len(matrix[0]) {
				tem[i] = matrix[i][i]
			} else {
				tem[i] = matrix[i][len(matrix[0])-1]
			}
		} else {
			if i < len(matrix) {
				tem[i] = matrix[i][i]
			} else {
				tem[i] = matrix[len(matrix)-1][i]
			}
		}
	}
	index, find := binarySearch(tem, target)
	if find {
		return true
	}
	if index < 0 || index == max(len(matrix), len(matrix[0])) {
		return false
	}
	if index+1 < len(matrix) {
		for i := 0; i < min(len(matrix[0]), index+1); i++ {
			if matrix[min(index, len(matrix)-1)][i] == target {
				return true
			}
			if matrix[min(index, len(matrix)-1)][i] > target {
				break
			}
		}
	}
	if index+1 < len(matrix[0]) {
		for i := index + 1; i < min(len(matrix), index+1); i++ {
			if matrix[i][min(index, len(matrix[0])-1)] == target {
				return true
			}
			if matrix[i][min(index, len(matrix[0])-1)] > target {
				break
			}
		}
	}
	return false
}

func binarySearchRecursive(source []int, target, low, high int) (index int, find bool) {
	if low > high {
		return high, false
	}

	mid := (low + high) / 2

	if source[mid] == target {
		return mid, true
	} else if source[mid] < target {
		return binarySearchRecursive(source, target, mid+1, high)
	} else {
		return binarySearchRecursive(source, target, low, mid-1)
	}
}

// binarySearch 是对外暴露的递归二分搜索函数
func binarySearch(source []int, target int) (index int, find bool) {
	return binarySearchRecursive(source, target, 0, len(source)-1)
}

/**
tem := make([]int, len(matrix))
	for i, _ := range matrix {
		tem[i] = matrix[i][0]
	}
	index, find := binarySearch(tem, target)
	if find {
		return true
	}
	if index < 0 {
		return false
	}
	index, find = binarySearch(matrix[index], target)
	if find {
		return true
	}
	if index < 0 {
		return false
	}
	index, find = binarySearch(matrix[0], target)
	if find {
		return true
	}
	if index < 0 {
		return false
	}
	for i, _ := range matrix {
		tem[i] = matrix[i][index]
	}
	index, find = binarySearch(tem, target)
	if find {
		return true
	}
	if index < 0 {
		return false
	}
	return false


*/

/**
Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:
编写一个在 m x n 整数矩阵 matrix 中搜索值 target 的有效算法。此矩阵具有以下属性：

Integers in each row are sorted in ascending from left to right.
每行中的整数按从左到右的升序排序。
Integers in each column are sorted in ascending from top to bottom.
每列中的整数按从上到下的升序排序。


Example 1: 示例 1：


Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
Output: true
Example 2: 示例 2：


Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
Output: false


Constraints: 约束：

m == matrix.length
n == matrix[i].length
1 <= n, m <= 300
-109 <= matrix[i][j] <= 109
All the integers in each row are sorted in ascending order.
每行中的所有整数都按升序排序。
All the integers in each column are sorted in ascending order.
每列中的所有整数按升序排序。
-109 <= target <= 109
*/
