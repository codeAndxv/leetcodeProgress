package leetcode

import "math"

func maximumGap(nums []int) int {
	dividend := 0
	for i := 8; i >= 0; i-- {
		dividend = dividend + int(math.Pow10(i))
		for j := 0; j < len(nums); j++ {

		}
	}
	return 0
}
