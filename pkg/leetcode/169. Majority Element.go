package leetcode

func majorityElement(nums []int) int {
	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v] += 1
		if countMap[v] > len(nums)/2 {
			return v
		}
	}
	return -1
}
