package leetcode

import "fmt"

func CircularGameLosers(n int, k int) []int {
	countMap := make(map[int]int)
	index := 0
	count := 0
	for true {
		index = (index + count*k) % n
		fmt.Println(index)
		countMap[index] += 1
		if countMap[index] >= 2 {
			break
		}
		count++
	}
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if countMap[i] == 0 {
			result = append(result, i+1)
		}
	}
	return result
}
