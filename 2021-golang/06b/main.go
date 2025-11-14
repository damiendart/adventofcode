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

func simulateDays(fishes []int, days int) int64 {
	var count int64
	var fishCounts [9]int64

	for _, fish := range fishes {
		fishCounts[fish]++
	}

	for day := 0; day < days; day++ {
		newFish := fishCounts[0]

		for i := 0; i < 8; i++ {
			fishCounts[i] = fishCounts[i+1]
		}

		fishCounts[6] += newFish
		fishCounts[8] = newFish
	}

	for _, fishCount := range fishCounts {
		count += fishCount
	}

	return count
}

func main() {
	var fishes []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		for _, value := range strings.Split(scanner.Text(), ",") {
			age, err := strconv.Atoi(value)

			if err != nil {
				log.Fatal(err)
			}

			fishes = append(fishes, age)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(simulateDays(fishes, 256))
}
