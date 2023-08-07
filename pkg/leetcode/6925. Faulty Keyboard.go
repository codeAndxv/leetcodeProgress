package leetcode

func finalString(s string) string {
	byteSlice := []byte{}
	for _, c1 := range s {
		if c1 == 'i' {
			reverseBytes(byteSlice)
		} else {
			byteSlice = append(byteSlice, byte(c1))
		}
	}
	return string(byteSlice)
}
func reverseBytes(byteSlice []byte) {
	left := 0
	right := len(byteSlice) - 1

	for left < right {
		// Swap elements at left and right indices
		byteSlice[left], byteSlice[right] = byteSlice[right], byteSlice[left]

		// Move the pointers towards the middle
		left++
		right--
	}
}
