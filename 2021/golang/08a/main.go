// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countDigits(input []string) int {
	var count int

	for _, line := range input {
		for _, digit := range strings.Fields(line) {
			if len(digit) == 7 ||
				len(digit) == 4 ||
				len(digit) == 3 ||
				len(digit) == 2 {
				count++
			}
		}
	}

	return count
}

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")

		input = append(input, line[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(countDigits(input))
}
