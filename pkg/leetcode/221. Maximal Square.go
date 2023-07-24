package leetcode

import "math"

func maximalSquare(matrix [][]byte) int {

	return maximalSquareBase(matrix, 0, 0)
}

func maximalSquareBase(matrix [][]byte, row int, column int) int {
	if row == len(matrix)-1 && column == len(matrix[0])-1 {
		return int(matrix[row][column])
	}
	if matrix[row][column] == 0 {
		// find first not equal 0
		for tem1 := row + 1; tem1 < len(matrix)-1 && matrix[tem1]; tem1++ {

		}
	}
	if row == len(matrix)-1 {
		return max(maximalSquareBase(matrix, row, column+1), int(matrix[row][column]))
	}
	if column == len(matrix[0])-1 {
		return max(maximalSquareBase(matrix, row+1, column), int(matrix[row][column]))
	}
	return max(maximalSquareBase(matrix, row+1, column), maximalSquareBase(matrix, row, column+1))
}

func tem() int {
	var sections []Section
	if row == len(matrix)-1 {
		tem1 := math.MaxInt32
		for index, value := range matrix[row] {
			for tem1 == math.MaxInt32 {
				if value == 1 {
					tem1 = index
					break
				}
			}
			if value != 1 {
				sections = append(sections, Section{
					Start: tem1,
					End:   index - 1,
				})
				tem1 = math.MaxInt32
			}
		}
		return sections
	}
}

type Section struct {
	Start int
	End   int
}
