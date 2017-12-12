package day12

import (
	"strconv"
	"strings"
)

func GetTree(input []string) map[int][]int {
	tree := make(map[int][]int)

	for _, line := range input {
		// gross input. first we'll split on <->
		values := strings.Split(line, " <-> ")
		node, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		childrenStrings := strings.Split(values[1], ", ")

		children := make([]int, len(childrenStrings))
		for i, child := range childrenStrings {
			c, err := strconv.Atoi(child)
			if err != nil {
				panic(err)
			}
			children[i] = c
		}

		// we now have node and children we can put in our map
		tree[node] = children
	}

	return tree
}

func CountChildren(root int, tree map[int][]int, total map[int]bool) {
	// start at root
	total[root] = true

	for _, child := range tree[root] {
		if !total[child] {
			CountChildren(child, tree, total)
		}
	}
}

func Solve1(input []string) int {
	tree := GetTree(input)
	total := make(map[int]bool)
	CountChildren(0, tree, total)

	return len(total)
}

func Solve2(input []string) int {
	tree := GetTree(input)
	totalGroups := 0
	for root, _ := range tree {
		group := make(map[int]bool)
		CountChildren(root, tree, group)

		// remove this entire group from the original tree
		for k, _ := range group {
			delete(tree, k)
		}

		totalGroups++
	}

	return totalGroups
}
