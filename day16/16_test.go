package day16

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

var input []string
var dancers = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}

func init() {
	iBytes, err := ioutil.ReadFile("./16.txt")
	if err != nil {
		panic(err)
	}

	input = strings.Split(string(iBytes), ",")
}

func TestSolve1(t *testing.T) {
	testDancers := []string{
		"a", "b", "c", "d", "e",
	}
	testInput := []string{
		"s1",
		"x3/4",
		"pe/b",
		"s2",
	}

	result := Solve1(testDancers, testInput)

	if result != "dcbae" {
		fmt.Println("Wanted dcbae, got", result)
		t.Fail()
	}

	fmt.Println("Answer 1:", Solve1(dancers, input))
}

func TestSolve2(t *testing.T) {
	fmt.Println("Answer 2:", Solve2(dancers, input))
}
