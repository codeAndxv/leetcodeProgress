package leetcode

func SolveNQueens(n int) [][]string {
	status := make([][]bool, n)
	for i := 0; i < n; i++ {
		status[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			status[i][j] = true
		}
	}
	result := make([][]string, 0)
	solveNQueensBase(n, 0, status, result)
	return nil
}

func solveNQueensBase(n int, row int, status [][]bool, result [][]string) {
	for i := 0; i < n; i++ {
		if status[row][i] {
			if row == n-1 {
				append(result)
			} else {
				render(n, status, row, i)
				solveNQueensBase(n, row+1, status, result)
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
	for i := max(a+b-n, 0); i < min(a+b, n); i++ {
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
	for i := max(a+b-n, 0); i < min(a+b, n); i++ {
		status[i][a+b-i] = true
	}
}
