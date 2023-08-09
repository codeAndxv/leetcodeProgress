package leetcode

func findContentChildren(g []int, s []int) int {
	quickSort(g)
	quickSort(s)
	i := 0
	j := 0
	num := 0
	for i < len(g) && j < len(s) {
		if s[j] < g[i] {
			j++
		} else {
			j++
			i++
			num++
		}
	}
	return num
}
