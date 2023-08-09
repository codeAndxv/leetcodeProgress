package leetcode

func subtractProductAndSum(n int) int {
	product := 1
	sum := 0
	for n != 0 {
		product = product * (n % 10)
		sum = sum + n%10
		n = n / 10
	}
	return product - sum
}
