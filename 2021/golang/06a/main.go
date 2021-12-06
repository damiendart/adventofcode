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

func simulateDays(fishes []int, days int) []int {

	for day := 0; day < days; day++ {
		var additionalFishes []int

		for i := range fishes {
			if fishes[i] == 0 {
				additionalFishes = append(additionalFishes, 8)
				fishes[i] = 6
			} else {
				fishes[i]--
			}
		}

		fishes = append(fishes, additionalFishes...)
	}

	return fishes
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

	fmt.Println(len(simulateDays(fishes, 256)))
}
