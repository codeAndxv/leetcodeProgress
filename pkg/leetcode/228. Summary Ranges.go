package leetcode

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	result := make([]string, 0)
	lastValue := nums[0]
	start := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == lastValue+1 {
			lastValue = nums[i]
		} else {
			if start == lastValue {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, lastValue))
			}
			lastValue = nums[i]
			start = nums[i]
		}
	}
	if start == lastValue {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, lastValue))
	}
	return result
}
