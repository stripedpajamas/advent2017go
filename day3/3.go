package day3

import (
	"math"
)

func FindMDInSpiral(x, y int) int {
	// finds distance between (x, y) and (0, 0)
	// the taxicab distance between (p1, p2) and (q1, q2) =
	// |p1 - q1| + |p2 - q2|
	// |p1 - 0| + |p2 - 0| in our case
	// |p1| + |p2|
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func MakeSpiral(limit int) [][]int {
	// make spiral
	// https://stackoverflow.com/a/398302/7703857
	output := make([][]int, limit)
	x, y := 0, 0
	dx := 0
	dy := -1

	for i := 0; i < limit; i++ {
		if (-limit/2 < x && x <= limit/2) && (-limit/2 < y && y <= limit/2) {
			output[i] = []int{x, y}
		}
		if x == y || (x < 0 && x == -y) || (x > 0 && x == 1-y) {
			dx, dy = -dy, dx
		}
		x, y = x+dx, y+dy
	}
	return output
}
