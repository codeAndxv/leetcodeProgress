package leetcode

func PartitionLabels(s string) []int {
	askIndexMap := make(map[string]int)
	source := []byte(s)
	for i, v := range source {
		askIndexMap[string(v)] = i
	}
	count := 0
	askIndex := 0
	result := make([]int, 0)
	for i, v := range s {
		askIndex = max(askIndex, askIndexMap[string(v)])
		count++
		if i == askIndex {
			result = append(result, count)
			count = 0
		}
	}
	return result
}
