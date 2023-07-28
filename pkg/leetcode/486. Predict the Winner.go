package leetcode

func PredictTheWinner(nums []int) bool {
	return PredictTheWinnerBase(nums, 0, 0, true)
}

func PredictTheWinnerBase(nums []int, score1 int, score2 int, one bool) bool {
	if len(nums) <= 0 {
		return score1 > score2 || (score1 == score2 && one)
	}
	if len(nums) == 1 {
		return score1+nums[0] > score2 || (score1+nums[0] == score2 && one)
	}
	result1 := PredictTheWinnerBase(nums[:len(nums)-1], score2, score1+nums[len(nums)-1], !one)
	result2 := PredictTheWinnerBase(nums[1:], score2, score1+nums[0], !one)
	return !result1 || !result2
}

/*
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
*/
