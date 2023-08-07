package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid) - 1
	n := len(obstacleGrid[0]) - 1
	record := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		record[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			record[i][j] = -1
		}
	}
	return uniquePathsWithObstaclesBase(obstacleGrid, m, n, 0, 0, record)
}

func uniquePathsWithObstaclesBase(obstacleGrid [][]int, m int, n int, x int, y int, record [][]int) int {
	if obstacleGrid[x][y] == 1 {
		return 0
	}
	if x == m && y == n {
		return 1
	}
	count := 0
	if x < m {
		if record[x+1][y] == -1 {
			record[x+1][y] = uniquePathsWithObstaclesBase(obstacleGrid, m, n, x+1, y, record)
		}
		count += record[x+1][y]
	}
	if y < n {
		if record[x][y+1] == -1 {
			record[x][y+1] = uniquePathsWithObstaclesBase(obstacleGrid, m, n, x, y+1, record)
		}
		count += record[x][y+1]
	}
	record[x][y] = count
	return record[x][y]
}
