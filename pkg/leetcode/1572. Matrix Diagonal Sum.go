package leetcode

func diagonalSum(mat [][]int) int {
	sum := 0
	for i := 0; i < len(mat); i++ {
		sum += mat[i][i]
	}
	for i := 0; i < len(mat); i++ {
		if 2*i != len(mat)-1 {
			sum += mat[i][len(mat)-1-i]
		}
	}
	return sum
}
