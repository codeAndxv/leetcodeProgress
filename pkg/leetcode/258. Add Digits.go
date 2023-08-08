package leetcode

func addDigits(num int) int {
	for num/10 > 0 {
		num = addSelf(num)
	}
	return num
}

func addSelf(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	return sum
}
