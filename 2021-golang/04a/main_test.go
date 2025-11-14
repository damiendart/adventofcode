// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"reflect"
	"testing"
)

func TestCalculateScore(t *testing.T) {
	inputBoards := []board{
		board{
			cells: [25]cell{
				cell{value: 22, marked: false},
				cell{value: 13, marked: false},
				cell{value: 17, marked: false},
				cell{value: 11, marked: false},
				cell{value: 0, marked: false},
				cell{value: 8, marked: false},
				cell{value: 2, marked: false},
				cell{value: 23, marked: false},
				cell{value: 4, marked: false},
				cell{value: 24, marked: false},
				cell{value: 21, marked: false},
				cell{value: 9, marked: false},
				cell{value: 14, marked: false},
				cell{value: 16, marked: false},
				cell{value: 7, marked: false},
				cell{value: 6, marked: false},
				cell{value: 10, marked: false},
				cell{value: 3, marked: false},
				cell{value: 18, marked: false},
				cell{value: 5, marked: false},
				cell{value: 1, marked: false},
				cell{value: 12, marked: false},
				cell{value: 20, marked: false},
				cell{value: 15, marked: false},
				cell{value: 19, marked: false},
			},
		},
		board{
			cells: [25]cell{
				cell{value: 3, marked: false},
				cell{value: 15, marked: false},
				cell{value: 0, marked: false},
				cell{value: 2, marked: false},
				cell{value: 22, marked: false},
				cell{value: 9, marked: false},
				cell{value: 18, marked: false},
				cell{value: 13, marked: false},
				cell{value: 17, marked: false},
				cell{value: 5, marked: false},
				cell{value: 19, marked: false},
				cell{value: 8, marked: false},
				cell{value: 7, marked: false},
				cell{value: 25, marked: false},
				cell{value: 23, marked: false},
				cell{value: 20, marked: false},
				cell{value: 11, marked: false},
				cell{value: 10, marked: false},
				cell{value: 24, marked: false},
				cell{value: 4, marked: false},
				cell{value: 14, marked: false},
				cell{value: 21, marked: false},
				cell{value: 16, marked: false},
				cell{value: 12, marked: false},
				cell{value: 6, marked: false},
			},
		},
		board{
			cells: [25]cell{
				cell{value: 14, marked: false},
				cell{value: 21, marked: false},
				cell{value: 17, marked: false},
				cell{value: 24, marked: false},
				cell{value: 4, marked: false},
				cell{value: 10, marked: false},
				cell{value: 16, marked: false},
				cell{value: 15, marked: false},
				cell{value: 9, marked: false},
				cell{value: 19, marked: false},
				cell{value: 18, marked: false},
				cell{value: 8, marked: false},
				cell{value: 23, marked: false},
				cell{value: 26, marked: false},
				cell{value: 20, marked: false},
				cell{value: 22, marked: false},
				cell{value: 11, marked: false},
				cell{value: 13, marked: false},
				cell{value: 6, marked: false},
				cell{value: 5, marked: false},
				cell{value: 2, marked: false},
				cell{value: 0, marked: false},
				cell{value: 12, marked: false},
				cell{value: 3, marked: false},
				cell{value: 7, marked: false},
			},
		},
	}

	inputNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	if output := calculateScore(inputBoards, inputNumbers); output != 4512 {
		t.Errorf("Expected 4512, got %d", output)
	}
}

func TestCheckBoardHorizontal(t *testing.T) {
	input := board{
		cells: [25]cell{
			cell{value: 22, marked: false},
			cell{value: 13, marked: false},
			cell{value: 17, marked: false},
			cell{value: 11, marked: false},
			cell{value: 0, marked: false},
			cell{value: 8, marked: false},
			cell{value: 2, marked: false},
			cell{value: 23, marked: false},
			cell{value: 4, marked: false},
			cell{value: 24, marked: false},
			cell{value: 21, marked: false},
			cell{value: 9, marked: false},
			cell{value: 14, marked: false},
			cell{value: 16, marked: false},
			cell{value: 7, marked: false},
			cell{value: 6, marked: false},
			cell{value: 10, marked: false},
			cell{value: 3, marked: false},
			cell{value: 18, marked: false},
			cell{value: 5, marked: false},
			cell{value: 1, marked: false},
			cell{value: 12, marked: false},
			cell{value: 20, marked: false},
			cell{value: 15, marked: false},
			cell{value: 19, marked: false},
		},
	}

	if checkBoard(input) != false {
		t.Error("Expected false, got true")
	}

	input.cells[5].marked = true
	input.cells[6].marked = true
	input.cells[7].marked = true
	input.cells[8].marked = true
	input.cells[9].marked = true

	if checkBoard(input) != true {
		t.Error("Expected true, got false")
	}
}

func TestCheckBoardVertical(t *testing.T) {
	input := board{
		cells: [25]cell{
			cell{value: 22, marked: false},
			cell{value: 13, marked: false},
			cell{value: 17, marked: false},
			cell{value: 11, marked: false},
			cell{value: 0, marked: false},
			cell{value: 8, marked: false},
			cell{value: 2, marked: false},
			cell{value: 23, marked: false},
			cell{value: 4, marked: false},
			cell{value: 24, marked: false},
			cell{value: 21, marked: false},
			cell{value: 9, marked: false},
			cell{value: 14, marked: false},
			cell{value: 16, marked: false},
			cell{value: 7, marked: false},
			cell{value: 6, marked: false},
			cell{value: 10, marked: false},
			cell{value: 3, marked: false},
			cell{value: 18, marked: false},
			cell{value: 5, marked: false},
			cell{value: 1, marked: false},
			cell{value: 12, marked: false},
			cell{value: 20, marked: false},
			cell{value: 15, marked: false},
			cell{value: 19, marked: false},
		},
	}

	if checkBoard(input) != false {
		t.Error("Expected false, got true")
	}

	input.cells[2].marked = true
	input.cells[7].marked = true
	input.cells[12].marked = true
	input.cells[17].marked = true
	input.cells[22].marked = true

	if checkBoard(input) != true {
		t.Error("Expected true, got false")
	}
}

