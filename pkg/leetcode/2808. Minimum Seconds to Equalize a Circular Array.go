package leetcode

import "math"

func minimumSeconds(nums []int) int {
	indexMap := make(map[int][]int)
	for i, v := range nums {
		if _, exist := indexMap[v]; !exist {
			indexMap[v] = []int{i}
		} else {
			indexMap[v] = append(indexMap[v], i)
		}
	}
	minStep := math.MaxInt32
	for _, v := range indexMap {
		minStep = min(minStep, needStep(v, len(nums)))
	}
	return minStep
}

func needStep(indexs []int, numlength int) int {
	maxDistance := 0
	if len(indexs) == 1 {
		return numlength / 2
	}
	for i := 0; i < len(indexs); i++ {
		if i == 0 {
			maxDistance = max(maxDistance, (indexs[i]+numlength-indexs[len(indexs)-1])/2)
		} else {
			maxDistance = max(maxDistance, (indexs[i]-indexs[i-1])/2)
		}
	}
	return maxDistance
}

/*
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
*/
