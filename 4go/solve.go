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
				if b.checkBingo(row, col) { // if bingo, return true & score
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

// Could rewrite so both functions use the same method and just pick first / last out of order of bingo.
func solvep1(allBoards []board, nums []int) int {
	for _, num := range nums {
		for _, b := range allBoards {
			if done := b.markBoard(num); done {
				return b.calcScore(num)
			}
		}
	}
	panic("no bingo!! should not happen")
}

func solvep2(allBoards []board, nums []int) int {
	currLastNum := -1
	var currLastBoard board

	for _, b := range allBoards {
		if b, numsUsed := numsNeededToBingo(b, nums); numsUsed != -1 {
			if numsUsed > currLastNum {
				currLastBoard = b
				currLastNum = numsUsed
			}
		}
	}
	return currLastBoard.calcScore(nums[currLastNum])
}
func numsNeededToBingo(b board, nums []int) (board, int) {
	for totalNum, currNum := range nums {
		if done := b.markBoard(currNum); done {
			return b, totalNum
		}
	}
	return b, -1
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

	println("part1", solvep1(allBoards, nums))
	println("part2", solvep2(allBoards, nums))
}
