package leetcode

/*func SubarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	sumsMap := make(map[int][]int)
	sumsMap[0] = []int{-1}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, exist := sumsMap[sum]; !exist {
			sumsMap[sum] = []int{i}
		} else {
			sumsMap[sum] = append(sumsMap[sum], i)
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		count += len(sumsMap[sum-k])
		sumsMap[sum] = sumsMap[sum][:len(sumsMap[sum])-1]
		sum -= nums[i]
	}
	if k == 0 {
		count -= len(nums)
	}
	return count
}
*/

func SubarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre]++
	}
	return count
}

/**

 */

/*func SubarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}
*/
