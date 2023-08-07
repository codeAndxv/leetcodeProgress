package leetcode

func Rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix)-1; j++ {
			tem := matrix[i+j][j]
			matrix[i+j][j] = matrix[len(matrix)-1-j][i+j]
			matrix[len(matrix)-1-j][i+j] = matrix[len(matrix)-1-i-j][len(matrix)-1-j]
			matrix[len(matrix)-1-i-j][len(matrix)-1-j] = matrix[j][len(matrix)-1-i-j]
			matrix[j][len(matrix)-1-i-j] = tem
		}
	}
}
