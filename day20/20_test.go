package day20

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

var input [][]byte

func init() {
	iBytes, err := ioutil.ReadFile("./20.txt")
	if err != nil {
		panic(err)
	}

	input = bytes.Split(iBytes, []byte("\n"))
}

func TestSolve1(t *testing.T) {
	test := [][]byte{
		[]byte("p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>"),
		[]byte("p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>"),
	}
	result := Solve1(parseInput(test))
	if result != 0 {
		t.Fail()
	}

	fmt.Println("Answer 1:", Solve1(parseInput(input)))
}

func TestSolve2(t *testing.T) {
	test := [][]byte{
		[]byte("p=<-6,0,0>, v=<3,0,0>, a=<0,0,0>"),
		[]byte("p=<-4,0,0>, v=<2,0,0>, a=<0,0,0>"),
		[]byte("p=<-2,0,0>, v=<1,0,0>, a=<0,0,0>"),
		[]byte("p=<3,0,0>, v=<-1,0,0>, a=<0,0,0>"),
		[]byte("p=<0,-8,0>, v=<0,2,0>, a=<0,0,0>"),
		[]byte("p=<0,-12,0>, v=<0,3,0>, a=<0,0,0>"),

		[]byte("p=<1,-6,0>, v=<0,1,0>, a=<0,1,0>"),
		[]byte("p=<1,-12,0>, v=<0,3,0>, a=<0,1,0>"),
	}
	result := Solve2(parseInput(test))

	if result != 1 {
		fmt.Println("Wanted 1, got", result)
		t.Fail()
	}

	fmt.Println("Answer 2:", Solve2(parseInput(input)))
}
