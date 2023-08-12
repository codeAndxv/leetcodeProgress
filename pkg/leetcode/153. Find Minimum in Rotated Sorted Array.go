package leetcode

func findMin(nums []int) int {

}

func findMinBase(nums []int, start int, end int) int {
	mid := (start + end) / 2
	if nums[start] > nums[end] {
		findMinBase(nums, mid, end)
	}
}
