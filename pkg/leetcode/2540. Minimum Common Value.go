package leetcode

func GetCommon(nums1 []int, nums2 []int) int {
	i := 0
	j := 0
	result := -1
	for i < len(nums1) && j < len(nums2) {
		for i < len(nums1) && nums1[i] < nums2[j] {
			i++
		}
		for i < len(nums1) && j < len(nums2) && nums2[j] < nums1[i] {
			j++
		}

		if i < len(nums1) && j < len(nums2) && nums1[i] == nums2[j] {
			result = nums1[i]
			break
		}
	}
	return result
}
