package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type p struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return -x
}

// PART 1
// highest y-zenith will reach the lowest Y value in the zone.
// All arcs will start at y=0 and then reach y=0 AGAIN on the way down.
// Using this property, its easy to see which value will pass the y=0 with the highest velocity and also be in the zone.
// e.g starting yVelocity 7 will be at y==0 at step 15 and will then have a velocity of -(startingY)-1
func p1(lowestY int) int {
	delta := getMaxY(lowestY)
	return calcTriangleNumber(delta) // triangleNumber of downward velocity == same as highest point reached.
}

// returns velocity at y=0 for starting point Y
func getMaxY(yFrom int) int {
	return abs(yFrom) - 1
}

// calc "triangle" number. (probably wrong terminology but its the same fromula .. )
// can be used to calc max X and Y value of x,y start.
func calcTriangleNumber(n int) int {
	return (n * (n + 1)) / 2
}

// p2 : for each step, count all X and Y starting-values that will be in the zone this step for SOME respective Y/X value. (that is - consider only one direction )
// Then, go through each step and create all pairs.
func p2(xFrom, xTo, yFrom, yTo int) int {
	// get minX velocity to reach target area this
	minX := getMinstartingX(xFrom)
	maxStep := 0 // save maxStep, that is - the last Y step that any starting Y be in the zone.

	// step -> starting values that is in the zone this step
	ymap := make(map[int][]int)
	for y := yFrom; y <= getMaxY(yFrom); y++ {
		ySteps := calcYInZone(y, yFrom, yTo)
		for _, s := range ySteps {
			if _, ok := ymap[s]; !ok {
				ymap[s] = make([]int, 0)
			}
			ymap[s] = append(ymap[s], y)
			if s > maxStep {
				maxStep = s
			}
		}
	}
	// use maxStep for x-startingValues that stay in the zone indefinetly.
	xmap := make(map[int][]int)
	for i := 0; i <= maxStep; i++ {
		xmap[i] = []int{}
	}

	for x := minX; x <= xTo; x++ {
		xSteps, indefinetly := calcXInZone(x, xFrom, xTo)
		if indefinetly { // stays in zone indefinetly, fill in all steps after first step which reached the zone.
			from := xSteps[0]
			for i := from; i <= maxStep; i++ {
				xmap[i] = append(xmap[i], x)
			}
		} else {
			for _, s := range xSteps {
				xmap[s] = append(xmap[s], x)
			}
		}
	}

	// create all combinations, since there will be overlaps, -- a pair can be in zone for more than 1 step -- save count in map.
	points := make(map[p]bool)
	for xStep, xCount := range xmap {
		for _, x := range xCount {
			for _, y := range ymap[xStep] {
				new := p{x, y}
				if _, ok := points[new]; !ok {
					points[new] = true
				}
			}
		}
	}
	return len(points)
}

// baby helper func to get lowestX value
func getMinstartingX(xFrom int) int {
	min := 0
	for i := 0; i < xFrom; i++ {
		min = calcTriangleNumber(i)
		if min >= xFrom {
			return i
		}
	}
	panic("no starting x found should not happen")
}

// from starting x value, return the step this trajectory is inside the target zone.
// Returns true if it remains in targetZone indefinetly.
// => [5], true  => startingX is in targetzone for all steps=>5
// => [4 5], false  => only in targetzone steps 4,5
func calcXInZone(xStart, xFrom, xTo int) ([]int, bool) {
	delta := xStart
	dist := 0
	steps := 0
	stepsInZone := []int{}
	for delta != 0 {
		dist += delta
		delta--
		steps++

		if dist >= xFrom && dist <= xTo {
			stepsInZone = append(stepsInZone, steps)
		}
	}

	if dist >= xFrom && dist <= xTo {
		return stepsInZone[0:1], true
	}
	return stepsInZone, false
}

// same as above but no indefinetly cases to consider
func calcYInZone(yStart, yFrom, yTo int) []int {
	delta := yStart
	height := 0
	steps := 0
	stepsInZone := []int{}
	for height > yFrom {
		height += delta
		delta--
		steps++

		if height >= yFrom && height <= yTo {
			stepsInZone = append(stepsInZone, steps)
		}
	}
	return stepsInZone
}

func main() {

	input, _ := ioutil.ReadFile("day17.txt")
	split := strings.Split(string(input), ", ")
	xSplit := strings.Split(split[0], "..")
	ySplit := strings.Split(split[1], "..")

	xFrom, _ := strconv.Atoi(strings.TrimPrefix(xSplit[0], "target area: x="))
	xTo, _ := strconv.Atoi(xSplit[1])
	yFrom, _ := strconv.Atoi(strings.TrimPrefix(ySplit[0], "y="))
	yTo, _ := strconv.Atoi(ySplit[1])

	fmt.Printf("P1 : %d \n", p1(yFrom))
	fmt.Printf("P2 : %d \n", p2(xFrom, xTo, yFrom, yTo))

}
