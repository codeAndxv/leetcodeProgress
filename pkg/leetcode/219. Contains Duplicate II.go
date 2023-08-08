package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)
	for i, v := range nums {
		if _, exist := indexMap[v]; !exist {
			indexMap[v] = i
		} else {
			if i-indexMap[v] <= k {
				return true
			} else {
				indexMap[v] = i
			}
		}
	}
	return false
}
