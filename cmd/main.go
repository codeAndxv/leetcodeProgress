package main

import (
	"fmt"
	"temgo/pkg/leetcode"
)

func main() {
	fmt.Println(leetcode.MinPathSum([][]int{
		{1, 3, 1}, {1, 5, 1}, {4, 2, 1},
	}))
}
