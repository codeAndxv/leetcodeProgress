package leetcode

func integerReplacement(n int) int {
	step := 0
	for n != 1 {
		step++
		if n%2 == 0 {
			n = n / 2
			continue
		}
		if n == 3 {
			n = n - 1
			continue
		}
		// 对比哪一种走的步长
		if getPow(n+1) > getPow(n-1) {
			n = n + 1
		} else {
			n = n - 1
		}
	}
	return step
}

func getPow(a int) int {
	step := 0
	for a%2 == 0 {
		a = a / 2
		step++
	}
	return step
}

func IntegerReplacement(n int) int {
	return integerReplacement(n)
}
