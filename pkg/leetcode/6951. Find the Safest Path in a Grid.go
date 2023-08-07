package leetcode

import "math"

func maximumSafenessFactor(grid [][]int) int {
	thieves := make([]Point, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			thieves = append(thieves, Point{
				x: i,
				y: j,
			})
		}
	}
	return maximumSafenessFactorBase(grid, 0, math.MaxInt32, 0, 0, thieves)
}

func maximumSafenessFactorBase(grid [][]int, safeDistance int, currentDistance int, x int, y int, thieves []Point) int {
	if x == len(grid)-1 && y == len(grid)-1 {
		return computeSafeDistance(thieves, x, y)
	}
	temDistance := computeSafeDistance(thieves, x, y)
	if temDistance <= safeDistance {
		return temDistance
	}
	currentDistance = min(currentDistance, temDistance)
	var distance int
	if x < len(grid)-1 {
		distance = max(distance, maximumSafenessFactorBase(grid, safeDistance, currentDistance, x+1, y, thieves))
	}
	if y < len(grid)-1 {
		distance = max(distance, maximumSafenessFactorBase(grid, safeDistance, currentDistance, x, y+1, thieves))
	}
	return min(currentDistance, distance)
}

func computeSafeDistance(thieves []Point, x int, y int) int {
	distance := math.MaxInt32
	for _, value := range thieves {
		distance = min(distance, abs(value.x-x)+abs(value.y-y))
	}
	return distance
}

type Point struct {
	x int
	y int
}
