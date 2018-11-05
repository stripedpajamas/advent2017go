package day22

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var testData = "..#\n#..\n..."
var input string

func init() {
	iBytes, err := ioutil.ReadFile("./22.txt")
	if err != nil {
		panic(err)
	}

	input = string(iBytes)
}

func TestSolve1(t *testing.T) {
	testStart, testInput := parseInput(testData)
	output := Solve1(testInput, testStart, 70)
	if output != 41 {
		t.Fail()
	}

	start, realInput := parseInput(input)
	fmt.Println("Answer 1:", Solve1(realInput, start, 10000))
}

func TestSolve2(t *testing.T) {
	testStart, testInput := parseInputV2(testData)
	output := Solve2(testInput, testStart, 100)
	if output != 26 {
		t.Fail()
	}

	start, realInput := parseInputV2(input)
	fmt.Println("Answer 1:", Solve2(realInput, start, 10000000))
}
