package leetcode

func topKFrequent(nums []int, k int) []int {
	countsMap := make(map[int]int)
	for _, v := range nums {
		countsMap[v] += 1
	}
	counts := make([]int, 0, len(countsMap))
	// 遍历 map，将键添加到切片中
	for _, v := range countsMap {
		counts = append(counts, v)
	}
	heapSort(counts, k)
	result := make([]int, 0)
	for key, v := range countsMap {
		if v >= counts[len(counts)-k] {
			result = append(result, key)
		}
	}
	return result
}
