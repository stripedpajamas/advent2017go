package day13

import (
	"fmt"
	"strconv"
	"strings"
)

type Scanner struct {
	rang        []bool
	reverse     bool
	previousIdx int
	currentIdx  int
}

func makeState(input []string) []*Scanner {
	stateSize, err := strconv.Atoi(strings.Split(input[len(input)-1], ": ")[0])
	if err != nil {
		panic(err)
	}
	state := make([]*Scanner, stateSize+1)

	// parse input into initial state
	for _, line := range input {
		split := strings.Split(line, ": ")
		d, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		state[d] = &Scanner{
			rang:        make([]bool, r),
			previousIdx: 0,
			currentIdx:  1,
			reverse:     false,
		}
		state[d].rang[0] = true
	}

	return state
}

func printState(state []*Scanner) {
	fmt.Println("***")
	for _, s := range state {
		if s != nil {
			fmt.Println(s.rang)
		}
	}
	fmt.Println("***")
}

func resetState(state []*Scanner) {
	for _, v := range state {
		if v != nil {
			v.previousIdx = 0
			v.currentIdx = 1
			v.reverse = false
			v.rang[0] = true
			for i := 1; i < len(v.rang); i++ {
				v.rang[i] = false
			}
		}
	}
}

func incrementScanners(state []*Scanner) {
	for _, scanner := range state {
		if scanner != nil {
			// we are dealing with a valid scanner
			scannerLen := len(scanner.rang)
			scanner.rang[scanner.previousIdx] = false
			scanner.rang[scanner.currentIdx] = true

			if scanner.reverse {
				if scanner.currentIdx == 0 {
					// if we're at the first index, stop reversing
					scanner.currentIdx++
					scanner.previousIdx--
					scanner.reverse = false
				} else {
					scanner.currentIdx--
					scanner.previousIdx--
				}
			} else {
				if scanner.currentIdx == scannerLen-1 {
					// if we're at the last index, reverse
					scanner.currentIdx--
					scanner.previousIdx++
					scanner.reverse = true
				} else {
					scanner.currentIdx++
					scanner.previousIdx++
				}
			}
		}
	}
}

func goodDelay(delay int, invalidPeriods map[int][]int) bool {
	for period, answers := range invalidPeriods {
		for _, answer := range answers {
			if delay%period == answer {
				return false
			}
		}
	}
	return true
}

func Solve1(input []string) int {
	severity := 0
	state := makeState(input)

	// the main event loop
	// each scanner moves up or down within their range
	// and packet proceeds one index further and mark whether or not i'm caught

	for packetLoc, s := range state {
		if s != nil && s.rang[0] {
			// i am caught, update severity
			severity += packetLoc * len(s.rang)
		}

		// move scanners forward
		incrementScanners(state)
	}

	return severity
}

func Solve2(input []string) int {
	state := makeState(input)
	invalidPeriods := make(map[int][]int) // map period to delay%period to exclusions

	delay := 0

	for ; ; delay++ {
		// first check if we are in a delay we know will not work
		if !goodDelay(delay, invalidPeriods) {
			continue
		}
		caught := false
		clock := 0
		packetLoc := 0
		for ; ; clock++ {
			if clock >= delay {
				if state[packetLoc] != nil && state[packetLoc].rang[0] {
					// i am caught, try a new delay and update invalid periods
					rang := len(state[packetLoc].rang)
					var period int
					if rang > 1 {
						period = (rang * 2) - 2
					} else {
						period = rang
					}
					invalidPeriods[period] = append(invalidPeriods[period], delay%period)
					caught = true
					break
				} else if packetLoc == len(state)-1 {
					// i am at the end of the line without getting caught
					break
				}
				packetLoc++
			}

			// move scanners forward
			incrementScanners(state)
		}

		if caught == false {
			break
		}
		resetState(state)
	}

	return delay
}
