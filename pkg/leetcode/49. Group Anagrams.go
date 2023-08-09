package leetcode

func groupAnagrams(strs []string) [][]string {
	wordMap := make(map[string][]string)
	for _, v := range strs {
		key := sortWord(v)
		if _, exist := wordMap[key]; !exist {
			wordMap[key] = []string{v}
		} else {
			wordMap[key] = append(wordMap[key], v)
		}
	}
	result := make([][]string, 0)
	for _, v := range wordMap {
		result = append(result, v)
	}
	return result
}

func sortWord(word string) string {
	temBytes := []byte(word)
	for i := 0; i < len(temBytes); i++ {
		for j := i + 1; j < len(temBytes); j++ {
			if temBytes[i] > temBytes[j] {
				tem := temBytes[i]
				temBytes[i] = temBytes[j]
				temBytes[j] = tem
			}
		}
	}
	return string(temBytes)
}
