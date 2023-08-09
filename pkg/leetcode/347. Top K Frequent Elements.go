package leetcode

func TopKFrequent(nums []int, k int) []int {
	flagMap := make(map[int]bool)
	for _, v := range nums {
		flagMap[v] = true
	}
	keys := make([]int, 0, len(flagMap))
	// 遍历 map，将键添加到切片中
	for key := range flagMap {
		keys = append(keys, key)
	}
	heapSort(keys, k)
	return keys[len(keys)-k:]
}
