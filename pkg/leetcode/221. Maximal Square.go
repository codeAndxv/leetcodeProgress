package leetcode

func MaximalSquare(matrix [][]byte) int {
	var squares []Square
	for index1, value1 := range matrix {
		for index2, value2 := range value1 {
			if value2 == '1' {
				squares = append(squares, Square{
					Left: index2,
					Top:  index1,
					Edge: 1,
				})
			}
		}
	}
	if len(squares) == len(matrix)*len(matrix[0]) {
		return min(len(matrix), len(matrix[0])) * min(len(matrix), len(matrix[0]))
	}
	for index, square := range squares {
		squares[index] = burst(square, matrix)
	}
	maxArea := 0
	for _, square := range squares {
		maxArea = max(maxArea, square.Edge*square.Edge)
	}
	return maxArea
}

func burst(square Square, matrix [][]byte) Square {
	step := 0
	exit := false
	for !exit && step+1+square.Left < len(matrix[0]) && step+1+square.Top < len(matrix) {
		for tem1 := square.Left; tem1 <= square.Left+step+1; tem1++ {
			if matrix[square.Top+step+1][tem1] == '0' {
				exit = true
				break
			}
		}
		if exit {
			break
		}

		for tem2 := square.Top; tem2 <= square.Top+step+1; tem2++ {
			if matrix[tem2][square.Left+step+1] == '0' {
				exit = true
				break
			}
		}
		if exit {
			break
		}
		step += 1
	}
	square.Edge += step
	return square
}

type Square struct {
	Left int
	Top  int
	Edge int
}

/*func maximalSquareBase(matrix [][]byte, row int, column int) int {
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
}*/

/*
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
*/
