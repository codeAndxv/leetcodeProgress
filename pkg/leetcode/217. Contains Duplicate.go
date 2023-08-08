package leetcode

func containsDuplicate(nums []int) bool {
	countMap := make(map[int]uint8)
	for _, v := range nums {
		if _, exist := countMap[v]; !exist {
			countMap[v] = 1
		} else {
			return true
		}
	}
	return false
}
