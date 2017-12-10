package day9

func Solve1(input string) int {
	totalScore := 0 // because everything is wrapped in {}

	ignoreNext := false // flag for when we encounter a !
	inGarbage := false  // flag for when we are inside <...>
	inGroup := false    // flag for when we are inside {...}
	scoreModifier := 0  // goes up if you're deeper in, down if you come back out

	for _, x := range input {
		// fmt.Printf("Total: %d, char: %s, ignoreNext: %t, inGarbage: %t, inGroup: %t, modifier: %d\n", totalScore, string(x), ignoreNext, inGarbage, inGroup, scoreModifier)
		if ignoreNext {
			// previous char was a !
			ignoreNext = !ignoreNext
			continue
		}
		switch string(x) {
		case ">":
			if inGarbage {
				inGarbage = false
			}
		case "<":
			inGarbage = true
		case "{":
			if !inGarbage {
				inGroup = true
				scoreModifier++
			}
		case "}":
			if !inGarbage && inGroup {
				totalScore += scoreModifier
				scoreModifier--
			}
		case "!":
			ignoreNext = true
		}
	}
	return totalScore
}

func Solve2(input string) int {
	totalGarbage := 0

	ignoreNext := false // flag for when we encounter a !
	inGarbage := false  // flag for when we are inside <...>

	for _, x := range input {
		// fmt.Printf("Total: %d, char: %s, ignoreNext: %t, inGarbage: %t, inGroup: %t, modifier: %d\n", totalScore, string(x), ignoreNext, inGarbage, inGroup, scoreModifier)
		if ignoreNext {
			// previous char was a !
			ignoreNext = !ignoreNext
			continue
		}

		switch string(x) {
		case ">":
			if inGarbage {
				inGarbage = false
			}
		case "<":
			if !inGarbage {
				inGarbage = true
			} else {
				totalGarbage++
			}
		case "!":
			ignoreNext = true
		default:
			if inGarbage {
				totalGarbage++
			}
		}
	}
	return totalGarbage
}
