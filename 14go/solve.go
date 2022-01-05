package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// I think this solution is pretty nice but reads terrible due to all IFS since no default-ish maps

// instead of saving string, save pairs in dict.
// in each step, iterate over current pairs and generate the next steps pairs.

func solve1Alt(pairTable map[string][]string, charTable map[string]string, initString string, numSteps int) int {

	// After lookign at reddit, you can apparently just count all the letters in pairs and divide by 2..
	// oh well I count letter in a separate dict,
	// since after each pair expansion, only one new letter is created in final string. NN => NB BN - only B is new,
	charCount := make(map[string]int)

	// Generate initial pairs, count initial chars - just for ease sake last char is counted after loop
	currTemplate := make(map[string]int)
	runes := []rune(initString)
	for i := 0; i < len(initString)-1; i++ {
		first, second := string(runes[i]), string(runes[i+1])
		concat := fmt.Sprintf("%s%s", first, second)
		if _, ok := currTemplate[concat]; !ok {
			currTemplate[concat] = 1
		} else {
			currTemplate[concat] += 1
		}
		if _, ok := charCount[first]; !ok {
			charCount[first] = 1
		} else {
			charCount[first] += 1
		}
	}
	//
	// count steps !
	for i := 0; i < numSteps; i++ {
		nextTemplate := make(map[string]int)
		for k, v := range currTemplate {
			newPairs := pairTable[k]
			newChar := charTable[k]

			if _, ok := nextTemplate[newPairs[0]]; !ok {
				nextTemplate[newPairs[0]] = v
			} else {
				nextTemplate[newPairs[0]] += v
			}
			if _, ok := nextTemplate[newPairs[1]]; !ok {
				nextTemplate[newPairs[1]] = v
			} else {
				nextTemplate[newPairs[1]] += v
			}

			if _, ok := charCount[newChar]; !ok {
				charCount[newChar] = v
			} else {
				charCount[newChar] += v
			}
		}
		currTemplate = nextTemplate
	}

	// add the last character from original string
	charCount[string(runes[len(runes)-1])] += 1

	smallest := charCount["B"] // just start with whatever smallest / largest char
	largest := charCount["B"]
	for _, v := range charCount {
		if v < smallest {
			smallest = v
		}
		if v > largest {
			largest = v
		}
	}
	return (largest - smallest)
}

func main() {

	input, _ := ioutil.ReadFile("day14.txt")
	iSplit := strings.Split(string(input), "\n\n")
	table := make(map[string]string)
	for _, l := range strings.Split(iSplit[1], "\n") {
		splt := strings.Split(l, " -> ")
		table[splt[0]] = splt[1]
	}
	tablePairs := make(map[string][]string) // for smarter p2 algo, create tabl that returns pairs EG  NN => B now is NN => [NB, BN]
	for k, v := range table {
		s := strings.Split(k, "")
		concat1 := fmt.Sprintf("%s%s", s[0], v)
		concat2 := fmt.Sprintf("%s%s", v, s[1])
		tablePairs[k] = []string{concat1, concat2}
	}

	polyTemplate := iSplit[0]
	fmt.Printf("part 1 %d \n", solve1Alt(tablePairs, table, polyTemplate, 10))
	fmt.Printf("part 2 %d \n", solve1Alt(tablePairs, table, polyTemplate, 40))
}
