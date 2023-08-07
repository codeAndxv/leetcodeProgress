package leetcode

func lengthOfLastWord(s string) int {
	start := -1
	end := -1
	tems := []byte(s)
	for i := len(tems) - 1; i >= 0; i-- {
		if tems[i] == ' ' {
			if end != -1 {
				break
			}
		} else {
			if end != -1 {
				start = i
			} else {
				end = i
				start = i
			}
		}
	}
	return end - start + 1
}
