package leetcode

func CanIWin(maxChoosableInteger int, desiredTotal int) bool {
	allChoose := make([]int, maxChoosableInteger)
	for index := 0; index < len(allChoose); index++ {
		allChoose[index] = index
	}
	return canIWinBase(allChoose, desiredTotal, true)
}

func canIWinBase(allChoose []int, desiredTotal int, one bool) bool {
	if len(allChoose) <= 0 {
		return (one && desiredTotal == 0) || (!one && desiredTotal != 0)
	}
	result := false
	for index, value := range allChoose {
		if !result {
			break
		}
		result = result || !canIWinBase(append(allChoose[:index], allChoose[index+1:]...), desiredTotal-value, !one)
	}
	return result
}
