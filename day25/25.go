package day25

type Inst struct {
	write     int
	move      int
	nextState string
}

func CountOnes(tape map[int]int) int {
	total := 0
	for _, x := range tape {
		total += x
	}
	return total
}

/*
  input is a map of StateKey -> [2]Inst
*/
func Solve1(startingState string, iterations int, instructions map[string][]Inst) int {
	tape := make(map[int]int)
	cursor := 0
	state := startingState
	var nextCursor, nextVal int
	for iterations > 0 {
		nextVal = instructions[state][tape[cursor]].write
		nextCursor += instructions[state][tape[cursor]].move
		state = instructions[state][tape[cursor]].nextState
		tape[cursor] = nextVal
		cursor = nextCursor
		iterations--
	}
	return CountOnes(tape)
}
