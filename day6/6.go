package day6

import (
	"crypto/sha1"
)

func Solve1(input []byte) int {
	cycles := 0
	seen := make(map[[20]byte]bool, len(input))
	state := make([]byte, len(input))
	copy(state, input)

	for !seen[sha1.Sum(state)] {
		// make this state as seen in the map
		seen[sha1.Sum(state)] = true

		// find the highest val
		var max byte
		maxIdx := 0
		for i, x := range state {
			if x > max {
				max = x
				maxIdx = i
			}
		}

		// empty max value across state
		state[maxIdx] = 0
		for i := (maxIdx + 1) % len(state); max > 0; i++ {
			state[i%len(state)]++
			max--
		}

		// count the cycles
		cycles++
	}

	return cycles
}

func Solve2(input []byte) int {
	cycles := 0
	seen := make(map[[20]byte]int, len(input))
	state := make([]byte, len(input))
	copy(state, input)

	for seen[sha1.Sum(state)] == 0 {
		// make this state as seen in the map
		seen[sha1.Sum(state)] = cycles

		// find the highest val
		var max byte
		maxIdx := 0
		for i, x := range state {
			if x > max {
				max = x
				maxIdx = i
			}
		}

		// empty max value across state
		state[maxIdx] = 0
		for i := (maxIdx + 1) % len(state); max > 0; i++ {
			state[i%len(state)]++
			max--
		}

		// count the cycles
		cycles++
	}

	// we want to know how many cycles have happened between the current iteration
	// and the iteration that produced the duplicate hash

	return cycles - seen[sha1.Sum(state)]
}
