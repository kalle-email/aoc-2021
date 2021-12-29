package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type line struct {
	from point
	to   point
}
type point struct {
	x, y int
}

func (l *line) lineLength() int {
	len := 0
	if l.from.x != l.to.x {
		len = l.from.x - l.to.x
	} else {
		len = l.from.y - l.to.y
	}
	if len < 0 {
		return -len
	}
	return len
}
func parseString(fromX, fromY, toX, toY string) line {
	fX, _ := strconv.Atoi(fromX)
	fY, _ := strconv.Atoi(fromY)
	tX, _ := strconv.Atoi(toX)
	tY, _ := strconv.Atoi(toY)
	line := line{point{fX, fY}, point{tX, tY}}
	return line
}
func filterStraights(lines []line) []line {
	straight := []line{}

	for _, l := range lines {
		if l.from.x == l.to.x || l.from.y == l.to.y {
			straight = append(straight, l)
		}
	}
	return straight
}

//
func solve(lines []line) int {

	m := make(map[point]int) // using map to count occurences.

	for _, l := range lines {
		xDir, yDir := 0, 0
		currX, currY := l.from.x, l.from.y
		lineLength := l.lineLength()
		// set directions!
		if l.from.x < l.to.x {
			xDir = 1
		} else if l.from.x > l.to.x {
			xDir = -1
		}
		if l.from.y < l.to.y {
			yDir = 1
		} else if l.from.y > l.to.y {
			yDir = -1
		}
		// iterate over line with the help of lineLength and directions.
		for i := 0; i <= lineLength; i++ {
			p := point{currX, currY}
			if num, ok := m[p]; ok {
				m[p] = num + 1
			} else {
				m[p] = 1
			}
			currX += xDir
			currY += yDir
		}
	}
	count := 0
	for _, v := range m {
		if v >= 2 {
			count += 1
		}
	}
	return count
}

func main() {
	input, _ := ioutil.ReadFile("day5.txt")
	s := strings.Split(string(input), "\n")
	allLines := []line{}
	for _, ss := range s {
		splt := strings.Split(ss, " -> ")
		from := strings.Split(splt[0], ",")
		to := strings.Split(splt[1], ",")
		l := parseString(from[0], from[1], to[0], to[1])
		allLines = append(allLines, l)
	}
	straightLines := filterStraights(allLines)

	fmt.Printf("part1: %d \n", solve(straightLines)) // part1
	fmt.Printf("part2: %d \n", solve(allLines))      // part1
}
