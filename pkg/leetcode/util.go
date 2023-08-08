package leetcode

import "math"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func pow(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sort(source []int) []int {
	if len(source) < 2 {
		return source
	}
	for i := 0; i < len(source); i++ {
		for j := i + 1; j < len(source); j++ {
			if source[i] > source[j] {
				source[i], source[j] = source[j], source[i]
			}
		}
	}
	return source
}
