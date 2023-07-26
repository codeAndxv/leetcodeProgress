package leetcode

func threeSum(nums []int) [][]int {
	nums = bubbleSort(nums)
	numMap := make(map[int]int)
	for _, value := range nums {
		numMap[value] += 1
	}
	resultMap := make(map[int]map[int]int)
	for index1 := 0; index1 < len(nums); index1++ {
		for index2 := index1 + 1; index2 < len(nums); index2++ {
			tem := -(nums[index1] + nums[index2])
			sameNum := 0
			if tem == nums[index1] {
				sameNum++
			}
			if tem == nums[index2] {
				sameNum++
			}
			if numMap[tem] > sameNum {
				if (resultMap[nums[index1]] == nil || resultMap[nums[index1]][nums[index2]] == 0) && (resultMap[nums[index2]] == nil || resultMap[nums[index2]][nums[index1]] == 0) {

				}
				if !(resultMap[nums[index1]][nums[index2]] > 0 || resultMap[nums[index2]][nums[index1]] > 0) {
					resultMap[nums[index1]][nums[index2]] += 1
				}
			}
		}
	}
	return nil
}

func bubbleSort(nums []int) []int {
	for index1 := 0; index1 < len(nums)-1; index1++ {
		for index2 := index1; index2 < len(nums)-1; index2++ {
			if nums[index2] > nums[index2+1] {
				tem := nums[index2]
				nums[index2] = nums[index2+1]
				nums[index2+1] = tem
			}
		}
	}
	return nums
}
