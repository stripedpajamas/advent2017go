package day9

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var input string

func init() {
	iBytes, err := ioutil.ReadFile("./9.txt")
	if err != nil {
		panic(err)
	}

	input = string(iBytes)
}

func TestSolve1(t *testing.T) {
	tests := []string{
		"{}",
		"{{{}}}",
		"{{},{}}",
		"{{{},{},{{}}}}",
		"{<a>,<a>,<a>,<a>}",
		"{{<ab>},{<ab>},{<ab>},{<ab>}}",
		"{{<!!>},{<!!>},{<!!>},{<!!>}}",
		"{{<a!>},{<a!>},{<a!>},{<ab>}}",
	}

	results := make([]int, len(tests))

	for i, test := range tests {
		results[i] = Solve1(test)
	}

	want := []int{1, 6, 5, 16, 1, 9, 9, 3}

	for i, result := range results {
		if result != want[i] {
			fmt.Printf("Test: %s, wanted %d, got %d\n", tests[i], want[i], result)
			t.Fail()
		}
	}

	fmt.Println("Answer 1:", Solve1(input))
}

func TestSolve2(t *testing.T) {
	tests := []string{
		"<>",
		"<random characters>",
		"<<<<>",
		"<{!>}>",
		"<!!>",
		"<!!!>>",
		"<{o\"i!a,<{i<a>",
		"{<{o\"i!a,<{i<a>}",
		"{{<a!>},{<a!>},{<a!>},{<ab>}}", //17
	}

	results := make([]int, len(tests))

	for i, test := range tests {
		results[i] = Solve2(test)
	}

	want := []int{0, 17, 3, 2, 0, 0, 10, 10, 17}

	for i, result := range results {
		if result != want[i] {
			fmt.Printf("Test: %s, wanted %d, got %d\n", tests[i], want[i], result)
			t.Fail()
		}
	}

	fmt.Println("Answer 2:", Solve2(input))
}
