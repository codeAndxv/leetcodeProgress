package leetcode

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	for i := len(answer) - 2; i >= 0; i-- {
		tem := i + 1
		for tem < len(answer) {
			if temperatures[i] < temperatures[tem] {
				answer[i] = tem - i
				break
			} else {
				if answer[tem] == 0 {
					answer[i] = 0
				} else {
					tem = tem + answer[tem]
				}
			}
		}
	}
	return answer
}
