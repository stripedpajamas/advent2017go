package day17

func Solve1(input int) int {
	// make a 2018-length array
	buff := make([]int, 2018)

	buffLen := 1
	currentPosition := 0

	// loop through 2017 and add the values to buff
	for i := 1; i < 2018; i++ {
		stepsTaken := (currentPosition + input) % buffLen
		insertIdx := stepsTaken + 1
		copy(buff[insertIdx+1:], buff[insertIdx:])
		buff[insertIdx] = i // set the new value
		buffLen++
		currentPosition = insertIdx
	}

	return buff[(currentPosition+1)%2018]
}

func Solve2(input int) int {
	buff := make([]int, 2)

	buffLen := 1
	currentPosition := 0

	for i := 1; i < 50000001; i++ {
		stepsTaken := (currentPosition + input) % buffLen
		insertIdx := stepsTaken + 1
		// no need to mess with arrays at all since i'm only concerned about idx 1
		if insertIdx == 1 {
			buff[insertIdx] = i // set the new value
		}
		buffLen++
		currentPosition = insertIdx
	}

	return buff[1]
}
