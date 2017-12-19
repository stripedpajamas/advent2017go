package day19

func Solve2(input [][]byte) int {
	total := 0
	//row, col := 0, 0

	for rowIdx, row := range input {
		rowTotal := 0
		for colIdx, x := range row {
			switch string(x) {
			case "-":
				rowTotal++
				// check for this:
				// |
				// -
				// |
				if rowIdx > 0 && rowIdx < len(input)-1 {
					// we aren't in the first or last row
					if colIdx < len(input[rowIdx-1]) && colIdx < len(input[rowIdx+1]) {
						above := string(input[rowIdx-1][colIdx])
						below := string(input[rowIdx+1][colIdx])
						switch above {
						case "|", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
							switch below {
							case "|", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
								rowTotal++
							}
						}
					}
				}
			case "|":
				rowTotal++
				// check for this:
				// -|-
				if colIdx > 0 && colIdx < len(row)-1 {
					// we aren't in the first or last column
					left := string(row[colIdx-1])
					right := string(row[colIdx+1])
					switch left {
					case "-", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
						switch right {
						case "-", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
							rowTotal++
						}
					}
				}
			case "+", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
				rowTotal++
			}
		}
		//fmt.Printf("Row %d has %d weirdos\n", rowIdx, rowTotal)
		total += rowTotal
	}
	return total
}
