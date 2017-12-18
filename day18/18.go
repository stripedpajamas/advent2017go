package day18

import (
	"strconv"
	"strings"
	"time"
)

func Solve1(input []string) int {
	register := make(map[string]int)
	soundsPlayed := make([]int, 0)
	recovered := make([]int, 0)
	line := 0

	// execution loop
outer:
	for ; line < len(input); line++ {
		instruction := strings.Split(input[line], " ")

		switch instruction[0] {
		case "snd":
			soundsPlayed = append(soundsPlayed, register[instruction[1]])
		case "set":
			var newVal int
			if v, err := strconv.Atoi(instruction[2]); err == nil {
				newVal = v
			} else {
				newVal = register[instruction[2]]
			}
			register[instruction[1]] = newVal
		case "add":
			var newVal int
			if v, err := strconv.Atoi(instruction[2]); err == nil {
				newVal = v
			} else {
				newVal = register[instruction[2]]
			}
			register[instruction[1]] += newVal
		case "mul":
			var newVal int
			if v, err := strconv.Atoi(instruction[2]); err == nil {
				newVal = v
			} else {
				newVal = register[instruction[2]]
			}
			register[instruction[1]] *= newVal
		case "mod":
			var newVal int
			if v, err := strconv.Atoi(instruction[2]); err == nil {
				newVal = v
			} else {
				newVal = register[instruction[2]]
			}
			register[instruction[1]] %= newVal
		case "rcv":
			if register[instruction[1]] != 0 {
				recovered = append(recovered, soundsPlayed[len(soundsPlayed)-1])
				break outer
			}
		case "jgz":
			var conditionVal int
			var newVal int
			if v, err := strconv.Atoi(instruction[2]); err == nil {
				newVal = v
			} else {
				newVal = register[instruction[2]]
			}
			if v, err := strconv.Atoi(instruction[1]); err == nil {
				conditionVal = v
			} else {
				conditionVal = register[instruction[1]]
			}
			if conditionVal > 0 {
				line += newVal - 1
			}
		}
	}

	return recovered[0]
}

func Program(programId int, recvChan, sendChan, done chan int, input []string) int {
	register := make(map[string]int)
	register["p"] = programId
	line := 0
	sent := 0

	// execution loop
outer:
	for ; line < len(input); line++ {
		instruction := strings.Split(input[line], " ")

		switch instruction[0] {
		case "snd":
			var newVal int
			if v, err := strconv.Atoi(instruction[1]); err == nil {
				newVal = v
			} else {
				newVal = register[instruction[1]]
			}
			sendChan <- newVal
			sent++
		case "set":
			var newVal int
			if v, err := strconv.Atoi(instruction[2]); err == nil {
				newVal = v
			} else {
				newVal = register[instruction[2]]
			}
			register[instruction[1]] = newVal
		case "add":
			var newVal int
			if v, err := strconv.Atoi(instruction[2]); err == nil {
				newVal = v
			} else {
				newVal = register[instruction[2]]
			}
			register[instruction[1]] += newVal
		case "mul":
			var newVal int
			if v, err := strconv.Atoi(instruction[2]); err == nil {
				newVal = v
			} else {
				newVal = register[instruction[2]]
			}
			register[instruction[1]] *= newVal
		case "mod":
			var newVal int
			if v, err := strconv.Atoi(instruction[2]); err == nil {
				newVal = v
			} else {
				newVal = register[instruction[2]]
			}
			register[instruction[1]] %= newVal
		case "rcv":
			select {
			case newVal := <-recvChan:
				register[instruction[1]] = newVal
			case <-time.After(time.Second * 4):
				break outer
			}
		case "jgz":
			var conditionVal int
			var newVal int
			if v, err := strconv.Atoi(instruction[2]); err == nil {
				newVal = v
			} else {
				newVal = register[instruction[2]]
			}
			if v, err := strconv.Atoi(instruction[1]); err == nil {
				conditionVal = v
			} else {
				conditionVal = register[instruction[1]]
			}
			if conditionVal > 0 {
				line += newVal - 1
			}
		}
	}

	done <- sent

	return sent
}

func Solve2(input []string) int {
	io0 := make(chan int, 100)
	io1 := make(chan int, 100)

	result0 := make(chan int)
	result1 := make(chan int)

	go Program(0, io0, io1, result0, input)
	go Program(1, io1, io0, result1, input)

	// we only care about how many sends Program 1 did
	<-result0
	retVal := <-result1

	return retVal
}
