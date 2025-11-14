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
	"unicode"
)

type node struct {
	neighbours []string
}

type nodes map[string]node

func containsElement(strings []string, needle string) bool {
	for _, str := range strings {
		if str == needle {
			return true
		}
	}

	return false
}

func isLower(str string) bool {
	for _, c := range str {
		if !unicode.IsLower(c) && unicode.IsLetter(c) {
			return false
		}
	}

	return true
}

func traverseCaves(journeys [][]string, caves nodes) [][]string {
	var newJourneys [][]string

	for _, journey := range journeys {
		if journey[len(journey)-1] == "end" {
			newJourneys = append(newJourneys, journey)
		} else {
			for _, neighbour := range caves[journey[len(journey)-1]].neighbours {
				if isLower(neighbour) && containsElement(journey, neighbour) {
					continue
				}

				newJourney := make([]string, len(journey))

				copy(newJourney, journey)

				newJourney = append(newJourney, neighbour)
				newJourneys = append(newJourneys, newJourney)

				newJourneys = traverseCaves(newJourneys, caves)
			}
		}
	}

	return newJourneys
}

func calculatePaths(input []string) [][]string {
	caves := make(nodes)

	for _, line := range input {
		cavesInput := strings.Split(line, "-")

		if cave, ok := caves[cavesInput[0]]; ok {
			cave.neighbours = append(cave.neighbours, cavesInput[1])
			caves[cavesInput[0]] = cave
		} else {
			newCave := node{neighbours: []string{cavesInput[1]}}
			caves[cavesInput[0]] = newCave
		}

		if cave, ok := caves[cavesInput[1]]; ok {
			cave.neighbours = append(cave.neighbours, cavesInput[0])
			caves[cavesInput[1]] = cave
		} else {
			newCave := node{neighbours: []string{cavesInput[0]}}
			caves[cavesInput[1]] = newCave
		}
	}

	return traverseCaves(
		[][]string{
			[]string{"start"},
		},
		caves,
	)
}

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(calculatePaths(input)))
}
