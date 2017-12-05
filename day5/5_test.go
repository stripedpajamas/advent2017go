package day5

import (
	"io/ioutil"
	"bytes"
	"strconv"
	"testing"
	"fmt"
)

var input []int

func init() {
	inputBytes, err := ioutil.ReadFile("./5.txt")
	if err != nil {
		panic(err)
	}

	inputSplit := bytes.Split(inputBytes, []byte("\n"))
	input = make([]int, len(inputSplit))

	// parse the numbers
	for i, b := range inputSplit {
		tmp, err := strconv.Atoi(string(b))
		if err != nil {
			panic(err)
		}

		input[i] = tmp
	}
}

func TestSolve1(t *testing.T) {
	test1 := []int{
		0,
		3,
		0,
		1,
		-3,
	}

	if Solve1(test1) != 5 {
		t.Fail()
	}

	testingInput := make([]int, len(input))
	copy(testingInput, input)
	fmt.Println("Answer 1:", Solve1(testingInput))
}

func TestSolve2(t *testing.T) {
	test1 := []int{
		0,
		3,
		0,
		1,
		-3,
	}

	if Solve2(test1) != 10 {
		t.Fail()
	}

	testingInput := make([]int, len(input))
	copy(testingInput, input)
	fmt.Println("Answer 2:", Solve2(testingInput))
}
