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
	"strings"
)

func countDigits(input [][]string) int {
	var total int

	for _, line := range input {
		var value int

		digits := strings.Fields(line[1])
		knownPatterns := make(map[int]string)

		for _, pattern := range strings.Fields(line[0]) {
			if len(pattern) == 2 {
				knownPatterns[1] = pattern
			} else if len(pattern) == 3 {
				knownPatterns[7] = pattern
			} else if len(pattern) == 4 {
				knownPatterns[4] = pattern
			} else if len(pattern) == 7 {
				knownPatterns[8] = pattern
			}
		}

		for i, digit := range digits {
			if len(digit) == 2 {
				value += int(math.Pow(10, float64(len(digits)-i-1)))
			} else if len(digit) == 3 {
				value += 7 * int(math.Pow(10, float64(len(digits)-i-1)))
			} else if len(digit) == 4 {
				value += 4 * int(math.Pow(10, float64(len(digits)-i-1)))
			} else if len(digit) == 7 {
				value += 8 * int(math.Pow(10, float64(len(digits)-i-1)))
			} else if len(digit) == 6 {
				regex := regexp.MustCompile("([" + knownPatterns[4] + "])")

				if len(regex.FindAllStringSubmatch(digit, -1)) == 4 {
					value += 9 * int(math.Pow(10, float64(len(digits)-i-1)))
				} else {
					regex = regexp.MustCompile("([" + knownPatterns[1] + "])")

					if len(regex.FindAllStringSubmatch(digit, -1)) != 2 {
						value += 6 * int(math.Pow(10, float64(len(digits)-i-1)))
					}
				}
			} else {
				regex := regexp.MustCompile("([" + knownPatterns[1] + "])")

				if len(regex.FindAllStringSubmatch(digit, -1)) == 2 {
					value += 3 * int(math.Pow(10, float64(len(digits)-i-1)))
				} else {
					regex = regexp.MustCompile("([" + knownPatterns[4] + "])")

					if len(regex.FindAllStringSubmatch(digit, -1)) == 2 {
						value += 2 * int(math.Pow(10, float64(len(digits)-i-1)))
					} else {
						value += 5 * int(math.Pow(10, float64(len(digits)-i-1)))
					}
				}

			}
		}

		total += value
	}

	return total
}

func main() {
	var input [][]string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")

		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(countDigits(input))
}
