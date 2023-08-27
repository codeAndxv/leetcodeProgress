package leetcode

func MinOperations(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum < target {
		return -1
	}
	diff := sum - target
	opTimes := 0
	for diff != 0 {
		tem := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] >= diff && nums[i] < nums[tem] {
				tem = i
			}
		}
		if diff > nums[tem]/2 {
			diff -= nums[tem] / 2
		}
		nums[tem] = nums[tem] / 2
		opTimes++
	}
	return opTimes
}
