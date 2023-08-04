package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	status := make([]int, len(gas))
	for index := 0; index < len(status); index++ {
		status[index] = gas[index] - cost[index]
	}

	for index := 0; index < len(status); index++ {
		sum := 0
		for start := index; ; start = (start + 1) % len(status) {
			sum = sum + status[start]
			if sum < 0 {
				break
			}

			if start == (index-1+len(status))%len(status) {
				return index
			}
		}
	}
	return -1
}
