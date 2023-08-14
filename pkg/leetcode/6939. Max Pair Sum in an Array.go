package leetcode

func maxSum(nums []int) int {
	quickSort(nums)
	bitNumMap := make(map[int][]int)
	for _, v := range nums {
		temBit := findMaxBit(v)
		if _, ok := bitNumMap[temBit]; !ok {
			bitNumMap[temBit] = []int{v}
		} else {
			bitNumMap[temBit] = append(bitNumMap[temBit], v)
		}
	}
	maxValue := -1
	for _, v := range bitNumMap {
		if len(v) > 1 {
			maxValue = max(maxValue, v[len(v)-2]+v[len(v)-1])
		}
	}
	return maxValue
}
func findMaxBit(v int) int {
	maxBit := 0
	for v != 0 {
		maxBit = max(maxBit, v%10)
		v = v / 10
	}
	return maxBit
}
