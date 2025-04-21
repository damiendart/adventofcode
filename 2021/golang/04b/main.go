// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type board struct {
	cells    [25]cell
	complete bool
}

type cell struct {
	value  int
	marked bool
}

func calculateScore(boards []board, drawnNumbers []int) int {
	var lastWinningBoard board
	var lastWinningNumber int
	var sum int

	filteredBoards := boards

	for _, number := range drawnNumbers {
		for i := range filteredBoards {
			boards[i] = markBoard(boards[i], number)

			if !boards[i].complete && checkBoard(boards[i]) {
				boards[i].complete = true
				lastWinningBoard = boards[i]
				lastWinningNumber = number
			}
		}
	}

	for _, cell := range lastWinningBoard.cells {
		if !cell.marked {
			sum += cell.value
		}
	}

	return sum * lastWinningNumber
}

func checkBoard(board board) bool {
	for i := 0; i < 5; i++ {
		if board.cells[5*i].marked &&
			board.cells[5*i+1].marked &&
			board.cells[5*i+2].marked &&
			board.cells[5*i+3].marked &&
			board.cells[5*i+4].marked {
			return true
		}

		if board.cells[i].marked &&
			board.cells[i+5].marked &&
			board.cells[i+10].marked &&
			board.cells[i+15].marked &&
			board.cells[i+20].marked {
			return true
		}
	}

	return false
}

func markBoard(board board, number int) board {
	for i := range board.cells {
		if board.cells[i].value == number {
			board.cells[i].marked = true
		}
	}

	return board
}

func parseInput(input string) ([]board, []int) {
	var boards []board
	var drawnNumbers []int

	drawnNumbersInput := strings.SplitN(input, "\n", 2)

	for _, item := range strings.Split(drawnNumbersInput[0], ",") {
		number, err := strconv.Atoi(item)

		if err != nil {
			log.Fatal(err)
		}

		drawnNumbers = append(drawnNumbers, number)
	}

	if len(drawnNumbersInput) > 1 {
		for _, boardInput := range strings.Split(drawnNumbersInput[1], "\n\n") {
			var board board

			for j, cellInput := range strings.Fields(boardInput) {
				number, err := strconv.Atoi(cellInput)

				if err != nil {
					log.Fatal(err)
				}

				cell := new(cell)

				cell.value = number

				board.cells[j] = *cell
			}

			boards = append(boards, board)
		}
	}

	return boards, drawnNumbers
}

func main() {
	input, err := io.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal(err)
	}

	boards, drawnNumbers := parseInput(string(input))

	fmt.Println(calculateScore(boards, drawnNumbers))
}
