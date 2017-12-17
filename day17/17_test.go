package day17

import (
	"fmt"
	"testing"
)

func TestSolve1(t *testing.T) {
	result := Solve1(3)

	if result != 638 {
		fmt.Println("Wanted 638, got", result)
		t.Fail()
	}

	fmt.Println("Answer 1:", Solve1(345))
}

func TestSolve2(t *testing.T) {
	result := Solve2(3)

	if result != 1222153 {
		fmt.Println("Wanted 1222153, got", result)
		t.Fail()
	}
	fmt.Println("Answer 2:", Solve2(345))
}
