package leetcode

func FindPeakElement(nums []int) int {
	return findPeakElementBase(nums, 0, len(nums)-1)
}

func findPeakElementBase(nums []int, low int, high int) int {
	mid := (low + high) / 2
	flagL := true
	flagR := true
	if mid+1 != len(nums) {
		flagR = nums[mid] > nums[mid+1]
	}
	if mid-1 != -1 {
		flagL = nums[mid] > nums[mid-1]
	}
	if flagL && flagR {
		return mid
	} else {
		if mid > low {
			tem := findPeakElementBase(nums, low, mid-1)
			if tem != -1 {
				return tem
			}
		}
		if mid < high {
			return findPeakElementBase(nums, mid+1, high)
		}
	}
	return -1
}
