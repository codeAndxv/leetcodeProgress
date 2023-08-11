package leetcode

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	records := make([]map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		tem := make(map[int]bool)
		if i == 0 {
			tem[0] = true
			if nums[i] == sum {
				return true
			}
			if nums[i] < sum {
				tem[nums[i]] = true
			}
		} else {
			for j, _ := range records[i-1] {
				tem[j] = true
				if j+nums[i] == sum {
					return true
				}
				if j+nums[i] < sum {
					tem[j+nums[i]] = true
				}
			}
		}
		records[i] = tem
	}
	return false
}

/**
// 01背包状态转移方程为：dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
// 显然我们能够优化成滚动数组，采用一位数组即可 dp[j] 表示负重最大为 j 的背包最多能装多少货物，使得价值最大
// 此时 dp[j] = max(dp[j], dp[j-weight[i]]+value[i])，此时 每次遍历 j，都从后往前，因为 dp[j] 依赖于前面的 dp[j]，从后往前不会导致语义出现问题
// 背包最大负重为 sum/2，并且 sum 不能为奇数，否则不能切分成两个和相等的子集
func canPartition(nums []int) bool {
	var sum int
	for i := range nums {
		sum += nums[i]
	}
	if (sum & 1) == 1 {
		return false
	}
	dp := make([]int, sum/2+1)
	// 以下为滚动数组，当 j < nums[i] 时，显然 dp[j] = dp[j]，相当于复用上一层的 dp[j]，即未选取当前的物品 i
	for i := range nums {
		for j := sum / 2; j >= nums[i]; j-- {
			if dp[j] < dp[j-nums[i]]+nums[i] {
				dp[j] = dp[j-nums[i]] + nums[i]
			}
		}
	}
	return dp[sum/2] == sum/2
}
*/

/*func canPartitionBase(nums []int, index int, left int) bool {
	if left == 0 {
		return true
	}
	if left < 0 {
		return false
	}
	if index == len(nums) {
		return false
	}
	if canPartitionBase(nums, index+1, left-nums[index]) {
		return true
	}
	return canPartitionBase(nums, index+1, left)
}
*/
