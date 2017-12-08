package day8

import (
	"bytes"
	"strconv"
)

func Solve1and2(input [][]byte) (int, int) {
	// b inc 5 if a > 1
	register := make(map[string]int)
	highest := 0

	for _, line := range input {
		split := bytes.Split(line, []byte(" "))
		// [b, inc,  5, if, a, <, 1]
		// first parse condition at idx 4,5,6
		var condition bool
		conditionTarget := register[string(split[4])]
		conditionReq, err := strconv.Atoi(string(split[6]))
		if err != nil {
			panic(err)
		}
		switch string(split[5]) {
		case "<":
			condition = conditionTarget < conditionReq
		case ">":
			condition = conditionTarget > conditionReq
		case "<=":
			condition = conditionTarget <= conditionReq
		case ">=":
			condition = conditionTarget >= conditionReq
		case "!=":
			condition = conditionTarget != conditionReq
		case "==":
			condition = conditionTarget == conditionReq
		}

		if condition {
			amount, err := strconv.Atoi(string(split[2]))
			if err != nil {
				panic(err)
			}
			if string(split[1]) == "dec" {
				register[string(split[0])] -= amount
			} else if string(split[1]) == "inc" {
				register[string(split[0])] += amount
			}
			if register[string(split[0])] > highest {
				highest = register[string(split[0])]
			}
		}
	}

	// after altering the register, find the highest value
	max := 0
	for _, val := range register {
		if val > max {
			max = val
		}
	}
	return max, highest
}
