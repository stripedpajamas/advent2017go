package day6

import (
	"fmt"
	"testing"
)

var input []byte = []byte{10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6}
var input2 []byte = []byte{5, 1, 10, 0, 1, 7, 13, 14, 3, 12, 8, 10, 7, 12, 0, 6}

func TestSolve1(t *testing.T) {
	test1 := []byte{0, 2, 7, 0}

	if Solve1(test1) != 5 {
		t.Error("Wanted 5, got", Solve1(test1))
	}

	fmt.Println("Answer 1:", Solve1(input))
}

func TestSolve2(t *testing.T) {
	test1 := []byte{0, 2, 7, 0}

	if Solve2(test1) != 4 {
		t.Error("Wanted 4, got", Solve2(test1))
	}

	fmt.Println("Answer 2:", Solve2(input))
}
