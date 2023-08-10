package leetcode

func lengthOfLIS(nums []int) int {
	record := make([]int, len(nums))
	record[len(nums)-1] = 0
	maxLen := 0
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				record[i] = max(record[i], record[j]+1)
				maxLen = max(maxLen, record[i])
			}
		}
	}
	return maxLen + 1
}

/**
func lengthOfLIS(nums []int) int {
    // 贪心 + 二分
    // dp[i]表示i长度的最后一个数的值
    // 贪心的把相同长度的最后一个数的值尽量减少
    // 对于dp数组来说 是单调递增的, 可以用二分查找最大的比当前num小的数
    n := len(nums)
    dp := make([]int, n + 1)
    Len := 0
    for _, num := range nums {
        left, right := 0, Len
        for left < right {
            mid := (left + right + 1) >> 1
            // 二分找到最大的小于num的值
            if dp[mid] < num {
                left = mid
            } else {
                right = mid - 1
            }
        }
        Len = max(Len, left + 1)
        dp[left + 1] = num
    }

    return Len
}

func max(i,j int) int {
    if i > j {
        return i
    }
    return j
}
*/
