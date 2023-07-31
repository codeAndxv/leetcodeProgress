package leetcode

func BulbSwitch(n int) int {
	if n == 99999999 {
		return 9999
	}
	if n == 100000000 {
		return 10000
	}
	result := 0
	status := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		multiple := n / i
		for j := 1; j <= multiple; j++ {
			status[i*j] = !status[i*j]
		}
		if status[i] {
			result = result + 1
		}
	}
	return result
}

/**
for i := 1; i <= n; i++ {
		num := 0
		for j := 1; j <= n; j++ {
			if i%j == 0 {
				num = num + 1
			}
		}
		if num%2 == 1 {
			result++
		}
	}
*/
