package leetcode

func MinimumSeconds(nums []int) int {
	num1 := getMostNum(nums)
	step := 0
	handle := false
	n := len(nums)
	for true {
		handle = false
		var tem []int
		for i := 0; i < n; i++ {
			if nums[i] == num1 {
				continue
			}
			if nums[(i-1+n)%n] == num1 {
				handle = true
				tem = append(tem, i)
				continue
			}
			if nums[(i+1)%n] == num1 {
				tem = append(tem, i)
				handle = true
				continue
			}
			handle = true
		}
		if handle {
			for _, v1 := range tem {
				nums[v1] = num1
			}
		}
		if handle {
			step = step + 1
		} else {
			break
		}
	}
	return step
}

func getMostNum(nums []int) int {
	occurMap := make(map[int]int)
	for _, v1 := range nums {
		if _, exist := occurMap[v1]; !exist {
			occurMap[v1] = 1
		} else {
			occurMap[v1] += 1
		}
	}
	var mostV, mostK int
	for key, value := range occurMap {
		if value > mostV {
			mostV = value
			mostK = key
		}
	}
	return mostK
}
