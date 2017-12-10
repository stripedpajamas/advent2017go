package day10

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

var input []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255}
var inputLengths []int = []int{63, 144, 180, 149, 1, 255, 167, 84, 125, 65, 188, 0, 2, 254, 229, 24}
var inputBLengths []byte = []byte{63, 144, 180, 149, 1, 255, 167, 84, 125, 65, 188, 0, 2, 254, 229, 24}

func TestSolve1(t *testing.T) {
	testList := []int{0, 1, 2, 3, 4}
	testLengths := []int{3, 4, 1, 5}
	if Solve1(testList, testLengths) != 12 {
		t.Fail()
	}

	fmt.Println("Answer 1:", Solve1(input, inputLengths))
}

func TestSolve2(t *testing.T) {
	puzzleInput := []byte("63,144,180,149,1,255,167,84,125,65,188,0,2,254,229,24")
	tests := [][]byte{
		[]byte{},
		[]byte("AoC 2017"),
		[]byte("1,2,3"),
		[]byte("1,2,4"),
	}

	results := make([][]byte, len(tests))

	for i, test := range tests {
		results[i] = Solve2(test)
	}

	wantHex := []string{
		"a2582a3a0e66e6e86e3812dcb672a272",
		"33efeb34ea91902bb2f59c9920caa6cd",
		"3efbe78a8d82f29979031a4aa0b16a9d",
		"63960835bcdc130f0b66d7ff4f6a5a8e",
	}

	want := make([][]byte, len(wantHex))

	for i, h := range wantHex {
		tmp, err := hex.DecodeString(h)
		if err != nil {
			panic(err)
		}
		want[i] = tmp
	}

	for i, result := range results {
		if !bytes.Equal(result, want[i]) {
			fmt.Printf("Test: %v, wanted %x, got %x\n", tests[i], want[i], result)
			t.Fail()
		}
	}

	fmt.Printf("Answer 2: %x\n", Solve2(puzzleInput))
}
