package leetcode

func maxDistToClosest(seats []int) int {
	maxDis := 0
	lastSeat := -1
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			if lastSeat == -1 {
				maxDis = max(maxDis, i)
				lastSeat = i
			} else {
				maxDis = max(maxDis, (i-lastSeat)/2)
				lastSeat = i
			}
		}
	}
	maxDis = max(maxDis, len(seats)-1-lastSeat)
	return maxDis
}
