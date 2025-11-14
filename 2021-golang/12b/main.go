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

func containsElement(strings []string, needle string) int {
	var count int

	for _, str := range strings {
		if str == needle {
			count++
		}
	}

	return count
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
				if isLower(neighbour) {
					if neighbour == "start" {
						continue
					}

					if journey[0] == neighbour && containsElement(journey[1:], neighbour) > 1 {
						continue
					}

					if journey[0] != neighbour && containsElement(journey[1:], neighbour) > 0 {
						continue
					}

					if journey[0] == "" {
						alternativeJourney := make([]string, len(journey))

						copy(alternativeJourney, journey)

						if neighbour != "end" {
							alternativeJourney[0] = neighbour
						}

						alternativeJourney = append(alternativeJourney, neighbour)
						newJourneys = append(newJourneys, alternativeJourney)
					}
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

	paths := traverseCaves(
		[][]string{
			[]string{"", "start"},
		},
		caves,
	)

	var uniquePaths [][]string

	uniquePaths = append(uniquePaths, paths[0])

	for _, path := range paths[1:] {
		unique := true

		for _, uniquePath := range uniquePaths {
			equal := true

			if len(path) != len(uniquePath) {
				equal = false
			} else {
				for i := 1; i < len(path); i++ {
					if path[i] != uniquePath[i] {
						equal = false
						break
					}
				}
			}

			if equal {
				unique = false
				break
			}
		}

		if unique {
			uniquePaths = append(uniquePaths, path)
		}
	}

	return uniquePaths
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
