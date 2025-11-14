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

func calculateRiskSum(input [][]int) int {
	var riskSum int

	for i, line := range input {
		for j, value := range line {
			if (j > 0 && value >= line[j-1]) ||
				(j < len(line)-1 && value >= line[j+1]) ||
				(i > 0 && value >= input[i-1][j]) ||
				(i < len(input)-1 && value >= input[i+1][j]) {
				continue
			}

			riskSum += value + 1
		}
	}

	return riskSum
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

	fmt.Println(calculateRiskSum(input))
}
