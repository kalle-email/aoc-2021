package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"unicode"
)

func solvep1(g map[string][]string) int {

	queue := [][]string{{"start"}}
	allPaths := [][]string{}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		last := path[len(path)-1]
		if last == "end" {
			allPaths = append(allPaths, path)
		}
		for _, neighbor := range g[last] {
			if isLowerCase(neighbor) { // small cave, visit only once
				if !visited(path, neighbor) {
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, neighbor)
					queue = append(queue, newPath)
				}
			} else { // big cave, always isnert
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return len(allPaths)
}

// p2 iterative solution :
// when smallcave, push both regular path (no double visits) + path with this smallcave allowed twice on the stack.

func solvep2(g map[string][]string) int {

	queue := [][]string{{"start"}}
	allPaths := [][]string{}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		last := path[len(path)-1]
		if last == "end" {
			allPaths = append(allPaths, path)
		}

		for _, neighbor := range g[last] {
			if isLowerCase(neighbor) { // smallcave,
				if !visited(path, neighbor) {
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, neighbor)
					queue = append(queue, newPath)
				} else if !hasDouble(path) {
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, neighbor)
					queue = append(queue, newPath)
				}
			} else { // big cave, always isnert
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return len(allPaths)
}

// recursive solutions just for "fun"
func recP1(g map[string][]string, path []string) []string {

	last := path[len(path)-1]
	res := []string{}
	if last == "end" {
		return append(res, strings.Join(path, "->"))
	}
	for _, neighbor := range g[last] {
		if isLowerCase(neighbor) { // end/start/smallcave, visit only once
			if !visited(path, neighbor) {
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, neighbor)
				res = append(res, recP1(g, newPath)...)
			}
		} else { // big cave, always isnet
			newPath := make([]string, len(path))
			copy(newPath, path)
			newPath = append(newPath, neighbor)
			res = append(res, recP1(g, newPath)...)
		}
	}
	return res
}

// I think this is cleaner for p2
func p2Rec(g map[string][]string, path []string, visitedTwice bool) []string {

	last := path[len(path)-1]
	res := []string{}
	if last == "end" {
		return append(res, strings.Join(path, "->"))
	}
	for _, neighbor := range g[last] {
		if isLowerCase(neighbor) { // end/start/smallcave
			if !visited(path, neighbor) {
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, neighbor)
				res = append(res, p2Rec(g, newPath, visitedTwice)...)
			} else if !visitedTwice {
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, neighbor)
				res = append(res, p2Rec(g, newPath, true)...)
			}
		} else { // big cave, always insert
			newPath := make([]string, len(path))
			copy(newPath, path)
			newPath = append(newPath, neighbor)
			res = append(res, p2Rec(g, newPath, visitedTwice)...)
		}
	}
	return res

}

func main() {

	input, _ := ioutil.ReadFile("day12.txt")
	lines := strings.Split(string(input), "\n")
	coolG := make(map[string][]string)
	for _, l := range lines {
		split := strings.Split(l, "-")
		from, to := split[0], split[1]
		if _, seen := coolG[from]; !seen {
			coolG[from] = []string{}
		}
		if _, seen := coolG[to]; !seen {
			coolG[to] = []string{}
		}
		coolG[from] = append(coolG[from], to)
		if from != "start" { // instead of handling start and ends in function ill just do this
			coolG[to] = append(coolG[to], from)
		}
	}
	coolG["end"] = []string{} // see above :)

	start := time.Now()
	p1Iterative := solvep1(coolG)
	p1ItSpeed := time.Since(start)

	start = time.Now()
	p1Rec := recP1(coolG, []string{"start"})
	p1RRecSpeed := time.Since(start)

	start = time.Now()
	p2Rec := p2Rec(coolG, []string{"start"}, false)
	p2RecSpeed := time.Since(start)

	start = time.Now()
	p2It := solvep2(coolG)
	p2ItSpeed := time.Since(start)

	fmt.Printf("p1Iterative, %d, p1Rec %d\n", p1Iterative, len(p1Rec))
	fmt.Printf("it speed, %d, rec speed %d\n", p1ItSpeed, p1RRecSpeed)

	fmt.Printf("p2Iterative, %d, p2Rec %d\n", p2It, len(p2Rec))
	fmt.Printf("it speed, %d, rec speed %d\n", p2ItSpeed, p2RecSpeed)

}

// bunch of util func for all the fun experiements, too lazy 2 clean now

func visited(path []string, cave string) bool {
	for _, c := range path {
		if c == cave {
			return true
		}
	}
	return false
}

// check if cave is in path, if cave == doubleCave, check if doublecave is 2 times in path.
// lil helper, only to check if uppecase
func isLowerCase(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func hasDouble(l []string) bool {
	m := make(map[string]bool)
	for _, r := range l {
		if !isLowerCase(r) {
			continue
		}
		if _, ok := m[r]; !ok {
			m[r] = true
		} else {
			return true
		}
	}
	return false
}
