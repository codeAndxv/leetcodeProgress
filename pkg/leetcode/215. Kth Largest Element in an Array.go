package leetcode

func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums) - 1
	stand := nums[0]
	for right > left {
		for right >= 0 && right > left {
			if nums[right] >= stand {
				right--
			} else {
				break
			}
		}
		for left < len(nums) && right > left {
			if nums[left] <= stand {
				left++
			} else {
				break
			}
		}
		if right > left {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[left], nums[0] = stand, nums[left]
	if left < len(nums)-k {
		return findKthLargest(nums[left+1:], k)
	}
	if left > len(nums)-k {
		return findKthLargest(nums[:left], k-(len(nums)-left))
	}
	return nums[left]
}

/**
var num1, num2, num3 int
var nums1, nums3 []int
for _, value := range nums {
	if value == nums[0] {
		num2++
		continue
	}
	if value > nums[0] {
		num3++
		nums3 = append(nums3, value)
		continue
	}
	if value < nums[0] {
		num1++
		nums1 = append(nums1, value)
		continue
	}
}
if num3 >= k {
	return FindKthLargest(nums3, k)
}
if num3+num2 >= k {
	return nums[0]
}
return FindKthLargest(nums1, k-num3-num2)
*/
