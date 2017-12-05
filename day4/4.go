package day4

import (
	"bytes"
	"sort"
)

func Solve1(input [][]byte) int {
	totalValid := 0

	for _, passphrase := range input {
		p := bytes.Split(passphrase, []byte(" "))
		valid := true
		for i, word := range p {
			for j, otherWord := range p {
				if i == j {
					// don't see if word == word
					continue
				}
				if bytes.Equal(word, otherWord) {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			totalValid++
		}
	}
	return totalValid
}

func Solve2(input [][]byte) int {
	totalValid := 0

	for _, passphrase := range input {
		p := bytes.Split(passphrase, []byte(" "))
		// sort all the words
		for _, w := range p {
			stringSlice := make([]string, len(w))
			// convert the bytes of each word to letters
			for idx, letter := range w {
				stringSlice[idx] = string(letter)
			}
			sort.Strings(stringSlice)
			// convert back to bytes
			for idx, letter := range stringSlice {
				for _, actualLetter := range letter {
					w[idx] = byte(actualLetter)
				}
			}
		}
		valid := true
		for i, word := range p {
			for j, otherWord := range p {
				if i == j {
					// don't see if word == word
					continue
				}
				if bytes.Equal(word, otherWord) {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			totalValid++
		}
	}
	return totalValid
}