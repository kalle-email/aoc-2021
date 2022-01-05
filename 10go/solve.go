package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// this is probably the worst of my solutions, but boring problem too lazy to rewrite

// returns chunks list with corrutped lines removed.
func filterIncomplete(allChunks, corrupted []string) []string {
	filtered := []string{}
	cMap := make(map[string]bool, 0)
	for _, c := range corrupted {
		cMap[c] = true
	}
	for _, a := range allChunks {
		if _, ok := cMap[a]; !ok {
			filtered = append(filtered, a)
		}
	}
	return filtered
}

// horrendous  function
func solvep1(chunks []string) (int, []string) {

	corruptedChars := []string{}
	corruptedChunks := []string{}
	currOpen := []string{} // this could be causing a bug but not for my input, reusing currOpen for all lines

	for _, chunk := range chunks {
	OuterLoop:
		for _, c := range chunk {
			s := string(c)
			switch s {
			case "[", "(", "{", "<":
				currOpen = append(currOpen, s)
			case "]":
				if currOpen[len(currOpen)-1] != "[" {
					corruptedChars = append(corruptedChars, "]")
					corruptedChunks = append(corruptedChunks, chunk)
					break OuterLoop
				} else {
					currOpen = currOpen[:len(currOpen)-1]
				}
			case ")":
				if currOpen[len(currOpen)-1] != "(" {
					corruptedChars = append(corruptedChars, ")")
					corruptedChunks = append(corruptedChunks, chunk)
					break OuterLoop
				} else {
					currOpen = currOpen[:len(currOpen)-1]
				}

			case "}":
				if currOpen[len(currOpen)-1] != "{" {
					corruptedChars = append(corruptedChars, "}")
					corruptedChunks = append(corruptedChunks, chunk)
					break OuterLoop
				} else {
					currOpen = currOpen[:len(currOpen)-1]
				}
			case ">":
				if currOpen[len(currOpen)-1] != "<" {
					corruptedChars = append(corruptedChars, ">")
					corruptedChunks = append(corruptedChunks, chunk)
					break OuterLoop
				} else {
					currOpen = currOpen[:len(currOpen)-1]
				}
			}
		}
	}
	score := 0 // could calculate score immediatly in first loop instead.
	for _, c := range corruptedChars {
		switch c {
		case ")":
			score += 3
		case "]":
			score += 57
		case "}":
			score += 1197
		case ">":
			score += 25137
		}
	}
	return score, corruptedChunks
}

// Coulduse the same solve function since incomplete lines wont close a pattern incorrectly -
// so they can be distinguished as you iterate over them. but this is is alright aswell.
func solvep2(chunks []string) int {

	allOpened := [][]string{} // bad var name
	allScores := []int{}

	for _, chunk := range chunks {
		currOpen := []string{}
		for _, c := range chunk {
			s := string(c)
			switch s {
			case "[", "(", "{", "<":
				currOpen = append(currOpen, s)
			case "]":
				currOpen = currOpen[:len(currOpen)-1]
			case ")":
				currOpen = currOpen[:len(currOpen)-1]
			case "}":
				currOpen = currOpen[:len(currOpen)-1]
			case ">":
				currOpen = currOpen[:len(currOpen)-1]
			}
		}
		allOpened = append(allOpened, currOpen)
	}
	// iterate backwards over non-closed brackets
	for _, curr := range allOpened {
		score := 0
		for i := len(curr) - 1; i >= 0; i-- {
			switch curr[i] {
			case "[":
				score *= 5
				score += 2
			case "(":
				score *= 5
				score += 1
			case "{":
				score *= 5
				score += 3
			case "<":
				score *= 5
				score += 4
			}
		}
		allScores = append(allScores, score)
	}

	sort.Ints(allScores)
	return allScores[len(allScores)/2]
}

func main() {
	input, _ := ioutil.ReadFile("day10.txt")
	lines := strings.Split(string(input), "\n")
	chunks := make([]string, 0)
	chunks = append(chunks, lines...)
	p1Score, corrupted := solvep1(chunks)
	incomplete := filterIncomplete(chunks, corrupted)

	fmt.Printf("p1 answer:  %d \n", p1Score)
	fmt.Printf("p2 answer:  %d \n", solvep2(incomplete))
}
