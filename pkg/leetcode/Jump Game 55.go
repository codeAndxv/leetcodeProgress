package leetcode

func canJump(nums []int) bool {
	trap := false
	for index, value := range nums {
		if value == 0 && index != len(nums)-1 {
			trap = true
			for left := index - 1; left >= 0; left-- {
				if nums[left]+left > index {
					trap = false
					break
				}
			}
		}
		if trap {
			break
		}
	}
	return !trap
}
