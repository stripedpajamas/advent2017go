package day7

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

var inputTest [][]byte
var input [][]byte

func init() {
	inputBytes, err := ioutil.ReadFile("./7.txt")
	if err != nil {
		panic(err)
	}

	input = bytes.Split(inputBytes, []byte("\n"))

	inputBytes, err = ioutil.ReadFile("./7_test.txt")
	if err != nil {
		panic(err)
	}

	inputTest = bytes.Split(inputBytes, []byte("\n"))
}

func TestSolve1(t *testing.T) {
	treeTest := GetTree(inputTest)
	if Solve1(treeTest) != "tknk" {
		t.Fail()
	}

	tree := GetTree(input)
	fmt.Println("Answer 1:", Solve1(tree))
}

func TestSolve2(t *testing.T) {
	treeTest := GetTree(inputTest)
	if Solve2(treeTest) != 60 {
		t.Fail()
	}

	tree := GetTree(input)
	fmt.Println("Answer 2:", Solve2(tree))
}
