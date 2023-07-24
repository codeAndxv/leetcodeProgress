package leetcode

func maxProduct(nums []int) int {
	var maxAccum int
	var lastMaxAccum int
	var lastMinAccum int
	for index, value := range nums {
		if index == 0 {
			maxAccum = value
			lastMaxAccum = value
			lastMinAccum = value
			continue
		}
		tem := lastMaxAccum
		lastMaxAccum = max(max(lastMaxAccum*value, lastMinAccum*value), value)
		lastMinAccum = min(min(tem*value, lastMinAccum*value), value)

		maxAccum = max(lastMaxAccum, maxAccum)
	}
	return maxAccum
}
