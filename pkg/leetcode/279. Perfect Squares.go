package leetcode

import "math"

func NumSquares(n int) int {
	record := make([]int, n+1)
	for i := 0; i <= n; i++ {
		record[i] = -1
	}
	return numSquaresBase(n, record, 0, math.MaxInt32)
}

func numSquaresBase(n int, record []int, currentStep int, limitStep int) int {
	x := int(math.Sqrt(float64(n)))
	if record[n] != -1 {
		return record[n]
	}
	leastStep := n
	if x*x == n {
		record[n] = 1
		return record[n]
	}

	for i := 0; i < x; i++ {
		tem := numSquaresBase(n-(x-i)*(x-i), record, currentStep+1, leastStep) + 1
		if tem != -1 {
			leastStep = min(leastStep, tem)
		}
	}
	record[n] = leastStep
	return record[n]
}
