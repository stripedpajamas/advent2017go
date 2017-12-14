package day14

import (
	"fmt"
	"testing"
)

var input string = "ljoxqyyw"

func TestSolve1(t *testing.T) {
	test := "flqrgnkx"
	result := Solve1(test)

	if result != 8108 {
		fmt.Println("Wanted 8108, got", result)
		t.Fail()
	}

	fmt.Println("Answer 1:", Solve1(input))
}

func TestSolve2(t *testing.T) {
	test := hashList("flqrgnkx")
	result := Solve2(test)

	if result != 1242 {
		fmt.Println("Wanted 1242, got", result)
		t.Fail()
	}

	hashedInput := hashList(input)

	fmt.Println("Answer 2:", Solve2(hashedInput))
}
