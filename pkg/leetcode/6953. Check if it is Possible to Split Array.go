package leetcode

func CanSplitArray(nums []int, m int) bool {
	if len(nums) < 3 {
		return true
	}
	record := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		record[i] = make([]int, len(nums))
	}
	return canSplitArrayBase(nums, 0, len(nums)-1, m, record)
}

func canSplitArrayBase(nums []int, start int, end int, m int, record [][]int) bool {
	if start == end {
		return true
	}
	if record[start][end] == 1 {
		return true
	}
	if record[start][end] == -1 {
		return false
	}
	sum := 0
	for i := start; i <= end; i++ {
		sum += nums[i]
	}
	if sum < m {
		return false
	}

	solve := false
	for i := start + 1; i <= end; i++ {
		solve = solve || (canSplitArrayBase(nums, start, i-1, m, record) && canSplitArrayBase(nums, i, end, m, record))
	}
	if solve {
		record[start][end] = 1
	} else {
		record[start][end] = -1
	}
	return solve
}
