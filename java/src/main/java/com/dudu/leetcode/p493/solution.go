package main

import "fmt"

/**
 * datetime: 1/1/2020 1:54 PM
 * user: luzhao
 * email: luzhao233@gmail.com
 */
func reversePairs(nums []int) int {
	result := 0
	sortNums := make([]int, 0)
	// inverted sort
	for j := len(nums) - 1; j >= 0; j-- {
		num := nums[j]
		index := 0
		for _, tem := range sortNums {
			if tem < num {
				result++
			}
			if tem < 2*num {
				index++
			}
			if !(tem < num || tem < 2*num) {
				break
			}
		}
		if index >= len(sortNums) {
			sortNums = append(sortNums, 2*num)
		} else {
			sortNums = append(sortNums[:index+1], sortNums[index:]...)
			sortNums[index] = 2 * num
		}
	}
	return result
}

func main() {
	fmt.Print(reversePairs([]int{-240, -185, -241, -4, -315, 188}))
}
