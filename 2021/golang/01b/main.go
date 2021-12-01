package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var count int
	var measurements []int
	var measurementWindows []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		measurements = append(measurements, measurement)
	}

	for i := 2; i < len(measurements); i++ {
		measurementWindows = append(
			measurementWindows,
			measurements[i-2]+measurements[i-1]+measurements[i],
		)
	}

	for i := 1; i < len(measurementWindows); i++ {
		if measurementWindows[i] > measurementWindows[i-1] {
			count++
		}
	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
