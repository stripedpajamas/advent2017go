package day11

import (
	"math"
	"strings"
)

func Solve1(input string) int {
	// split strings on comma
	split := strings.Split(input, ",")

	x, y := 0, 0 // start at the origin

	// iterate over them to get the coordinates
	for _, dir := range split {
		switch dir {
		case "ne":
			y++
		case "se":
			x++
		case "s":
			x++
			y--
		case "sw":
			y--
		case "nw":
			x--
		case "n":
			x--
			y++
		}
	}

	// the algorithm for measuring distance in this system is
	// if the signs of x and y are the same, abs(x+y), else max(abs(x), abs(y))
	if (x < 0 && y < 0) || (x >= 0 && y >= 0) {
		// signs are the same
		return int(math.Abs(float64(x + y)))
	}
	// signs are different
	return int(math.Max(math.Abs(float64(x)), math.Abs(float64(y))))
}

func Solve2(input string) int {
	// split strings on comma
	split := strings.Split(input, ",")

	x, y, dis, maxDis := 0, 0, 0, 0 // start at the origin

	// iterate over them to get the coordinates
	for _, dir := range split {
		switch dir {
		case "ne":
			y++
		case "se":
			x++
		case "s":
			x++
			y--
		case "sw":
			y--
		case "nw":
			x--
		case "n":
			x--
			y++
		}
		// the algorithm for measuring distance in this system is
		// if the signs of x and y are the same, abs(x+y), else max(abs(x), abs(y))
		if (x < 0 && y < 0) || (x >= 0 && y >= 0) {
			// signs are the same
			dis = int(math.Abs(float64(x + y)))
		} else {
			// signs are different
			dis = int(math.Max(math.Abs(float64(x)), math.Abs(float64(y))))
		}
		if dis > maxDis {
			maxDis = dis
		}
	}
	return maxDis
}
