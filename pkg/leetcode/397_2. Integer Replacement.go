package leetcode

func integerReplacement2(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return integerReplacement2(n/2) + 1
	} else {
		return min(integerReplacement2(n-1)+1, integerReplacement2(n+1)+1)
	}
}
