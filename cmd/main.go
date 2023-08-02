package main

import (
	"fmt"
	"temgo/pkg/leetcode"
)

func main() {
	fmt.Println(leetcode.SearchMatrix([][]int{{1, 3, 5, 7, 9}, {2, 4, 6, 8, 10}, {11, 13, 15, 17, 19}, {12, 14, 16, 18, 20}, {21, 22, 23, 24, 25}}, 11))
}

/*
1   3   5   7    9
2   4   6   8   10
11  13  15  17  19
12, 14, 16, 18, 20
21, 22, 23, 24, 25
*/
