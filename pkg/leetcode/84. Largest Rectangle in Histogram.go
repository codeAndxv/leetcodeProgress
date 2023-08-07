package leetcode

func largestRectangleArea(heights []int) int {
	leftLessMaxs := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		if i == 0 {
			leftLessMaxs[i] = 0
		} else {
			if heights[i] > heights[i-1] {
				leftLessMaxs[i] = i
			} else {
				tem := leftLessMaxs[i-1]
				for true {
					if tem > 0 {
						if heights[tem-1] >= heights[i] {
							tem = leftLessMaxs[tem-1]
						} else {
							break
						}
					} else {
						break
					}
				}
				leftLessMaxs[i] = tem
			}
		}
	}
	rightLessMins := make([]int, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		if i == len(heights)-1 {
			rightLessMins[i] = i
		} else {
			if heights[i] > heights[i+1] {
				rightLessMins[i] = i
			} else {
				tem := rightLessMins[i+1]
				for true {
					if tem < len(heights)-1 {
						if heights[tem+1] >= heights[i] {
							tem = rightLessMins[tem+1]
						} else {
							break
						}
					} else {
						break
					}
				}
				rightLessMins[i] = tem
			}
		}
	}
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		maxArea = max(maxArea, (rightLessMins[i]-leftLessMaxs[i]+1)*heights[i])
	}
	return maxArea
}
