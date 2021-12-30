package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct { // ugly struct  to have tuples ..
	row, col int
}

func resetFlashed(octiGrid [][]int, flashed map[point]bool) {
	for p := range flashed {
		octiGrid[p.row][p.col] = 0
	}
}

func tickAll(octiGrid [][]int) {
	for row := range octiGrid {
		for col := range octiGrid[row] {
			octiGrid[row][col]++
		}
	}
}

// not sure if there is a better way of doing this in golang
func notifyNeighs(octiGrid [][]int, currPoint point) {
	row, col := currPoint.row, currPoint.col
	if row > 0 { // above
		octiGrid[row-1][col]++
	}
	if row > 0 && col != 0 { // above left
		octiGrid[row-1][col-1]++
	}
	if row > 0 && col != len(octiGrid[row])-1 { // above right
		octiGrid[row-1][col+1]++
	}
	// sides
	if col != 0 {
		octiGrid[row][col-1]++
	}
	if col != len(octiGrid[row])-1 {
		octiGrid[row][col+1]++
	}
	// below
	if row != len(octiGrid)-1 {
		octiGrid[row+1][col]++
	}
	if row != len(octiGrid)-1 && col != 0 {
		octiGrid[row+1][col-1]++
	}
	if row != len(octiGrid)-1 && col != len(octiGrid[row])-1 {
		octiGrid[row+1][col+1]++
	}
}

func getFlashing(octiGrid [][]int, alreadyFlashed map[point]bool) []point {

	currFlashing := []point{}
	for row := range octiGrid {
		for col := range octiGrid[row] {
			_, hasFlashed := alreadyFlashed[point{row, col}]
			if octiGrid[row][col] > 9 && !hasFlashed {
				currFlashing = append(currFlashing, point{row, col})
			}
		}
	}
	return currFlashing
}

// return total flashes, which step first 100-octi-flash
// thatis p1-answer, p2-answer
func solve(octiGrid [][]int) (int, int) {
	currFlashing := make([]point, 0)
	alreadyFlashed := make(map[point]bool, 0)

	totalFlashed := 0 //p1 first 100 steps
	allFlashStep := 0 // p2 first step all 100 octi flash at the same time

	for i := 0; i < 500; i++ { // 500 steps is enough for p2
		// tick all
		tickAll(octiGrid)
		currFlashing = append(currFlashing, getFlashing(octiGrid, alreadyFlashed)...)
		for _, p := range currFlashing {
			alreadyFlashed[p] = true
		}

		for len(currFlashing) > 0 { // iterate all flashing octi this step
			currOct := currFlashing[len(currFlashing)-1] // pop new
			currFlashing = currFlashing[:len(currFlashing)-1]
			if i < 100 {
				totalFlashed++
			}

			notifyNeighs(octiGrid, currOct)
			currFlashing = append(currFlashing, getFlashing(octiGrid, alreadyFlashed)...) // get new!
			for _, p := range currFlashing {
				alreadyFlashed[p] = true
			}
		}
		if (len(alreadyFlashed)) >= 100 {
			allFlashStep = i + 1
			break
		}
		resetFlashed(octiGrid, alreadyFlashed)
		currFlashing = make([]point, 0)
		alreadyFlashed = make(map[point]bool, 0)
	}
	return totalFlashed, allFlashStep
}

func main() {

	input, _ := ioutil.ReadFile("day11.txt")
	lines := strings.Split(string(input), "\n")
	grid := make([][]int, 0)
	for _, r := range lines {
		row := make([]int, 0)
		for _, c := range r {
			d, _ := strconv.Atoi(string(c))
			row = append(row, d)
		}
		grid = append(grid, row)
	}

	p1, p2 := solve(grid)
	fmt.Printf("part1: %d\npart2 : %d \n", p1, p2)
}
