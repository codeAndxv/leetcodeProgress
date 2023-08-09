package leetcode

func CountDigitOne(n int) int {
	temn := n
	count := 0
	for temn != 0 {
		count++
		temn = temn / 10
	}
	result := 0
	for i := count; i > 0; i-- {
		tem1 := n / pow(10, i)
		tem2 := (n / pow(10, i-1)) % 10
		if tem2 > 1 {
			result += (tem1 + 1) * pow(10, i-1)
		} else if tem2 == 1 {
			result += tem1*pow(10, i-1) + n%pow(10, i-1) + 1
		} else {
			result += tem1 * pow(10, i-1)
		}
	}
	return result
}
