package leetcode

func divideArray(nums []int) bool {
	resultMap := make(map[int]int)
	for _, value := range nums {
		if _, exist := resultMap[value]; !exist {
			resultMap[value] = 1
		} else {
			resultMap[value] = resultMap[value] + 1
		}
	}
	for _, value := range resultMap {
		if value%2 == 1 {
			return false
		}
	}
	return true
}
