package leetcode

func setZeroes(matrix [][]int) {
	flag := make([]int, len(matrix[0]))
	for _, value := range matrix {
		for i := 0; i < len(flag); i++ {
			if value[i] == 0 {
				flag[i] = 1
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		setZe := false
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				setZe = true
				break
			}
		}
		if setZe {
			for j := 0; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for index, value := range flag {
		if value == 1 {
			for i := 0; i < len(matrix); i++ {
				matrix[i][index] = 0
			}
		}
	}
}
