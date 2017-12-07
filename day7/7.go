package day7

import (
	"strconv"
	"strings"
)

type Program struct {
	Name        string
	TrueWeight  int
	Weight      int
	BranchNames []string
	Branches    map[string]*Program
	HasParent   bool
}

func GetTree(input [][]byte) map[string]*Program {
	// we have the lines split on \n, but we need to parse them
	// ghgsrqm (122) -> jajds, mkhupy
	// ktlj (57)
	// are the two types of lines we will encounter

	Tree := make(map[string]*Program)

	for _, line := range input {
		s := strings.Split(string(line), " ")
		// [ghgsrqm, (122), ->, jajds, mkhupy]
		// [ktlf, (57)]
		name := s[0]
		weight, err := strconv.Atoi(strings.Trim(s[1], "()"))
		if err != nil {
			panic(err)
		}

		Tree[name] = &Program{
			Name:   name,
			Weight: weight,
		}

		if len(s) > 2 {
			// s has branches
			Tree[name].BranchNames = make([]string, len(s)-len(s[:3]))
			Tree[name].Branches = make(map[string]*Program)

			for i, branch := range s[3:] {
				Tree[name].BranchNames[i] = strings.Trim(branch, ", ")
			}
		}
	}

	// mapping branch names to structs
	// also noting who could be the root
	for _, root := range Tree {
		if len(root.BranchNames) > 0 {
			// working with a root
			for _, branch := range root.BranchNames {
				root.Branches[branch] = Tree[branch]
				Tree[branch].HasParent = true
			}
		}
	}

	for _, root := range Tree {
		// add true weight property
		root.TrueWeight = FindTrueWeight(root)
	}

	return Tree
}

func FindTrueWeight(root *Program) int {
	total := root.Weight
	for _, branch := range root.Branches {
		total += FindTrueWeight(branch)
	}
	return total
}

func FindUnbalance(root *Program) string {
	// should return the name of an unbalanced branch or something

	// make a map of true weights to how many branches have that weight
	branchWeights := make(map[int]int)

	for _, info := range root.Branches {
		branchWeights[info.TrueWeight]++
	}

	// find the true weight that only occurred once
	for weight, occ := range branchWeights {
		if occ == 1 {
			// find the one
			for name, info := range root.Branches {
				if info.TrueWeight == weight {
					return FindUnbalance(root.Branches[name])
				}
			}
		}
	}

	return root.Name
}

func Solve1(tree map[string]*Program) string {
	// loop through and find the one that doesn't have a parent
	for name, info := range tree {
		if !info.HasParent {
			return name
		}
	}

	return "Not found"
}

func Solve2(tree map[string]*Program) int {
	root := Solve1(tree)
	unbalance := FindUnbalance(tree[root])

	var rootOfUnbalance string

	// find the unbalanced node's root and its branches and their true weights
	for _, possibleRoot := range tree {
		for _, branch := range possibleRoot.Branches {
			if branch.Name == unbalance {
				rootOfUnbalance = possibleRoot.Name
				break
			}
		}
		if rootOfUnbalance != "" {
			break
		}
	}

	var desiredWeight int

	for _, branch := range tree[rootOfUnbalance].Branches {
		if branch.Name == unbalance {
			continue
		} else {
			desiredWeight = branch.TrueWeight
			break
		}
	}

	difference := desiredWeight - tree[unbalance].TrueWeight

	return tree[unbalance].Weight + difference
}
