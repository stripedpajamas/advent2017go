package day3

import (
	"testing"
	"fmt"
)

func TestFindMDInSpiral(t *testing.T) {
	/*
	Data from square 1 is carried 0 steps, since it's at the access port.
	Data from square 12 is carried 3 steps, such as: down, left, left.
	Data from square 23 is carried only 2 steps: up twice.
	Data from square 1024 must be carried 31 steps.
	 */

	 one := FindMDInSpiral(0, 0)
	 twelve := FindMDInSpiral(2, 1)
	 twentyThree := FindMDInSpiral(0, -2)
	 tenTwentyFour := FindMDInSpiral(-15, 16)

	 if one != 0 || twelve != 3 || twentyThree != 2 || tenTwentyFour != 31 {
	 	t.Fail()
	 }

	 fmt.Println("Answer 1:", FindMDInSpiral(181, -257))
}

func TestMakeSpiral(t *testing.T) {
	s := MakeSpiral(300000)
	x, y := s[0][0], s[0][1]
	if x != 0 || y != 0 {
		t.Error("For 0, 0 I got", x, y)
	}

	x, y = s[11][0], s[11][1]
	if x != 2 || y != 1 {
		t.Error("For 2, 1 I got", x, y)
	}

	x, y = s[22][0], s[22][1]
	if x != 0 || y != -2 {
		t.Error("For 0, -2 I got", x, y)
	}
}
