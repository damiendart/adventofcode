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
	"strconv"
	"strings"
)

func calculateCostOfCheapestPosition(input []int) int {
	var cheapestCost int
	var furthestCrabPosition int

	for _, position := range input {
		if position > furthestCrabPosition {
			furthestCrabPosition = position
		}
	}

	for i := 0; i <= furthestCrabPosition; i++ {
		var fuelCost int

		for _, position := range input {
			if i == position {
				continue
			}

			distance := int(math.Abs(float64(i - position)))

			fuelCost += (int(math.Pow(float64(distance), 2)) + distance) / 2
		}

		if i == 0 || fuelCost < cheapestCost {
			cheapestCost = fuelCost
		}
	}

	return cheapestCost
}

func main() {
	var positions []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		for _, value := range strings.Split(scanner.Text(), ",") {
			position, err := strconv.Atoi(value)

			if err != nil {
				log.Fatal(err)
			}

			positions = append(positions, position)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(calculateCostOfCheapestPosition(positions))
}
