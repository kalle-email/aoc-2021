package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type board struct {
	num    [][]int
	marked [][]bool
}

func (b *board) calcScore(finalNum int) int {
	unmarkedSum := 0
	for row := range b.marked {
		for col := 0; col < len(b.num[0]); col++ {
			if !(b.marked[row][col]) {
				unmarkedSum += b.num[row][col]
			}
		}
	}
	return unmarkedSum * finalNum
}

// mark board and check if bingo! Returns true if bingo
func (b *board) markBoard(num int) bool {
	for row := range b.num {
		for col := 0; col < len(b.num[0]); col++ {
			if b.num[row][col] == num {
				b.marked[row][col] = true
				if b.checkBingo(row, col) {
					return true
				}
			}
		}
	}
	return false
}
func (b *board) checkBingo(row, col int) bool {
	rowBingo, colBingo := true, true

	for i := 0; i < len(b.num[0]); i++ { // row bingo
		if !(b.marked[row][i]) {
			rowBingo = false
		}
	}
	for i := 0; i < len(b.num); i++ { // col bingo
		if !(b.marked[i][col]) {
			colBingo = false
		}
	}
	return rowBingo || colBingo
}

// since p2 needs to loop through all boards regardless, solve both parts in one func.
func solve(currBoards []board, nums []int) (int, int) {

	solvedBoards := []int{} // score of solved boads
	var b board
	for _, num := range nums {
		for i := len(currBoards) - 1; i >= 0; i-- { // iterate backwards so we can remove ezpz
			b = currBoards[i]
			if done := b.markBoard(num); done {
				solvedBoards = append(solvedBoards, b.calcScore(num))
				currBoards = append(currBoards[:i], currBoards[i+1:]...)
			}
		}
	}
	return solvedBoards[0], solvedBoards[len(solvedBoards)-1]
}

func main() {
	input, _ := ioutil.ReadFile("day4.txt")
	s := strings.Split(string(input), "\n\n")
	nums := make([]int, 0)
	allBoards := make([]board, 0)
	for _, n := range strings.Split(s[0], ",") {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)
	}

	for _, rawBoard := range s[1:] {
		newBoard := board{}
		rows := strings.Split(rawBoard, "\n")
		for _, row := range rows {
			rowAsStr := strings.Fields(row)
			rowNums := make([]int, 0)
			for _, nStr := range rowAsStr {
				num, _ := strconv.Atoi(nStr)
				rowNums = append(rowNums, num)
			}
			newBoard.num = append(newBoard.num, rowNums)
			newBoard.marked = append(newBoard.marked, []bool{false, false, false, false, false})
		}
		allBoards = append(allBoards, newBoard)
	}

	p1, p2 := solve(allBoards, nums)

	println("part1", p1)
	println("part2", p2)
}
