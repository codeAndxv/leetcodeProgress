package leetcode

func minimumSum(n int, k int) int {
	result := 0
	for i := 1; i <= n; i++ {
		if i <= k/2 {
			result = result + i
		} else {
			if k%2 == 0 {
				result = result + i + k/2 - 1
			} else {
				result = result + i + k/2
			}

		}
	}
	return result
}
