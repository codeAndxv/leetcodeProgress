package leetcode

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		for j := i; j < len(matrix)-1-i; j++ {
			tem := matrix[i][j]
			matrix[i][j] = matrix[len(matrix)-1-j][i]
			matrix[len(matrix)-1-j][i] = matrix[len(matrix)-1-i][len(matrix)-1-j]
			matrix[len(matrix)-1-i][len(matrix)-1-j] = matrix[j][len(matrix)-1-i]
			matrix[j][len(matrix)-1-i] = tem
		}
	}
}
