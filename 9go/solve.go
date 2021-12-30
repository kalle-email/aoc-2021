package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type point struct { // ugly struct  to have tuples ..
	row, col int
}

// check if neighbors is lower!
func isHighPoint(neigs []int, currPoint int) bool {
	for _, n := range neigs {
		if currPoint >= n {
			return false
		}
	}
	return true
}
func solvep2(lowPoints []point, grid [][]int) int {

	allBasins := make([]int, 0)
	for _, p := range lowPoints {
		allBasins = append(allBasins, calculateBasin(grid, p.row, p.col))
	}
	// take 3 largest
	sort.Slice(allBasins, func(i, j int) bool { return allBasins[i] > allBasins[j] })
	total := 1
	for i := range allBasins[:3] {
		total *= allBasins[i]
	}
	return total
}

// akta foer om  global visited!
func calculateBasin(grid [][]int, startRow, startCol int) int {
	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[0]))
	}
	stack := make([]point, 0)
	basin := make([]point, 0)
	stack = append(stack, point{startRow, startCol})

	// maybe easier to read to but visitedGrid top of loop if corner case are taken in  consideration.
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		currVal := grid[curr.row][curr.col]
		if currVal == 9 {
			continue
		}
		basin = append(basin, curr)

		if curr.col > 0 && !visited[curr.row][curr.col-1] {
			if grid[curr.row][curr.col-1] > currVal {
				stack = append(stack, point{curr.row, curr.col - 1})
				visited[curr.row][curr.col-1] = true
			}
		}
		if curr.col < len(grid[0])-1 && !visited[curr.row][curr.col+1] {
			if grid[curr.row][curr.col+1] > currVal {
				stack = append(stack, point{curr.row, curr.col + 1})
				visited[curr.row][curr.col+1] = true
			}
		}
		if curr.row > 0 && !visited[curr.row-1][curr.col] {
			if grid[curr.row-1][curr.col] > currVal {
				stack = append(stack, point{curr.row - 1, curr.col})
				visited[curr.row-1][curr.col] = true
			}
		}

		if curr.row < len(grid)-1 && !visited[curr.row+1][curr.col] {
			if grid[curr.row+1][curr.col] > currVal {
				stack = append(stack, point{curr.row + 1, curr.col})
				visited[curr.row+1][curr.col] = true
			}
		}
	}
	return len(basin)
}

func solvep1(grid [][]int) []point {

	allLowPoints := make([]point, 0)

	for row := range grid {
		for col := range grid[row] {
			neighbors := make([]int, 0)
			if col > 0 {
				neighbors = append(neighbors, grid[row][col-1])
			}
			if col < len(grid[row])-1 {
				neighbors = append(neighbors, grid[row][col+1])
			}
			if row > 0 {
				neighbors = append(neighbors, grid[row-1][col])
			}
			if row < len(grid)-1 {
				neighbors = append(neighbors, grid[row+1][col])
			}
			if isHighPoint(neighbors, grid[row][col]) {
				allLowPoints = append(allLowPoints, point{row, col})
			}
		}
	}
	return allLowPoints
}

func main() {

	input, _ := ioutil.ReadFile("day9.txt")
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

	allLowPoints := solvep1(grid)
	p1Answer := 0
	for _, p := range allLowPoints {
		p1Answer += grid[p.row][p.col] + 1
	}
	fmt.Printf("part1, %d \n", p1Answer)                    // 475
	fmt.Printf("part2, %d \n", solvep2(allLowPoints, grid)) // 1092012
}
