package day11

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var input string

func init() {
	iBytes, err := ioutil.ReadFile("./11.txt")
	if err != nil {
		panic(err)
	}

	input = string(iBytes)
}

func TestSolve1(t *testing.T) {
	tests := []string{
		"ne,ne,ne",
		"ne,ne,sw,sw",
		"ne,ne,s,s",
		"se,sw,se,sw,sw",
	}

	results := make([]int, len(tests))

	for i, test := range tests {
		results[i] = Solve1(test)
	}

	want := []int{3, 0, 2, 3}

	for i, result := range results {
		if result != want[i] {
			fmt.Printf("Test: %s, wanted: %d, got: %d\n", tests[i], want[i], result)
			t.Fail()
		}
	}

	fmt.Println("Answer 1:", Solve1(input))
}

func TestSolve2(t *testing.T) {
	tests := []string{
		"ne,ne,ne",
		"ne,ne,sw,sw",
		"ne,ne,s,s",
		"se,sw,se,sw,sw",
	}

	results := make([]int, len(tests))

	for i, test := range tests {
		results[i] = Solve2(test)
	}

	want := []int{3, 2, 2, 3}

	for i, result := range results {
		if result != want[i] {
			fmt.Printf("Test: %s, wanted: %d, got: %d\n", tests[i], want[i], result)
			t.Fail()
		}
	}

	fmt.Println("Answer 2:", Solve2(input))
}
