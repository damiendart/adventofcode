// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type cell struct {
	x int
	y int
}

type line struct {
	startX int
	startY int
	endX   int
	endY   int
}

func calculateOverlaps(lines []line) int {
	var overlaps []cell

	for i := range lines {
		for j := range lines[:i] {
			for _, cell := range expandLine(lines[i]) {
				for _, cellToCompare := range expandLine(lines[j]) {
					if cell.x == cellToCompare.x && cell.y == cellToCompare.y {
						if len(overlaps) == 0 {
							overlaps = append(overlaps, cell)
						} else {
							unique := true

							for _, overlapCell := range overlaps {
								if cell.x == overlapCell.x && cell.y == overlapCell.y {
									unique = false
									break
								}
							}

							if unique {
								overlaps = append(overlaps, cell)
							}
						}
					}
				}
			}
		}
	}

	return len(overlaps)
}

func expandLine(line line) []cell {
	var cells []cell

	if line.startX == line.endX {
		if line.endY > line.startY {
			for i := line.startY; i <= line.endY; i++ {
				var cell cell

				cell.x = line.startX
				cell.y = i
				cells = append(cells, cell)
			}
		} else {
			for i := line.startY; i >= line.endY; i-- {
				var cell cell

				cell.x = line.startX
				cell.y = i
				cells = append(cells, cell)
			}
		}
	} else if line.startY == line.endY {
		if line.endX > line.startX {
			for i := line.startX; i <= line.endX; i++ {
				var cell cell

				cell.x = i
				cell.y = line.startY
				cells = append(cells, cell)
			}
		} else {
			for i := line.startX; i >= line.endX; i-- {
				var cell cell

				cell.x = i
				cell.y = line.startY
				cells = append(cells, cell)
			}
		}
	} else {
		dX := line.endX - line.startX
		dY := line.endY - line.startY

		for i := 0; i <= int(math.Abs(float64(dX))); i++ {
			var cell cell

			cell.x = line.startX + (dX / int(math.Abs(float64(dX))) * i)
			cell.y = line.startY + (dY / int(math.Abs(float64(dY))) * i)
			cells = append(cells, cell)
		}
	}

	return cells
}

func parseLine(input string) line {
	regex := regexp.MustCompile("([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)$")
	matches := regex.FindStringSubmatch(input)

	var line line

	startX, err := strconv.Atoi(matches[1])

	if err != nil {
		log.Fatal(err)
	}

	startY, err := strconv.Atoi(matches[2])

	if err != nil {
		log.Fatal(err)
	}

	endX, err := strconv.Atoi(matches[3])

	if err != nil {
		log.Fatal(err)
	}

	endY, err := strconv.Atoi(matches[4])

	if err != nil {
		log.Fatal(err)
	}

	line.startX = startX
	line.startY = startY
	line.endX = endX
	line.endY = endY

	return line
}

func main() {
	var lines []line

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		lines = append(lines, parseLine(input))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(calculateOverlaps(lines))
}
