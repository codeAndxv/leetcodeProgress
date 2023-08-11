package leetcode

import (
	"bytes"
)

func SolveNQueens(n int) [][]string {
	status := make([][]bool, n)
	for i := 0; i < n; i++ {
		status[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			status[i][j] = true
		}
	}
	result := make([][]string, 0)
	path := make([]int, n)
	for i := 0; i < n; i++ {
		path[i] = -1
	}
	solveNQueensBase(n, 0, status, path, &result)
	return result
}

func solveNQueensBase(n int, row int, status [][]bool, path []int, result *[][]string) {
	for i := 2; i < n; i++ {
		if status[row][i] {
			if row == n-1 {
				path[row] = i
				*result = append(*result, pathToStr(path))
				path[row] = -1
			} else {
				render(n, status, row, i)
				path[row] = i
				solveNQueensBase(n, row+1, status, path, result)
				path[row] = -1
				rollback(n, status, row, i)
			}
		}
	}
}

func render(n int, status [][]bool, a int, b int) {
	for i := 0; i < n; i++ {
		status[i][b] = false
	}
	for i := max(a-b, 0); i < min(n, n+a-b); i++ {
		status[i][-(a-b)+i] = false
	}
	for i := max(a+b-n+1, 0); i < min(a+b+1, n); i++ {
		status[i][a+b-i] = false
	}
}

func rollback(n int, status [][]bool, a int, b int) {
	for i := 0; i < n; i++ {
		status[i][b] = true
	}
	for i := max(a-b, 0); i < min(n, n+a-b); i++ {
		status[i][-(a-b)+i] = true
	}
	for i := max(a+b-n+1, 0); i < min(a+b+1, n); i++ {
		status[i][a+b-i] = true
	}
}

func pathToStr(path []int) []string {
	result := make([]string, len(path))
	for i := 0; i < len(path); i++ {
		var buffer bytes.Buffer
		for j := 0; j < len(path); j++ {
			if path[i] == j {
				buffer.WriteString("Q")
			} else {
				buffer.WriteString(".")
			}
		}
		result[i] = buffer.String()
	}
	return result
}
