package day19

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

var input [][]byte
var inputTest [][]byte

func init() {
	iBytes, err := ioutil.ReadFile("./19.txt")
	if err != nil {
		panic(err)
	}

	input = bytes.Split(iBytes, []byte("\n"))

	iBytes, err = ioutil.ReadFile("./19_test.txt")
	if err != nil {
		panic(err)
	}

	inputTest = bytes.Split(iBytes, []byte("\n"))
}

func TestSolve2(t *testing.T) {
	result := Solve2(inputTest)

	if result != 38 {
		fmt.Println("Wanted 38, got", result)
		t.Fail()
	}
	fmt.Println("=========================")
	fmt.Println("Answer 1:", Solve2(input))
}
