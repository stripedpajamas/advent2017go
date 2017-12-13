package day13

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

var input []string
var testInput []string

func init() {
	iBytes, err := ioutil.ReadFile("./13.txt")
	if err != nil {
		panic(err)
	}

	input = strings.Split(string(iBytes), "\n")

	iBytes, err = ioutil.ReadFile("./13_test.txt")
	if err != nil {
		panic(err)
	}

	testInput = strings.Split(string(iBytes), "\n")
}

func TestSolve1(t *testing.T) {
	result := Solve1(testInput)
	if result != 24 {
		fmt.Println("Wanted 24, got:", result)
		t.Fail()
	}

	fmt.Println("Answer 1:", Solve1(input))
}

func TestSolve2(t *testing.T) {
	result := Solve2(testInput)
	if result != 10 {
		fmt.Println("Wanted 10, got:", result)
		t.Fail()
	}

	fmt.Println("Answer 2:", Solve2(input))
}
