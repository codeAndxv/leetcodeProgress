package leetcode

import "strconv"

func PredictTheWinner(nums []int) bool {
	scores := make(map[string]int)
	return PredictTheWinnerBase(nums, scores, 0, len(nums)-1, true)
	score1 :=
	score2 := PredictTheWinnerBase(nums, scores, 0, len(nums)-1, false)
	return score1 > score2
}

func PredictTheWinnerBase(nums []int, scores map[string]int, start int, end int, first bool) int {
	if start > end {
		return 0
	}
	key := getKey(start, end)
	if _, exist := scores[key]; exist {
		return scores[key]
	}
	var result int
	if first {
		if start == end {
			result = nums[start]
		}
		result = max(PredictTheWinnerBase(nums, scores, start+1, end, false)+nums[start],
			PredictTheWinnerBase(nums, scores, start, end-1, false)+nums[end])
	} else {
		if start == end {
			result = 0
		}
		result = max(PredictTheWinnerBase(nums, scores, start+1, end, true), PredictTheWinnerBase(nums, scores, start, end-1, true))
	}
	scores[key] = result
	return result
}

func getKey(start int, end int) string {
	return strconv.Itoa(start) + "_" + strconv.Itoa(end)
}
