package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type pair struct { // inputPairs signal and nums !
	signals []string
	nums    []string
}

func solvep1(pairs []pair) int {
	nums := 0
	for _, p := range pairs {
		for _, n := range p.nums {
			switch len(n) {
			case 2, 4, 3, 7:
				nums++
			}
		}
	}
	return nums
}

// !NOT an acutal intersect!

// From strA perspective, check how many letters are not in strB. this makes this problem very simple
// e.g strA:abcf, strBb:abcde == 2,letters : ef not found in a, (even tho total !intersect is 3)
func numIntersection(a, b string) int {
	notOverlapping := 0
	for _, c := range a {
		if !strings.Contains(b, string(c)) {
			notOverlapping += 1
		}
	}
	return notOverlapping
}

// ugly helper, can rewrite this better
func calcNum(numStrings []string, numMap map[int]string) int {
	strMap := make(map[string]int)
	finalNum := 0
	for k, v := range numMap { // ugly reverse map
		strMap[v] = k
	}

	finalNum += strMap[numStrings[0]] * 1000 // 2 lazy to loop
	finalNum += strMap[numStrings[1]] * 100
	finalNum += strMap[numStrings[2]] * 10
	finalNum += strMap[numStrings[3]] * 1

	return finalNum
}

func solvep2(pairs []pair) int {

	total := 0
	for _, p := range pairs {
		numMap := make(map[int]string, 0)
		for _, n := range p.signals { // get all free nums
			switch len(n) {
			case 2:
				numMap[1] = n
			case 4:
				numMap[4] = n
			case 3:
				numMap[7] = n
			case 7:
				numMap[8] = n
			}
		}
		for _, n := range p.signals { // parse rest of nums
			switch len(n) {
			case 6: // 6,9,0
				if numIntersection(numMap[7], n) == 1 {
					numMap[6] = n
				} else if numIntersection(numMap[4], n) == 1 {
					numMap[0] = n
				} else {
					numMap[9] = n
				}
			case 5: // 5,3,2
				if numIntersection(numMap[7], n) == 1 && numIntersection(numMap[4], n) == 1 {
					numMap[5] = n
				} else if numIntersection(numMap[4], n) == 2 && numIntersection(numMap[7], n) == 1 {
					numMap[2] = n
				} else {
					numMap[3] = n
				}
			}
		}
		total += calcNum(p.nums, numMap)
	}

	return total
}

func SortString(w string) string { // FML, thanks SO
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {

	input, _ := ioutil.ReadFile("day8.txt")
	lines := strings.Split(string(input), "\n")
	allPairs := []pair{}
	for _, l := range lines {
		split := strings.Split(l, "|")
		p := pair{}

		// just noticed that a numbers corresponding signal is not ordered in the same way jesus christ thats 2 hours gone
		// sort string so I can pretend its a proper set ...
		for _, sig := range strings.Fields(split[0]) {
			sig = SortString(sig)
			p.signals = append(p.signals, sig)
		}
		for _, num := range strings.Fields(split[1]) {
			num = SortString(num)
			p.nums = append(p.nums, num)
		}
		allPairs = append(allPairs, p)
	}

	fmt.Printf("%+v", allPairs)
	fmt.Printf("part1 : %d \n", solvep1(allPairs))
	fmt.Printf("part2 : %d \n", solvep2(allPairs))
}
