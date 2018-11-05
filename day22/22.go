package day22

import (
	"fmt"
	"strings"
)

type Node struct {
	infected bool
}

type Location struct {
	row int
	col int
}

// directions:
// 0: up; 1: right; 2: down; 3: left
// so that we can use addition/subtraction mod 4
// to turn
type Virus struct {
	direction int
	row       int
	col       int
}

func (n *Node) toggleState() {
	n.infected = !n.infected
}

func (v *Virus) turnRight() {
	v.direction = (v.direction + 1) % 4
}

func (v *Virus) turnLeft() {
	v.direction = (4 + v.direction - 1) % 4
}

func (v *Virus) move() {
	switch v.direction {
	case 0:
		v.row--
	case 1:
		v.col++
	case 2:
		v.row++
	case 3:
		v.col--
	}
}

func (v *Virus) burst(current *Node) int {
	willInfect := false
	if current.infected {
		v.turnRight()
	} else {
		willInfect = true
		v.turnLeft()
	}
	current.toggleState()
	v.move()

	if willInfect {
		return 1
	}
	return 0
}

func parseInput(s string) (Location, map[Location]*Node) {
	split := strings.Split(s, "\n")
	output := make(map[Location]*Node)
	for rIdx, row := range split {
		for cIdx, val := range row {
			loc := Location{row: rIdx, col: cIdx}
			output[loc] = &Node{infected: val == '#'}
		}
	}
	return Location{row: len(split) / 2, col: len(split[0]) / 2}, output
}

func printGrid(grid map[Location]*Node) {
	for k, v := range grid {
		fmt.Printf("%d:%d => %v\n", k.row, k.col, v.infected)
	}
}

func Solve1(input map[Location]*Node, start Location, bursts int) int {
	// always start in the middle
	virus := new(Virus)
	virus.row = start.row
	virus.col = start.col

	total := 0
	for bursts > 0 {
		// if the virus is sitting on a node that doesn't
		// exist in the map already; create it
		if _, found := input[Location{virus.row, virus.col}]; !found {
			input[Location{virus.row, virus.col}] = new(Node)
		}
		total += virus.burst(input[Location{virus.row, virus.col}])
		bursts--
	}
	return total
}

// part 2
// ==================================================
// states:
// 0: clean; 1: weakened; 2: infected; 3: flagged
// addition % 4 is the flow of states
type NodeV2 struct {
	state int
}

type VirusV2 struct {
	Virus
}

func (n *NodeV2) toggleState() {
	n.state = (n.state + 1) % 4
}

func (v *VirusV2) reverse() {
	v.direction = (v.direction + 2) % 4
}

func (v *VirusV2) burst(current *NodeV2) int {
	willInfect := false
	switch current.state {
	case 0: // clean
		v.turnLeft()
	case 1: // weakened
		// do not move
		willInfect = true
	case 2: // infected
		v.turnRight()
	case 3: // flagged
		v.reverse()
	}
	current.toggleState()
	v.move()

	if willInfect {
		return 1
	}
	return 0
}

func parseInputV2(s string) (Location, map[Location]*NodeV2) {
	split := strings.Split(s, "\n")
	output := make(map[Location]*NodeV2)
	for rIdx, row := range split {
		for cIdx, val := range row {
			loc := Location{row: rIdx, col: cIdx}
			state := 0
			if val == '#' {
				state = 2
			}
			output[loc] = &NodeV2{state}
		}
	}
	return Location{row: len(split) / 2, col: len(split[0]) / 2}, output
}

func Solve2(input map[Location]*NodeV2, start Location, bursts int) int {
	// always start in the middle
	virus := new(VirusV2)
	virus.row = start.row
	virus.col = start.col

	total := 0
	for bursts > 0 {
		// if the virus is sitting on a node that doesn't
		// exist in the map already; create it
		if _, found := input[Location{virus.row, virus.col}]; !found {
			input[Location{virus.row, virus.col}] = new(NodeV2)
		}
		total += virus.burst(input[Location{virus.row, virus.col}])
		bursts--
	}
	return total
}
