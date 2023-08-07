package leetcode

func singleNumber2(nums []int) int {
	countMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, exist := countMap[nums[i]]; !exist {
			countMap[nums[i]] = 1
		} else {
			countMap[nums[i]] += 1
		}
	}
	for key, value := range countMap {
		if value == 1 {
			return key
		}
	}
	return 0
}
