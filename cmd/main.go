package main

import (
	"fmt"
	"temgo/pkg/leetcode"
)

func main() {
	fmt.Println(leetcode.minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}
