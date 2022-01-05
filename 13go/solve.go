package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct{ x, y int }

func solve(grid [][]bool, instruction []point) {
	// parsing instructions as points, x=0,y=7 --- solve down on 7
	for insNum, i := range instruction {
		x, y := i.x, i.y
		nextGrid := gridCopy(grid, y, x)
		if x == 0 { // fold down
			for col := range nextGrid[0] {
				for i := range grid[y:] {
					val := grid[i+y][col]
					if val {
						nextGrid[len(nextGrid)-i][col] = val
					}
				}
			}
		} else { // fold right
			for row := range nextGrid {
				for i := range grid[row][x:] {
					val := grid[row][i+x]
					if val {
						nextGrid[row][len(nextGrid[0])-i] = val
					}
				}
			}
		}
		grid = nextGrid
		if insNum == 0 {
			countDots(grid)
		}
	}
	printGrid(grid)
}

// returns a copy of grid with specified size x,y
func gridCopy(grid [][]bool, y, x int) [][]bool { // TODO BETTER VAR NAME THAN X Y FOR COL ROW
	if x == 0 {
		x = len(grid[0])
	}
	if y == 0 {
		y = len(grid)
	}
	if x > len(grid[0]) || y > len(grid) {
		panic("x / y larger than grid, should not happen")
	}
	newGrid := [][]bool{}
	for row := 0; row < y; row++ {
		tempRow := make([]bool, x)
		copy(tempRow, grid[row][0:x])
		newGrid = append(newGrid, tempRow)
	}
	return newGrid
}

func main() {
	input, _ := ioutil.ReadFile("day13.txt")
	lines := strings.Split(string(input), "\n\n")
	gridLines, folds := lines[0], lines[1]

	// this parsing is a disaster just put in separate function instead
	// parse cordinates, save max x & y values
	foldInstructions := []point{}
	allCords := []point{}
	maxX, maxY := 0, 0
	for _, l := range strings.Split(gridLines, "\n") {
		splt := strings.Split(l, ",")
		x, y := splt[0], splt[1]
		xInt, _ := strconv.Atoi(x)
		yInt, _ := strconv.Atoi(y)
		if xInt > maxX {
			maxX = xInt
		}
		if yInt > maxY {
			maxY = yInt
		}
		allCords = append(allCords, point{xInt, yInt})
	}
	grid := make([][]bool, 0)
	for i := 0; i < maxY+1; i++ { // create grid ...
		grid = append(grid, make([]bool, maxX+1))
	}
	for _, p := range allCords { // insert cords !
		grid[p.y][p.x] = true
	}
	for _, fold := range strings.Split(folds, "\n") { // split coords into points
		splt := strings.Split(strings.TrimPrefix(fold, "fold along "), "=")
		dir, val := splt[0], splt[1]
		if dir == "x" {
			xVal, _ := strconv.Atoi(val)
			foldInstructions = append(foldInstructions, point{xVal, 0})
		} else if dir == "y" {
			yVal, _ := strconv.Atoi(val)
			foldInstructions = append(foldInstructions, point{0, yVal})
		} else {
			panic("error in fold-instruction parsing")
		}
	}

	//p1 747,
	//p2 :
	// .##..###..#..#.####.###...##..#..#.#..#.
	// #..#.#..#.#..#....#.#..#.#..#.#..#.#..#.
	// #..#.#..#.####...#..#..#.#....#..#.####.
	// ####.###..#..#..#...###..#....#..#.#..#.
	// #..#.#.#..#..#.#....#....#..#.#..#.#..#.
	// #..#.#..#.#..#.####.#.....##...##..#..#.

	solve(grid, foldInstructions)
}

// print the grid for the fans .. (and part 2 apparently)
func printGrid(grid [][]bool) {
	for row := range grid {
		str := ""
		for _, c := range grid[row] {
			if c {
				str = fmt.Sprintf("%s#", str)
			} else {
				str = fmt.Sprintf("%s.", str)
			}
		}
		println(str)
	}
}
func countDots(grid [][]bool) {
	total := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] {
				total++
			}
		}
	}
	println("total dots", total)
}
