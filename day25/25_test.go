package day25

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

var instructionSet map[string][]Inst

func init() {
	inputBytes, err := ioutil.ReadFile("./25.txt")
	if err != nil {
		panic(err)
	}

	inputSplit := bytes.Split(inputBytes, []byte("\n"))

	instructionSet = make(map[string][]Inst)

	var currentState string
	var currentVal int
	for _, line := range inputSplit {
		s := string(line)
		if strings.Contains(s, "In state") {
			currentState = s[9:10]
			instructionSet[currentState] = make([]Inst, 2)
			continue
		}
		if strings.Contains(s, "value is") {
			currentVal, _ = strconv.Atoi(s[26:27])
			instructionSet[currentState][currentVal] = Inst{}
			continue
		}
		if strings.Contains(s, "Write the value") {
			value, _ := strconv.Atoi(s[22:23])
			instructionSet[currentState][currentVal].write = value
			continue
		}
		if strings.Contains(s, "Move one slot") {
			var value int
			switch s[27:31] {
			case "left":
				value = -1
			case "righ":
				value = 1
			}
			instructionSet[currentState][currentVal].move = value
			continue
		}
		if strings.Contains(s, "Continue with") {
			nextState := s[26:27]
			instructionSet[currentState][currentVal].nextState = nextState
			continue
		}
	}
}

func TestSolve1(t *testing.T) {
	fmt.Println("Answer 1:", Solve1("A", 12683008, instructionSet))
}
