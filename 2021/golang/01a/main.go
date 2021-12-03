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
)

func countMeasurementIncreases(measurements []int) int {
	var count int

	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			count++
		}
	}

	return count
}

func main() {
	var measurements []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		measurements = append(measurements, measurement)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(countMeasurementIncreases(measurements))
}
