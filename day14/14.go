package day14

import (
	"fmt"
	"math/bits"
	"strconv"

	"github.com/stripedpajamas/advent/day10"
)

func hashList(input string) [][]byte {
	// build our list of inputs to hash
	inputs := make([][]byte, 128)
	for i := int64(0); i < 128; i++ {
		current := []byte(input + "-")
		current = strconv.AppendInt(current, i, 10)
		inputs[i] = current
	}

	// loop through our inputs list and knot hash them
	outputs := make([][]byte, 128)
	for i, in := range inputs {
		outputs[i] = day10.Solve2(in)
	}

	return outputs
}

func cleanMap(badRegionId, goodRegionId int, hashMap map[string]int) {
	// replace all instances of badRegionId with goodRegionId in the hash map
	for key, id := range hashMap {
		if id == badRegionId {
			hashMap[key] = goodRegionId
		}
	}
}

func Solve1(input string) int {
	outputs := hashList(input)

	// count the ones in each resulting hash
	total := 0
	for _, hash := range outputs {
		for _, b := range hash {
			total += bits.OnesCount8(b)
		}
	}

	return total
}

func Solve2(input [][]byte) int {
	// build out a string list of binary rep
	strs := make([]string, 128)
	for i, hash := range input {
		var tmp string
		for _, b := range hash {
			tmp2 := strconv.FormatInt(int64(b), 2)
			tmp2 = "00000000" + tmp2
			tmp2 = tmp2[len(tmp2)-8:]
			tmp += tmp2
		}
		strs[i] = tmp
	}

	// to count up the regions we will make a map of hashes of coordinates to a bool
	// the bool is true if those coordinates are part of a region

	regionIdCounter := 1
	hashMap := make(map[string]int)
	for row, line := range strs {
		for col, val := range line {
			if string(val) == "1" {
				ownHash := fmt.Sprintf("%x", day10.Solve2([]byte{byte(row), byte(col)}))
				var left, up string // l, u
				if col != 0 {
					// we are not on left edge so we can check item to the left
					if string(line[col-1]) == "1" {
						left = fmt.Sprintf("%x", day10.Solve2([]byte{byte(row), byte(col - 1)}))
					}
				}
				if row != 0 {
					// we are not on the top so we can check the item above
					if string(strs[row-1][col]) == "1" {
						up = fmt.Sprintf("%x", day10.Solve2([]byte{byte(row - 1), byte(col)}))
					}
				}

				if left != "" && up != "" {
					// if there is a 1 to the left and a 1 above
					// we want to use the lower value and then replace every instance in our map with that
					if hashMap[left] > hashMap[up] {
						cleanMap(hashMap[left], hashMap[up], hashMap)
						hashMap[ownHash] = hashMap[up]
					} else if hashMap[up] > hashMap[left] {
						cleanMap(hashMap[up], hashMap[left], hashMap)
						hashMap[ownHash] = hashMap[left]
					} else {
						// they're equal, so no cleaning necessary
						hashMap[ownHash] = hashMap[left]
					}
				} else if left != "" {
					hashMap[ownHash] = hashMap[left]
				} else if up != "" {
					hashMap[ownHash] = hashMap[up]
				} else {
					// neither left or up are 1s, so this is a new region
					hashMap[ownHash] = regionIdCounter
					regionIdCounter++
				}
			}
		}
	}

	uniqueRegions := make(map[int]bool)
	// count up regions
	for _, id := range hashMap {
		uniqueRegions[id] = true
	}

	return len(uniqueRegions)
}
