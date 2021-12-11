// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func parseLine(input string) int {
	var stack []int32
	var score int

	for _, character := range input {
		switch character {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case '<':
			stack = append(stack, '>')
		default:
			if character == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return 0
			}
		}
	}

	for i := len(stack)/2 - 1; i >= 0; i-- {
		opposite := len(stack) - 1 - i
		stack[i], stack[opposite] = stack[opposite], stack[i]
	}

	for _, character := range stack {
		score *= 5

		switch character {
		case ')':
			score++
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		}
	}

	return score
}

func main() {
	var scores []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		score := parseLine(scanner.Text())

		if score > 0 {
			scores = append(scores, score)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(scores)

	fmt.Println(scores[len(scores)/2])
}
