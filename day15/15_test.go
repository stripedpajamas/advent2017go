package day15

import (
	"fmt"
	"testing"
)

func TestSolve1(t *testing.T) {
	result := Solve1(65, 8921)
	if result != 588 {
		fmt.Println("Wanted 588, got", result)
		t.Fail()
	}

	fmt.Println("Answer 1:", Solve1(699, 124))
}

func TestSolve2(t *testing.T) {
	result := Solve2(65, 8921)
	if result != 309 {
		fmt.Println("Wanted 309, got", result)
		t.Fail()
	}

	fmt.Println("Answer 2:", Solve2(699, 124))
}
