package day18

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

var input []string
var inputTest []string

func init() {
	iBytes, err := ioutil.ReadFile("./18.txt")
	if err != nil {
		panic(err)
	}

	input = strings.Split(string(iBytes), "\n")

	iBytes, err = ioutil.ReadFile("./18_test.txt")
	if err != nil {
		panic(err)
	}

	inputTest = strings.Split(string(iBytes), "\n")
}

func TestSolve1(t *testing.T) {
	result := Solve1(inputTest)

	if result != 4 {
		fmt.Println("Wanted 4, got", result)
		t.Fail()
	}
	fmt.Println("Answer 1:", Solve1(input))
}

func TestSolve2(t *testing.T) {
	test := []string{
		"snd 1",
		"snd 2",
		"snd p",
		"rcv a",
		"rcv b",
		"rcv c",
		"rcv d",
	}

	result := Solve2(test)
	if result != 3 {
		fmt.Println("Wanted 3, got", result)
		t.Fail()
	}

	fmt.Println("Answer 2:", Solve2(input))
}
