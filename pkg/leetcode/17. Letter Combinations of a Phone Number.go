package leetcode

func letterCombinations(digits string) []string {
	if digits == "" || len(digits) < 1 {
		return nil
	}
	collect := make([]string, 0)
	digitToLetters := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	letterCombinationsBase(digits, digitToLetters, 0, "", &collect)
	return collect
}

func letterCombinationsBase(digits string, digitToLetters map[string]string, index int, current string, collect *[]string) {
	if index == len(digits) {
		*collect = append(*collect, current)
		return
	}
	for _, v := range digitToLetters[string(digits[index])] {
		newStr := current + string(v)
		letterCombinationsBase(digits, digitToLetters, index+1, newStr, collect)
	}
}
