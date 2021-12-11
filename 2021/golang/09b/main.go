// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type cell struct {
	x int
	y int
}

func checkFlow(area [][]int, position cell) []cell {
	var validCells []cell

	validCells = append(validCells, position)

	if position.x > 0 &&
		area[position.y][position.x-1] > area[position.y][position.x] &&
		area[position.y][position.x-1] != 9 {
		newCell := cell{position.x - 1, position.y}

		validCells = append(validCells, checkFlow(area, newCell)...)
	}

	if position.x < len(area[0])-1 &&
		area[position.y][position.x+1] > area[position.y][position.x] &&
		area[position.y][position.x+1] != 9 {
		newCell := cell{position.x + 1, position.y}

		validCells = append(validCells, checkFlow(area, newCell)...)
	}

	if position.y > 0 &&
		area[position.y-1][position.x] > area[position.y][position.x] &&
		area[position.y-1][position.x] != 9 {
		newCell := cell{position.x, position.y - 1}

		validCells = append(validCells, checkFlow(area, newCell)...)
	}

	if position.y < len(area)-1 &&
		area[position.y+1][position.x] > area[position.y][position.x] &&
		area[position.y+1][position.x] != 9 {
		newCell := cell{position.x, position.y + 1}

		validCells = append(validCells, checkFlow(area, newCell)...)
	}

	return validCells
}

func calculateBasins(input [][]int) []int {
	var basins []int

	for i, line := range input {
		for j, value := range line {
			if (j > 0 && value >= line[j-1]) ||
				(j < len(line)-1 && value >= line[j+1]) ||
				(i > 0 && value >= input[i-1][j]) ||
				(i < len(input)-1 && value >= input[i+1][j]) {
				continue
			}

			var uniqueBasinCells []cell

			cell := cell{j, i}
			basinCells := checkFlow(input, cell)

			for _, cell := range basinCells {
				isUnique := true

				for _, uniqueCell := range uniqueBasinCells {
					if (uniqueCell.x == cell.x) && (uniqueCell.y == cell.y) {
						isUnique = false
						break
					}
				}

				if isUnique {
					uniqueBasinCells = append(uniqueBasinCells, cell)
				}
			}

			basins = append(basins, len(uniqueBasinCells))
		}
	}

	sort.Ints(basins)
	return basins
}

func main() {
	var input [][]int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var lineValues []int

		line := scanner.Text()

		for _, item := range strings.Split(line, "") {
			number, err := strconv.Atoi(item)

			if err != nil {
				log.Fatal(err)
			}

			lineValues = append(lineValues, number)
		}

		input = append(input, lineValues)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var value int

	basins := calculateBasins(input)

	for _, basin := range basins[len(basins)-3:] {
		if value == 0 {
			value += basin
		} else {
			value *= basin
		}
	}

	fmt.Println(value)
}
