package day8

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

var input [][]byte
var inputTest [][]byte

func init() {
	testBytes, err := ioutil.ReadFile("./8_test.txt")
	if err != nil {
		panic(err)
	}
	inputTest = bytes.Split(testBytes, []byte("\n"))

	realBytes, err := ioutil.ReadFile("./8.txt")
	if err != nil {
		panic(err)
	}
	input = bytes.Split(realBytes, []byte("\n"))
}

func TestSolve1and2(t *testing.T) {
	one, two := Solve1and2(inputTest)
	if one != 1 {
		t.Fail()
	}

	if two != 10 {
		t.Fail()
	}

	one, two = Solve1and2(input)
	fmt.Println("Answer 1:", one)
	fmt.Println("Answer 2:", two)
}
