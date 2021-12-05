// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"reflect"
	"testing"
)

func TestCalculateOverlaps(t *testing.T) {
	input := []line{
		line{0, 9, 5, 9},
		line{9, 4, 3, 4},
		line{2, 2, 2, 1},
		line{7, 0, 7, 4},
		line{0, 9, 2, 9},
		line{3, 4, 1, 4},
	}

	if output := calculateOverlaps(input); output != 5 {
		t.Errorf("Expected 5, got %#v", output)
	}
}

func TestExpandLine(t *testing.T) {
	input := line{0, 9, 5, 9}
	expected := []cell{
		cell{0, 9},
		cell{1, 9},
		cell{2, 9},
		cell{3, 9},
		cell{4, 9},
		cell{5, 9},
	}

	if output := expandLine(input); reflect.DeepEqual(output, expected) == false {
		t.Errorf("Expected %#v, got %#v", expected, output)
	}

	input = line{5, 9, 0, 9}
	expected = []cell{
		cell{5, 9},
		cell{4, 9},
		cell{3, 9},
		cell{2, 9},
		cell{1, 9},
		cell{0, 9},
	}

	if output := expandLine(input); reflect.DeepEqual(output, expected) == false {
		t.Errorf("Expected %#v, got %#v", expected, output)
	}

	input = line{9, 0, 9, 5}
	expected = []cell{
		cell{9, 0},
		cell{9, 1},
		cell{9, 2},
		cell{9, 3},
		cell{9, 4},
		cell{9, 5},
	}

	if output := expandLine(input); reflect.DeepEqual(output, expected) == false {
		t.Errorf("Expected %#v, got %#v", expected, output)
	}

	input = line{9, 5, 9, 0}
	expected = []cell{
		cell{9, 5},
		cell{9, 4},
		cell{9, 3},
		cell{9, 2},
		cell{9, 1},
		cell{9, 0},
	}

	if output := expandLine(input); reflect.DeepEqual(output, expected) == false {
		t.Errorf("Expected %#v, got %#v", expected, output)
	}
}

func TestIsValidLine(t *testing.T) {
	if output := isValidLine(line{9, 4, 3, 4}); output != true {
		t.Errorf("Expected true, got %#v", output)
	}

	if output := isValidLine(line{7, 0, 7, 4}); output != true {
		t.Errorf("Expected true, got %#v", output)
	}

	if output := isValidLine(line{6, 2, 8, 5}); output != false {
		t.Errorf("Expected false, got %#v", output)
	}
}

func TestParseLine(t *testing.T) {
	input := "0,9 -> 5,9"
	expected := line{0, 9, 5, 9}

	if output := parseLine(input); reflect.DeepEqual(output, expected) == false {
		t.Errorf("Expected %#v, got %#v", expected, output)
	}

	input = "1000,546 -> 432,2312"
	expected = line{1000, 546, 432, 2312}

	if output := parseLine(input); reflect.DeepEqual(output, expected) == false {
		t.Errorf("Expected %#v, got %#v", expected, output)
	}
}
