package leetcode

import (
	"strconv"
	"strings"
)

func MaxDiff(num int) int {
	stra := strconv.Itoa(num)
	strb := strconv.Itoa(num)
	var replace1 string
	// 遍历字符串中的每个字符
	for _, char := range stra {
		if char == '9' {
			continue
		}
		replace1 = string(char)
		break
	}
	if len(replace1) > 0 {
		stra = strings.Replace(stra, replace1, "9", -1)
	}

	var replace2, newValue string
	// 遍历字符串中的每个字符
	firstCan := true
	for index, char := range strb {
		if index == 0 {
			if char == '1' {
				firstCan = false
				continue
			} else {
				replace2 = string(char)
				newValue = "1"
				break
			}
		} else {
			if char == '0' || (!firstCan && char == '1') {
				continue
			} else {
				replace2 = string(char)
				newValue = "0"
				break
			}
		}
	}
	if len(replace2) > 0 {
		strb = strings.Replace(strb, replace2, newValue, -1)
	}
	numa, err := strconv.Atoi(stra)
	if err != nil {
		return 0
	}
	numb, err := strconv.Atoi(strb)
	if err != nil {
		return 0
	}
	return numa - numb
}
