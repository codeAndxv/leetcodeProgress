package leetcode

func wordBreak(s string, wordDict []string) bool {
	flags := make([]bool, len(s)+1)
	flags[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < len(wordDict); j++ {
			if i >= len(wordDict[j]) && s[i-len(wordDict[j]):i] == wordDict[j] {
				flags[i] = flags[i] || flags[i-len(wordDict[j])]
			}
			if flags[j] {
				continue
			}
		}
	}
	return flags[len(s)]
}
