package leetcode

import "math"

func minimumTotal(triangle [][]int) int {
	var matrix [][]int
	for i := 0; i < len(triangle); i++ {
		matrix = append(matrix, make([]int, i+1))
		for j := 0; j < i+1; j++ {
			matrix[i][j] = math.MaxInt32
		}
	}
	return minimumTotalBase(triangle, 0, 0, matrix)
}

func minimumTotalBase(triangle [][]int, layer int, index int, matrix [][]int) int {
	if layer == len(triangle)-1 {
		matrix[layer][index] = triangle[layer][index]
		return triangle[layer][index]
	}
	var left int
	var right int
	if matrix[layer+1][index] != math.MaxInt32 {
		left = matrix[layer+1][index]
	} else {
		matrix[layer+1][index] = minimumTotalBase(triangle, layer+1, index, matrix)
		left = minimumTotalBase(triangle, layer+1, index, matrix)
	}
	if index+1 < len(triangle[layer+1]) {
		if matrix[layer+1][index+1] != math.MaxInt32 {
			right = matrix[layer+1][index+1]
		} else {
			matrix[layer+1][index+1] = minimumTotalBase(triangle, layer+1, index+1, matrix)
			right = minimumTotalBase(triangle, layer+1, index+1, matrix)
		}
	}
	return min(left, right) + triangle[layer][index]
}
