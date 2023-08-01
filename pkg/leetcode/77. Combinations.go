package leetcode

func Combine(n int, k int) [][]int {
	start := make([]int, k)
	for index, _ := range start {
		start[index] = index
	}
	var result [][]int
	resume := true
	for resume {
		tem := make([]int, k)
		for index, _ := range tem {
			tem[index] = start[index] + 1
		}
		result = append(result, tem)
		start, resume = autoIncrement(start, n)
	}
	return result
}

func autoIncrement(source []int, n int) ([]int, bool) {
	tem := len(source) - 1
	for ; tem >= 0; tem-- {
		if source[tem] < n-(len(source)-tem) {
			break
		}
	}
	if tem == -1 {
		return nil, false
	}
	for i := tem; i < len(source); i++ {
		if tem == i {
			source[i] = source[i] + 1
		} else {
			source[i] = source[i-1] + 1
		}
	}
	return source, true
}
