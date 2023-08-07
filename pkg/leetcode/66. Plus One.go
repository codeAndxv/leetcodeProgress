package leetcode

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			if i == 0 {
				digits[i] = 0
				// 创建一个新的切片，长度比原来的切片多1
				newSlice := make([]int, len(digits)+1)
				// 将新的元素放在新切片的第一个位置
				newSlice[0] = 1
				// 将原来的切片中的元素复制到新切片中，从新元素后面开始复制
				copy(newSlice[1:], digits)
				return newSlice
			} else {
				digits[i] = 0
			}
		} else {
			digits[i] = digits[i] + 1
			break
		}
	}
	return digits
}
