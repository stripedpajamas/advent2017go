package day4

import (
	"io/ioutil"
	"bytes"
	"testing"
	"fmt"
)

var input [][]byte

func init() {
	inputBytes, err := ioutil.ReadFile("./4.txt")
	if err != nil {
		panic(err)
	}

	input = bytes.Split(inputBytes, []byte("\n"))
}

func TestSolve1(t *testing.T) {
	/*
	aa bb cc dd ee is valid.
	aa bb cc dd aa is not valid - the word aa appears more than once.
	aa bb cc dd aaa is valid - aa and aaa count as different words.
	 */

	test1 := [][]byte{
		[]byte("aa bb cc dd ee"),
	}
	test2 := [][]byte{
		[]byte("aa bb cc dd aa"),
	}
	test3 := [][]byte{
		[]byte("aa bb cc dd aaa"),
	}

	if Solve1(test1) != 1 {
		t.Fail()
	}
	if Solve1(test2) != 0 {
		t.Fail()
	}
	if Solve1(test3) != 1 {
		t.Fail()
	}

	fmt.Println("Answer 1:", Solve1(input))
}

func TestSolve2(t *testing.T) {
	/*
	abcde fghij is a valid passphrase.
	abcde xyz ecdab is not valid - the letters from the third word can be rearranged to form the first word.
	a ab abc abd abf abj is a valid passphrase, because all letters need to be used when forming another word.
	iiii oiii ooii oooi oooo is valid.
	oiii ioii iioi iiio is not valid - any of these words can be rearranged to form any other word.
	 */

	 test1 := [][]byte{
	 	[]byte("abcde fghij"),
	 }
	 test2 := [][]byte{
	 	[]byte("abcde xyz ecdab"),
	 }
	 test3 := [][]byte{
	 	[]byte("a ab abc abd abf abj"),
	 }
	 test4 := [][]byte{
	 	[]byte("iiii oiii ooii oooi oooo"),
	 }
	 test5 := [][]byte{
	 	[]byte("oiii ioii iioi iiio"),
	 }

	 if Solve2(test1) != 1 || Solve2(test3) != 1 || Solve2(test4) != 1 {
	 	t.Fail()
	 }
	 if Solve2(test2) != 0 || Solve2(test5) != 0 {
	 	t.Fail()
	 }

	 fmt.Println("Answer 2:", Solve2(input))
}
