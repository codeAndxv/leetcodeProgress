package leetcode

func trailingZeroes(n int) int {
	result := 0
	for i := 1; 5*i <= n; i++ {
		tem := 0
		tem2 := i
		for tem2 != 0 {
			tem++
			if tem2%5 == 0 {
				tem2 = tem2 / 5
			} else {
				break
			}
		}
		result = result + tem
	}
	return result
}
