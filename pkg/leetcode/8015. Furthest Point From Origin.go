package leetcode

func furthestDistanceFromOrigin(moves string) int {
	left := 0
	right := 0
	blank := 0
	for _, item := range moves {
		if item == 'L' {
			left++
		}
		if item == 'R' {
			right++
		}
		if item == '_' {
			blank++
		}
	}
	return max(left, right) + blank - min(left, right)
}
