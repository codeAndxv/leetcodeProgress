package leetcode

func reverseString(s []byte) {
	low := 0
	high := len(s) - 1
	for low < high {
		tem := s[low]
		s[low] = s[high]
		s[high] = tem
		low++
		high--
	}
}
