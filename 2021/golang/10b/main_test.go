// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import "testing"

func TestParseLine(t *testing.T) {
	input := "[({(<(())[]>[[{[]{<()<>>"

	if output := parseLine(input); output != 288957 {
		t.Errorf("Expected 288957, got %#v", output)
	}

	input = "[(()[<>])]({[<{<<[]>>("

	if output := parseLine(input); output != 5566 {
		t.Errorf("Expected 5566, got %#v", output)
	}

	input = "(((({<>}<{<{<>}{[]{[]{}"

	if output := parseLine(input); output != 1480781 {
		t.Errorf("Expected 1480781, got %#v", output)
	}

	input = "{<[[]]>}<{[{[{[]{()[[[]"

	if output := parseLine(input); output != 995444 {
		t.Errorf("Expected 995444, got %#v", output)
	}

	input = "<{([{{}}[<[[[<>{}]]]>[]]"

	if output := parseLine(input); output != 294 {
		t.Errorf("Expected 294, got %#v", output)
	}
}
