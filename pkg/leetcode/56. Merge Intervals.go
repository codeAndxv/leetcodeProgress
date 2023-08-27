package leetcode

func merge1(intervals [][]int) [][]int {
	sortIntervals(intervals)
	currentIntervals := intervals[0]
	result := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= currentIntervals[0] && intervals[i][0] <= currentIntervals[1] {
			currentIntervals[1] = max(currentIntervals[1], intervals[i][1])
		} else {
			result = append(result, currentIntervals)
			currentIntervals = intervals[i]
		}
	}
	result = append(result, currentIntervals)
	return result
}

func sortIntervals(intervals [][]int) {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
}
