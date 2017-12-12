package day12

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

var input []string
var testInput []string

func init() {
	iBytes, err := ioutil.ReadFile("./12.txt")
	if err != nil {
		panic(err)
	}

	input = strings.Split(string(iBytes), "\n")

	iBytes, err = ioutil.ReadFile("./12_test.txt")
	if err != nil {
		panic(err)
	}

	testInput = strings.Split(string(iBytes), "\n")
}

func TestSolve1(t *testing.T) {
	result := Solve1(testInput)

	if result != 6 {
		fmt.Println("Wanted 6, got", result)
		t.Fail()
	}

	fmt.Println("Answer 1:", Solve1(input))
}

func TestSolve2(t *testing.T) {
	result := Solve2(testInput)

	if result != 2 {
		fmt.Println("Wanted 2, got", result)
		t.Fail()
	}

	fmt.Println("Answer 2:", Solve2(input))
}