func TestMarkBoard(t *testing.T) {
	input := board{
		cells: [25]cell{
			cell{value: 22, marked: false},
			cell{value: 13, marked: false},
			cell{value: 17, marked: false},
			cell{value: 11, marked: false},
			cell{value: 0, marked: false},
			cell{value: 8, marked: false},
			cell{value: 2, marked: false},
			cell{value: 23, marked: false},
			cell{value: 4, marked: false},
			cell{value: 24, marked: false},
			cell{value: 21, marked: false},
			cell{value: 9, marked: false},
			cell{value: 14, marked: false},
			cell{value: 16, marked: false},
			cell{value: 7, marked: false},
			cell{value: 6, marked: false},
			cell{value: 10, marked: false},
			cell{value: 3, marked: false},
			cell{value: 18, marked: false},
			cell{value: 5, marked: false},
			cell{value: 1, marked: false},
			cell{value: 12, marked: false},
			cell{value: 20, marked: false},
			cell{value: 15, marked: false},
			cell{value: 19, marked: false},
		},
	}

	expected := input

	expected.cells[1].marked = true
	expected.cells[16].marked = true

	input = markBoard(input, 13)
	input = markBoard(input, 10)

	if reflect.DeepEqual(input, expected) == false {
		t.Errorf("Expected %#v, got %#v", expected, input)
	}
}

func TestParseInput(t *testing.T) {
	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

	expectedBoards := []board{
		board{
			cells: [25]cell{
				cell{value: 22, marked: false},
				cell{value: 13, marked: false},
				cell{value: 17, marked: false},
				cell{value: 11, marked: false},
				cell{value: 0, marked: false},
				cell{value: 8, marked: false},
				cell{value: 2, marked: false},
				cell{value: 23, marked: false},
				cell{value: 4, marked: false},
				cell{value: 24, marked: false},
				cell{value: 21, marked: false},
				cell{value: 9, marked: false},
				cell{value: 14, marked: false},
				cell{value: 16, marked: false},
				cell{value: 7, marked: false},
				cell{value: 6, marked: false},
				cell{value: 10, marked: false},
				cell{value: 3, marked: false},
				cell{value: 18, marked: false},
				cell{value: 5, marked: false},
				cell{value: 1, marked: false},
				cell{value: 12, marked: false},
				cell{value: 20, marked: false},
				cell{value: 15, marked: false},
				cell{value: 19, marked: false},
			},
		},
		board{
			cells: [25]cell{
				cell{value: 3, marked: false},
				cell{value: 15, marked: false},
				cell{value: 0, marked: false},
				cell{value: 2, marked: false},
				cell{value: 22, marked: false},
				cell{value: 9, marked: false},
				cell{value: 18, marked: false},
				cell{value: 13, marked: false},
				cell{value: 17, marked: false},
				cell{value: 5, marked: false},
				cell{value: 19, marked: false},
				cell{value: 8, marked: false},
				cell{value: 7, marked: false},
				cell{value: 25, marked: false},
				cell{value: 23, marked: false},
				cell{value: 20, marked: false},
				cell{value: 11, marked: false},
				cell{value: 10, marked: false},
				cell{value: 24, marked: false},
				cell{value: 4, marked: false},
				cell{value: 14, marked: false},
				cell{value: 21, marked: false},
				cell{value: 16, marked: false},
				cell{value: 12, marked: false},
				cell{value: 6, marked: false},
			},
		},
		board{
			cells: [25]cell{
				cell{value: 14, marked: false},
				cell{value: 21, marked: false},
				cell{value: 17, marked: false},
				cell{value: 24, marked: false},
				cell{value: 4, marked: false},
				cell{value: 10, marked: false},
				cell{value: 16, marked: false},
				cell{value: 15, marked: false},
				cell{value: 9, marked: false},
				cell{value: 19, marked: false},
				cell{value: 18, marked: false},
				cell{value: 8, marked: false},
				cell{value: 23, marked: false},
				cell{value: 26, marked: false},
				cell{value: 20, marked: false},
				cell{value: 22, marked: false},
				cell{value: 11, marked: false},
				cell{value: 13, marked: false},
				cell{value: 6, marked: false},
				cell{value: 5, marked: false},
				cell{value: 2, marked: false},
				cell{value: 0, marked: false},
				cell{value: 12, marked: false},
				cell{value: 3, marked: false},
				cell{value: 7, marked: false},
			},
		},
	}

	expectedDrawnNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	boards, drawnNumbers := parseInput(input)

	if reflect.DeepEqual(drawnNumbers, expectedDrawnNumbers) == false {
		t.Errorf("Expected %#v, got %#v", expectedDrawnNumbers, drawnNumbers)
	}

	if reflect.DeepEqual(boards, expectedBoards) == false {
		t.Errorf("Expected %#v, got %#v", expectedBoards, boards)
	}
}
