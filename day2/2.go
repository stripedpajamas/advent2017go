package day2

func Solve1(input [][]int64) int64 {
	total := int64(0)
	for _, row := range input {
		max := row[0]
		min := row[0]
		for _, n := range row {
			if n > max {
				max = n
			}
			if n < min {
				min = n
			}
		}
		total += max - min
	}
	return total
}

func Solve2(input [][]int64) int64 {
	total := int64(0)
	for _, row := range input {
		for i, n := range row {
			for j, m := range row {
				if i == j {
					continue
				}
				if n%m == 0 {
					total += n / m
				} else if m%n == 0 {
					total += m / n
				}
			}
		}
	}
	return total / 2
}
