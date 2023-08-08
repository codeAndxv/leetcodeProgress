package leetcode

func maxAbsoluteSum(nums []int) int {
	maxSum := 0
	temSum := 0
	for i := 0; i < len(nums); i++ {
		if temSum+nums[i] < 0 {
			temSum = 0
			continue
		}
		temSum += nums[i]
		maxSum = max(maxSum, temSum)
	}
	minSum := 0
	temSum = 0
	for i := 0; i < len(nums); i++ {
		if temSum+nums[i] > 0 {
			temSum = 0
			continue
		}
		temSum += nums[i]
		minSum = min(minSum, temSum)
	}
	return max(abs(minSum), maxSum)
}

/**
}
func maxAbsoluteSum(nums []int) int {
    n:=len(nums)
    preSum:=make([]int,n+1)
    preSum[0]=0
    for i:=1;i<=n;i++ {
        preSum[i]=preSum[i-1]+nums[i-1]
    }
    maxOne,minOne:=math.MinInt,math.MaxInt
    for i:=0;i<=n;i++ {
        maxOne=max(maxOne,preSum[i])
        minOne=min(minOne,preSum[i])
    }
    return maxOne-minOne
}
*/
