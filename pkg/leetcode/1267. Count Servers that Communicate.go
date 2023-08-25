package leetcode

func countServers(grid [][]int) int {
	flagMap := make(map[int]bool)
	for i := 0; i < len(grid); i++ {
		isFirst := false
		lastAdd := -1
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				if !isFirst {
					isFirst = true
					lastAdd = i*1000 + j
				} else {
					flagMap[lastAdd] = true
					flagMap[i*1000+j] = true
				}
			}
		}
	}

	for j := 0; j < len(grid[0]); j++ {
		isFirst := false
		lastAdd := -1
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == 1 {
				if !isFirst {
					isFirst = true
					lastAdd = i*1000 + j
				} else {
					flagMap[lastAdd] = true
					flagMap[i*1000+j] = true
				}
			}
		}
	}
	return len(flagMap)
}
