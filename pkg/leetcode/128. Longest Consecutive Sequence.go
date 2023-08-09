package leetcode

func longestConsecutive(nums []int) int {
	flagMap := make(map[int]bool)
	for _, v := range nums {
		flagMap[v] = true
	}
	result := 0
	for num, _ := range flagMap {
		tem := 1
		delete(flagMap, num)
		for i := 1; flagMap[num-i]; i++ {
			delete(flagMap, num-i)
			tem++
		}
		for i := 1; flagMap[num+i]; i++ {
			delete(flagMap, num+i)
			tem++
		}
		result = max(result, tem)
	}
	return result
}
