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

func CalculateMostCommonBitCount(bitCount BitCount) int {
	if bitCount.one >= bitCount.zero {
		return 1
	}

	return 0
}

func CalculateLeastCommonBitCount(bitCount BitCount) int {
	if bitCount.one >= bitCount.zero {
		return 0
	}

	return 1
}

func CalculateRate(report []int64, numberOfBits int, commonBitFunction func(BitCount) int) int {
	var bitCounts = make([]BitCount, numberOfBits)
	var rate int

	for _, line := range report {
		for i := 0; i < numberOfBits; i++ {
			if (line>>i)&1 == 1 {
				bitCounts[i].one++
			} else {
				bitCounts[i].zero++
			}
		}
	}

	for i, count := range bitCounts {
		rate = (rate &^ (1 << i)) | (commonBitFunction(count) << i)
	}

	return rate
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

	epsilonRate := CalculateRate(report, reportLineLength, CalculateLeastCommonBitCount)
	gammaRate := CalculateRate(report, reportLineLength, CalculateMostCommonBitCount)

	fmt.Println(epsilonRate * gammaRate)
}
