// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import "testing"

func TestParseLine(t *testing.T) {
	input := "{([(<{}[<>[]}>{[]{[(<()>"

	if output := parseLine(input); output != '}' {
		t.Errorf("Expected '}', got '%c'", output)
	}

	input = "[[<[([]))<([[{}[[()]]]"

	if output := parseLine(input); output != ')' {
		t.Errorf("Expected ')', got '%c'", output)
	}

	input = "[{[{({}]{}}([{[{{{}}([]"

	if output := parseLine(input); output != ']' {
		t.Errorf("Expected ']', got '%c'", output)
	}

	input = "[<(<(<(<{}))><([]([]()"

	if output := parseLine(input); output != ')' {
		t.Errorf("Expected ')', got '%c'", output)
	}

	input = "<{([([[(<>()){}]>(<<{{"

	if output := parseLine(input); output != '>' {
		t.Errorf("Expected '>', got '%c'", output)
	}
}
