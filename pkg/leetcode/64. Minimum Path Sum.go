package leetcode

import "math"

func minPathSum(grid [][]int) int {
	var matrix [][]int
	for i := 0; i < len(grid); i++ {
		matrix = append(matrix, make([]int, len(grid[i])))
		for j := 0; j < len(grid[i]); j++ {
			matrix[i][j] = math.MaxInt32
		}
	}
	return minPathSumBase(grid, 0, 0, matrix)
}

func minPathSumBase(grid [][]int, row int, column int, matrix [][]int) int {
	if matrix[row][column] != math.MaxInt32 {
		return matrix[row][column]
	}
	if row == len(grid)-1 && column == len(grid[0])-1 {
		matrix[row][column] = grid[row][column]
		return matrix[row][column]
	}
	if row == len(grid)-1 {
		matrix[row][column] = minPathSumBase(grid, row, column+1, matrix) + grid[row][column]
		return matrix[row][column]
	}
	if column == len(grid[0])-1 {
		matrix[row][column] = minPathSumBase(grid, row+1, column, matrix) + grid[row][column]
		return matrix[row][column]
	}
	matrix[row][column] = min(minPathSumBase(grid, row+1, column, matrix), minPathSumBase(grid, row, column+1, matrix)) + grid[row][column]
	return matrix[row][column]
}
