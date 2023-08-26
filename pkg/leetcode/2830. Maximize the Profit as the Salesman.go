package leetcode

import "fmt"

func maximizeTheProfit(n int, offers [][]int) int {
	records := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		for j := 0; j < len(offers); j++ {
			if offers[j][1]+1 == i {
				records[i] = max(max(records[offers[j][0]]+offers[j][2], records[i-1]), records[i])
			}
			records[i] = max(records[i], records[i-1])
		}
	}
	for i := 0; i <= n; i++ {
		fmt.Println(records[i])
	}
	return records[n]
}
