package day16

import (
	"fmt"
	"strconv"
	"strings"
)

var startValues = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
	"i": 8,
	"j": 9,
	"k": 10,
	"l": 11,
	"m": 12,
	"n": 13,
	"o": 14,
	"p": 15,
}

func makeDancersString(dancersMap map[string]int) string {
	dancers := make([]string, len(dancersMap))
	for dancer, idx := range dancersMap {
		dancers[idx] = dancer
	}
	return strings.Join(dancers, "")
}

func makeDancersMap(dancers []string) map[string]int {
	// maps dancer to index so we aren't messing with slices all day
	dancersMap := make(map[string]int)
	for idx, dancer := range dancers {
		dancersMap[dancer] = idx
	}
	return dancersMap
}

func handleSpin(dancers map[string]int, x int) {
	// s3 on abcde = cdeab
	dancersLen := len(dancers)
	for dancer, idx := range dancers {
		dancers[dancer] = (idx + x) % dancersLen
	}
}

func handleExchange(dancers map[string]int, aIdx, bIdx int) {
	// x3/4 on eabcd = eabdc
	var a, b string
	for dancer, idx := range dancers {
		if idx == aIdx {
			a = dancer
		}
		if idx == bIdx {
			b = dancer
		}
		if a != "" && b != "" {
			break
		}
	}
	dancers[a], dancers[b] = dancers[b], dancers[a]
}

func handlePartner(dancers map[string]int, a, b string) {
	// pe/b on eabdc = baedc
	dancers[a], dancers[b] = dancers[b], dancers[a]
}

func makeMovesList(moves []string) []func(map[string]int) {
	movesList := make([]func(map[string]int), len(moves))
	for idx, move := range moves {
		switch string(move[0]) {
		case "s": // s1
			spin, err := strconv.Atoi(string(move[1:]))
			if err != nil {
				panic(err)
			}
			movesList[idx] = func(dm map[string]int) {
				//fmt.Println("Running handle spin with:", spin)
				handleSpin(dm, spin)
			}
		case "x": // x1/2
			// the slash could be 2 digits in
			slash := strings.Index(move, "/")
			aIdx, err := strconv.Atoi(string(move[1:slash]))
			if err != nil {
				panic(err)
			}
			bIdx, err := strconv.Atoi(string(move[slash+1:]))
			if err != nil {
				panic(err)
			}
			movesList[idx] = func(dm map[string]int) {
				//fmt.Println("Running handle exchange with:", aIdx, bIdx)
				handleExchange(dm, aIdx, bIdx)
			}
		case "p": // pa/b
			slash := strings.Index(move, "/")
			a := string(move[1:slash])
			b := string(move[slash+1:])
			movesList[idx] = func(dm map[string]int) {
				//fmt.Println("Running handle partner with:", a, b)
				handlePartner(dm, a, b)
			}
		}
	}

	return movesList
}

func dance(dancersMap map[string]int, movesList []func(map[string]int)) {
	for _, move := range movesList {
		move(dancersMap)
	}
}

func checkMap(dancersMap map[string]int) bool {
	for dancer, _ := range dancersMap {
		if dancersMap[dancer] != startValues[dancer] {
			return false
		}
	}
	return true
}

func Solve1(dancers, moves []string) string {
	dancersMap := makeDancersMap(dancers)
	movesList := makeMovesList(moves)

	dance(dancersMap, movesList)

	return makeDancersString(dancersMap)
}

func Solve2(dancers, moves []string) string {
	// dance a billion times
	dancersMap := makeDancersMap(dancers)
	movesList := makeMovesList(moves)

	// found that map was going back to starting position every 60 iterations
	// so only had to run this 40 times for the answer :)
	for i := 0; i < 40; i++ {
		dance(dancersMap, movesList)
		if checkMap(dancersMap) {
			fmt.Println("Map is back to starting position. i = ", i)
		}
	}

	return makeDancersString(dancersMap)
}
