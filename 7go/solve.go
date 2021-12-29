package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// my dumb bruteforce solution, according to internet, math says solution is close to avg of crabs,
// i thought the cool part was to notice the triangle number ..
func solvep2(crabs []int) int {

	minFuelUsed := 99999999999999999
	currFuelUsed := 0
	max := crabs[len(crabs)-1] // crab-list is sorted

	for target := 0; target <= max; target++ {
		currFuelUsed = 0
		for _, c := range crabs {
			dist := abs(c - target)
			fuel := ((dist * (dist + 1)) / 2)
			currFuelUsed += fuel
			// could return early here if currFuel > minfuelused, nvm avg solution is the coolest regardless
		}
		if currFuelUsed < minFuelUsed {
			minFuelUsed = currFuelUsed
		}
	}
	return minFuelUsed
}

func solvep1(crabs []int) int {

	// crab-list is sorted
	middleCrab := crabs[(len(crabs) / 2)]
	totalFuel := 0
	for _, c := range crabs {
		totalFuel += abs(c - middleCrab)
	}
	return totalFuel
}

func main() {

	input, _ := ioutil.ReadFile("day7.txt")
	splt := strings.Split(string(input), ",")
	crabs := make([]int, 0)
	for _, n := range splt {
		num, _ := strconv.Atoi(n)
		crabs = append(crabs, num)
	}

	sort.Slice(crabs, func(i, j int) bool { return crabs[i] < crabs[j] })
	fmt.Printf("part1: %d \n", solvep1(crabs))
	fmt.Printf("part2: %d \n", solvep2(crabs))
}
