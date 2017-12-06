package day2

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"
)

var realInput [][]int64

func init() {
	// read file
	inputBytes, err := ioutil.ReadFile("2.txt")
	if err != nil {
		panic(err)
	}
	// split on lines
	inputSplit := bytes.Split(inputBytes, []byte("\n"))

	// split on spaces
	input := make([][][]byte, len(inputSplit))
	for i, line := range inputSplit {
		input[i] = bytes.Split(line, []byte(" "))
	}

	realInput = make([][]int64, len(input))

	// convert each inner inner byte slice to string then parse that into an int
	for i, s := range input {
		realInput[i] = make([]int64, len(s))
		for j, x := range s {
			tmp, err := strconv.ParseInt(string(x), 10, 64)
			if err != nil {
				panic(err)
			}
			realInput[i][j] = tmp
		}
	}
}

func TestSolve1(t *testing.T) {
	/*
		5 1 9 5
		7 5 3
		2 4 6 8
	*/
	test1 := [][]int64{
		[]int64{5, 1, 9, 5},
		[]int64{7, 5, 3},
		[]int64{2, 4, 6, 8},
	}

	if Solve1(test1) != 18 {
		t.Fail()
	}

	fmt.Println("Answer 1: ", Solve1(realInput))
}

func TestSolve2(t *testing.T) {
	test1 := [][]int64{
		[]int64{5, 9, 2, 8},
		[]int64{9, 4, 7, 3},
		[]int64{3, 8, 6, 5},
	}

	if Solve2(test1) != 9 {
		t.Fail()
	}

	fmt.Println("Answer 2: ", Solve2(realInput))
}
