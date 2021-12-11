// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateSynchronisedStep(input [][]int) int {
	steps := -1

	for {
		synchronised := true

		steps++

		for i, line := range input {
			for j := range line {
				if input[i][j] > 0 {
					synchronised = false
				}
			}
		}

		if synchronised {
			break
		}

		for i, line := range input {
			for j := range line {
				input[i][j]++
			}
		}

		flashing := true

		for flashing {
			for i, line := range input {
				for j := range line {
					if input[i][j] > 9 {
						if j > 0 && input[i][j-1] != 0 {
							input[i][j-1]++
						}

						if j < len(line)-1 && input[i][j+1] != 0 {
							input[i][j+1]++
						}

						if i > 0 && input[i-1][j] != 0 {
							input[i-1][j]++
						}

						if i < len(input)-1 && input[i+1][j] != 0 {
							input[i+1][j]++
						}

						if i > 0 && j > 0 && input[i-1][j-1] != 0 {
							input[i-1][j-1]++
						}

						if i > 0 && j < len(line)-1 && input[i-1][j+1] != 0 {
							input[i-1][j+1]++
						}

						if i < len(input)-1 && j > 0 && input[i+1][j-1] != 0 {
							input[i+1][j-1]++
						}

						if i < len(input)-1 && j < len(line)-1 && input[i+1][j+1] != 0 {
							input[i+1][j+1]++
						}

						input[i][j] = 0
					}
				}
			}

			flashing = false

			for i, line := range input {
				for j := range line {
					if input[i][j] > 9 {
						flashing = true
					}
				}
			}
		}
	}

	return steps
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

	fmt.Println(calculateSynchronisedStep(input))
}
