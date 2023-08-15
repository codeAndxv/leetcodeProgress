package main

import (
	"fmt"
	"temgo/pkg/leetcode"
)

func main() {
	fmt.Println(leetcode.FindReplaceString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
}
