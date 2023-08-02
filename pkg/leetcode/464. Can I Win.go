package leetcode

func CanIWin(maxChoosableInteger int, desiredTotal int) bool {
	return true
}

/**
	allChoose := make([]int, maxChoosableInteger)
	for index := 0; index < len(allChoose); index++ {
		allChoose[index] = index + 1
	}
	return canIWinBase(allChoose, desiredTotal, true)
}

func canIWinBase(allChoose []int, desiredTotal int, one bool) bool {
	if desiredTotal == 0 {
		return
	}
	for _, value := range allChoose {
		if desiredTotal == value {
			return true
		}
	}
	result := false
	answer := false
	for index := 0; index < len(allChoose); index++ {
		if result {
			break
		}
		if desiredTotal > allChoose[index] {
			tem := getNewSlice(allChoose, index)
			childResult := canIWinBase(tem, desiredTotal-allChoose[index], !one)
			answer = answer || childResult
			result = result || !childResult
			if result == true && one && len(tem) == 9 {
				fmt.Println(tem)
				fmt.Println(desiredTotal - allChoose[index])
			}
		}
	}
	return result && answer
}

func getNewSlice(source []int, index int) []int {
	tem := make([]int, len(source)-1)
	for index1, value := range source {
		if index1 == index {
			continue
		}
		if index1 > index {
			tem[index1-1] = value

		} else {
			tem[index1] = value
		}
	}
	return tem
}
*/
