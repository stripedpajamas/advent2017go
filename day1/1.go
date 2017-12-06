package day1

func Solve1(input []int64) int64 {
	var total int64

	inputLen := len(input)

	for i, x := range input {
		if x == input[(i+1)%inputLen] {
			total += x
		}
	}

	return total
}

func Solve2(input []int64) int64 {
	var total int64

	inputLen := len(input)
	half := inputLen / 2

	for i, x := range input {
		if x == input[(i+half)%inputLen] {
			total += x
		}
	}

	return total
}
