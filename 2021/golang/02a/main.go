// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Command struct {
	regex   *regexp.Regexp
	handler func(Position, int) Position
}

type Position struct {
	x     int
	depth int
}

var Commands = []Command{
	Command{
		regexp.MustCompile("^down ([0-9]+)$"),
		func(position Position, value int) Position {
			position.depth += value

			return position
		},
	},
	Command{
		regexp.MustCompile("^forward ([0-9]+)$"),
		func(position Position, value int) Position {
			position.x += value

			return position
		},
	},
	Command{
		regexp.MustCompile("^up ([0-9]+)$"),
		func(position Position, value int) Position {
			position.depth -= value

			return position
		},
	},
}

func ParseInstructions(instructions []string) Position {
	var position Position

	for _, instruction := range instructions {
		for _, command := range Commands {
			matches := command.regex.FindStringSubmatch(instruction)

			if len(matches) > 0 {
				value, err := strconv.Atoi(matches[1])

				if err != nil {
					log.Fatal(err)
				}

				position = command.handler(position, value)
			}
		}
	}

	return position
}

func main() {
	var instructions []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	position := ParseInstructions(instructions)

	fmt.Println(position.depth * position.x)
}
