package day5

func Solve1(input []int) int {
	escapeSteps := 1
	previousIdx := 0
	currentIdx := 0
	current := input[0]
	escaped := false

	for !escaped {
		// starting at the first one and then we're gonna jump around
		currentIdx += current
		if currentIdx >= len(input) || currentIdx < 0 {
			escaped = true
			break
		}
		input[previousIdx]++
		current = input[currentIdx]
		previousIdx = currentIdx
		escapeSteps++
	}

	return escapeSteps
}

func Solve2(input []int) int {
	escapeSteps := 1
	previousIdx := 0
	currentIdx := 0
	current := input[0]
	escaped := false

	for !escaped {
		// starting at the first one and then we're gonna jump around
		currentIdx += current
		if currentIdx >= len(input) || currentIdx < 0 {
			escaped = true
			break
		}
		if input[previousIdx] >= 3 {
			input[previousIdx]--
		} else {
			input[previousIdx]++
		}
		current = input[currentIdx]
		previousIdx = currentIdx
		escapeSteps++
	}

	return escapeSteps
}
