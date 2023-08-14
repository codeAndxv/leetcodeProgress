package main

import (
	"fmt"
)

func getBit(v int) int {
	count := 0
	for v != 0 {
		v = v / 10
		count++
	}
	return count
}

func main() {
	fmt.Println(getBit(101))
}
