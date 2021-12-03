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

type command struct {
	regex   *regexp.Regexp
	handler func(position, int) position
}

type position struct {
	x     int
	depth int
}

var commands = []command{
	command{
		regexp.MustCompile("^down ([0-9]+)$"),
		func(position position, value int) position {
			position.depth += value

			return position
		},
	},
	command{
		regexp.MustCompile("^forward ([0-9]+)$"),
		func(position position, value int) position {
			position.x += value

			return position
		},
	},
	command{
		regexp.MustCompile("^up ([0-9]+)$"),
		func(position position, value int) position {
			position.depth -= value

			return position
		},
	},
}

func parseInstructions(instructions []string) position {
	var position position

	for _, instruction := range instructions {
		for _, command := range commands {
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

	position := parseInstructions(instructions)

	fmt.Println(position.depth * position.x)
}
