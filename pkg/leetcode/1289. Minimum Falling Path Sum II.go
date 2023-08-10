package leetcode

import "math"

func minFallingPathSum(grid [][]int) int {
	records := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		records[i] = make([]int, len(grid))
	}
	for i := len(grid) - 1; i >= 0; i-- {
		if i == len(grid)-1 {
			for j := 0; j < len(grid); j++ {
				records[i][j] = grid[i][j]
			}
		} else {
			min1, index1, min2, _ := getMinInSlice(records[i+1])
			for j := 0; j < len(grid); j++ {
				if index1 != j {
					records[i][j] = min1 + grid[i][j]
				} else {
					records[i][j] = min2 + grid[i][j]
				}
			}
		}
	}
	min, _, _, _ := getMinInSlice(records[0])
	return min
}

func getMinInSlice(arr []int) (int, int, int, int) {
	minValue := math.MaxInt32
	index1 := -1
	minValue2 := math.MaxInt32
	index2 := -1
	for i, v := range arr {
		if v < minValue {
			minValue2 = minValue
			index2 = index1
			minValue = v
			index1 = i
		} else if v < minValue2 {
			minValue2 = v
			index2 = i
		}
	}
	return minValue, index1, minValue2, index2
}

/**

func minFallingPathSum(grid [][]int) int {
	records := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		records[i] = make([]int, len(grid))
	}
	for i := len(grid) - 1; i >= 0; i-- {
		for j := 0; j < len(grid); j++ {
			if i == len(grid)-1 {
				records[i][j] = grid[i][j]
			} else {
				records[i][j] = getMinInSlice(records[i+1], j) + grid[i][j]
			}
		}
	}
	return getMinInSlice2(records[0])
}

func getMinInSlice2(arr []int) int {
	minValue := math.MaxInt32
	for _, v := range arr {
		minValue = min(minValue, v)
	}
	return minValue
}

func getMinInSlice(arr []int, outIndex int) int {
	minValue := math.MaxInt32
	for i, v := range arr {
		if i != outIndex {
			minValue = min(minValue, v)
		}
	}
	return minValue
}
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
*/
