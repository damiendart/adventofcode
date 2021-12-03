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

type BitCount struct {
	one  int
	zero int
}

func CalculateMostCommonBitCount(bitCount BitCount) int64 {
	if bitCount.one >= bitCount.zero {
		return 1
	}

	return 0
}

func CalculateLeastCommonBitCount(bitCount BitCount) int64 {
	if bitCount.one >= bitCount.zero {
		return 0
	}

	return 1
}

func CalculateRate(report []int64, numberOfBits int, commonBitFunction func(BitCount) int64) int {
	filteredReport := make([]int64, len(report))

	copy(filteredReport, report)

	for i := numberOfBits - 1; i >= 0; i-- {
		if len(filteredReport) == 1 {
			break
		}
		bitCounts := make([]BitCount, numberOfBits)

		for _, line := range filteredReport {
			for i := len(bitCounts) - 1; i >= 0; i-- {
				if (line>>i)&1 == 1 {
					bitCounts[i].one++
				} else {
					bitCounts[i].zero++
				}
			}
		}

		var j int

		for _, line := range filteredReport {
			if (line>>i)&1 == commonBitFunction(bitCounts[i]) {
				filteredReport[j] = line
				j++
			}
		}

		filteredReport = filteredReport[:j]
	}

	return int(filteredReport[0])
}

func main() {
	var report []int64
	var reportLineLength int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		reportLineLength = len(line)
		value, err := strconv.ParseInt(line, 2, 16)

		if err != nil {
			log.Fatal(err)
		}

		report = append(report, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	generatorRating := CalculateRate(report, reportLineLength, CalculateMostCommonBitCount)
	scrubberRating := CalculateRate(report, reportLineLength, CalculateLeastCommonBitCount)

	fmt.Println(generatorRating * scrubberRating)
}
