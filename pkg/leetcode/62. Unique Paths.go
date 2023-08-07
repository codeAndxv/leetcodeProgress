package leetcode

func uniquePaths(m int, n int) int {
	record := make([][]int, m)
	for i := 0; i < m; i++ {
		record[i] = make([]int, n)
	}
	return uniquePathsBase(m, n, 1, 1, record)
}

func uniquePathsBase(m int, n int, x int, y int, record [][]int) int {
	if x == m || y == n {
		record[x-1][y-1] = 1
		return record[x-1][y-1]
	}
	count := 0
	if x < m {
		if record[x][y-1] == 0 {
			record[x][y-1] = uniquePathsBase(m, n, x+1, y, record)
		}
		count += record[x][y-1]
	}
	if y < n {
		if record[x-1][y] == 0 {
			record[x-1][y] = uniquePathsBase(m, n, x, y+1, record)
		}
		count += record[x-1][y]
	}
	return count
}
