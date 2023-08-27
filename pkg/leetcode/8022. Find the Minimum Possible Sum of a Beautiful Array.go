package leetcode

func minimumPossibleSum(n int, target int) int64 {
	var result int64
	for i := 1; i <= n; i++ {
		if i <= target/2 {
			result = result + int64(i)
		} else {
			if target%2 == 0 {
				result = result + int64(i+target/2-1)
			} else {
				result = result + int64(i+target/2)
			}
		}
	}
	return result
}
