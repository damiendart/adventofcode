// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parseLine(input string) int32 {
	var firstIllegalCharacter int32
	var stack []int32

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
				return character
			}
		}
	}

	return firstIllegalCharacter
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	illegalCharacterCounts := map[int32]int{
		')': 0,
		']': 0,
		'}': 0,
		'>': 0,
	}

	for scanner.Scan() {
		firstIllegalCharacter := parseLine(scanner.Text())
		illegalCharacterCounts[firstIllegalCharacter]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(
		illegalCharacterCounts[')']*3 +
			illegalCharacterCounts[']']*57 +
			illegalCharacterCounts['}']*1197 +
			illegalCharacterCounts['>']*25137,
	)
}
