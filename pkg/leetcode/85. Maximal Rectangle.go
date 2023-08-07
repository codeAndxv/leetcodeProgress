package leetcode

func maximalRectangle(matrix [][]byte) int {
	maxArea := 0
	heights := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j] += 1
			}
		}
		maxArea = max(maxArea, largestRectangleArea(heights))
	}
	return maxArea
}

/*func maximalRectangle(matrix [][]byte) int {
	maxArea := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			maxArea = max(maxArea, getMaxArea(matrix, i, j))
		}
	}
	return maxArea
}
func getMaxArea(matrix [][]byte, a int, b int) int {
	len1 := math.MaxInt32
	maxArea := 0
	for i := a; i < len(matrix) && matrix[i][b] == '1'; i++ {
		len2 := 0
		for j := b; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				len2++
			} else {
				break
			}
		}
		len1 = min(len1, len2)
		maxArea = max(maxArea, (i-a+1)*len1)
	}
	return maxArea
}
*/
