package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	row, col int
}

// get neighbors of grid entry
func getNeighbors(grid [][]int, currPoint point) []point {
	neighs := make([]point, 0)
	row, col := currPoint.row, currPoint.col
	if row > 0 {
		neighs = append(neighs, point{row - 1, col})
	}
	if row < len(grid)-1 {
		neighs = append(neighs, point{row + 1, col})
	}
	if col > 0 {
		neighs = append(neighs, point{row, col - 1})
	}
	if col < len(grid[row])-1 {
		neighs = append(neighs, point{row, col + 1})
	}
	return neighs
}

//djikstras for both solutions
func solve(grid [][]int) int {
	currCost := make(map[point]int)
	for row := range grid {
		for col := range grid[row] {
			currP := point{row, col}
			currCost[currP] = 999999999999
		}
	}
	currCost[point{0, 0}] = 0
	visited := make(map[point]bool)
	pq := []point{}

	visited[point{0, 0}] = true // using a visited map inst. of putting all elements in prioqueue.
	pq = pqInsert(point{0, 0}, pq, currCost)

	for len(pq) > 0 {
		curr, newPq := pqPop(pq)
		pq = newPq
		if curr.row == len(grid)-1 && curr.col == len(grid[0])-1 {
			return currCost[curr]
		}
		neigh := getNeighbors(grid, curr)
		for _, n := range neigh {
			alt := currCost[curr] + grid[n.row][n.col]
			if alt < currCost[n] {
				currCost[n] = alt
			}
			if _, ok := visited[n]; !ok {
				newPq := pqInsert(n, pq, currCost)
				visited[n] = true
				pq = newPq
			}
		}
	}
	return -1
}

func main() {
	grid := make([][]int, 0)
	input, _ := ioutil.ReadFile("day15.txt")
	for row, line := range strings.Split(string(input), "\n") {
		split := strings.Split(line, "")
		grid = append(grid, []int{})
		for _, s := range split {
			num, _ := strconv.Atoi(s)
			grid[row] = append(grid[row], num)
		}
	}
	myExpanded := getP2Board(grid)
	fmt.Printf("p1: %d \n", solve(grid))
	fmt.Printf("p2: %d \n", solve(myExpanded))
}

// my priorityqueue hack. simply a sorted slice with a "pop" function.
func pqPop(pq []point) (point, []point) {
	pop := pq[0]
	return pop, pq[1:]
}

func pqInsert(insert point, pq []point, dist map[point]int) []point {

	for i := range pq {
		if dist[insert] < dist[pq[i]] {
			pq = append(pq[:i+1], pq[i:]...)
			pq[i] = insert
			return pq
		}
	}
	return append(pq, insert)
}

func getP2Board(grid [][]int) [][]int {

	grid = extendDown(grid)
	grid = extendRight(grid)
	return grid
}

// helper func to copy all rows of a grid and append them with all cols+=1 or reset to 1.
func newChunk(grid [][]int) [][]int {
	newChunk := [][]int{}
	for row := range grid {
		newRow := []int{}
		for col := range grid[row] {
			curr := grid[row][col]
			curr += 1
			if curr == 10 {
				newRow = append(newRow, 1)
			} else {
				newRow = append(newRow, curr)
			}
		}
		newChunk = append(newChunk, newRow)
	}
	return newChunk
}

func extendDown(start [][]int) [][]int {

	fullGrid := start
	currChunk := start
	for i := 0; i < 4; i++ {
		nextChunk := newChunk(currChunk)
		fullGrid = append(fullGrid, nextChunk...)
		currChunk = nextChunk
	}
	return fullGrid
}

// extends to the right 4 times.
func extendRight(grid [][]int) [][]int {

	for row := range grid {
		nextRow := grid[row]
		for i := 0; i < 4; i++ { // expand 4 times
			currRow := nextRow
			nextRow = []int{}
			for _, curr := range currRow {
				curr += 1
				if curr == 10 {
					nextRow = append(nextRow, 1)
				} else {
					nextRow = append(nextRow, curr)
				}
			}
			grid[row] = append(grid[row], nextRow...)
		}
	}
	return grid
}
