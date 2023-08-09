package leetcode

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	/*if indexDiff == 1 {
		for i := 1; i < len(nums); i++ {
			if abs(nums[i]-nums[i-1]) <= valueDiff {
				return true
			}
		}
		return false
	} else {
		var temMin, temMin2, minIndex, temMax, temMax2, maxIndex int
		if nums[0] >= nums[1] {
			temMin = nums[1]
			temMin2 = nums[0]
			minIndex = 1
			temMax = nums[0]
			temMin = nums[1]
			maxIndex = 0
		} else {
			temMin = nums[0]
			temMin2 = nums[1]
			minIndex = 0
			temMax = nums[1]
			temMin = nums[0]
			maxIndex = 1
		}
		start := 0
		end := 1
		flagMap := make(map[int]bool)
		for i := 2; i < len(nums); i++ {
			if abs(nums[i]-temMin) <= valueDiff || abs(nums[i]-temMax) <= valueDiff {
				return true
			}
			if i-start < indexDiff {
				end = i
				flagMap[nums[i]] = true
			} else {
				if temMax == start {
					tem
				}
				start = start + 1
				end = i
			}
			if nums[i] > temMax {
				temMax2 = temMax
				temMax = nums[i]
			}
			if nums[i] < temMin {
				temMin2 = temMin
				temMin = nums[i]
			}
		}
	}*/
	return false
}
