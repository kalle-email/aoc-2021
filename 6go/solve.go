package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func solve(startingFish []int, days int) int {
	fishArray := startingFish
	var nextDay []int

	for i := 0; i < days; i++ {
		nextDay = make([]int, 9) // prolly bad way of clearing array :- )
		for j := 7; j >= 0; j-- {
			nextDay[j] = fishArray[j+1]
		}
		nextDay[8] = fishArray[0]
		nextDay[6] += fishArray[0]
		fishArray = nextDay
	}

	// sum it up
	totalFishes := 0
	for _, f := range fishArray {
		totalFishes += f
	}
	return totalFishes
}

func main() {

	input, _ := ioutil.ReadFile("day6.txt")
	splt := strings.Split(string(input), ",")
	startingFish := make([]int, 9)
	for _, n := range splt {
		num, _ := strconv.Atoi(n)
		startingFish[num]++
	}
	fmt.Printf("part1: %d \n", solve(startingFish, 80))  // 350917
	fmt.Printf("part2: %d \n", solve(startingFish, 256)) //1592918715629
}
