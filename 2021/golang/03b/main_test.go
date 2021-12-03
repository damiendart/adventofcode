// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"strconv"
	"testing"
)

func TestCalculateMostCommonBitCount(t *testing.T) {
	if output := CalculateMostCommonBitCount(BitCount{2, 0}); output != 1 {
		t.Errorf("Expected 1, got %#v", output)
	}

	if output := CalculateMostCommonBitCount(BitCount{0, 2}); output != 0 {
		t.Errorf("Expected 0, got %#v", output)
	}

	if output := CalculateMostCommonBitCount(BitCount{2, 2}); output != 1 {
		t.Errorf("Expected 1, got %#v", output)
	}
}

func TestCalculateLeastCommonBitCount(t *testing.T) {
	if output := CalculateLeastCommonBitCount(BitCount{2, 0}); output != 0 {
		t.Errorf("Expected 0, got %#v", output)
	}

	if output := CalculateLeastCommonBitCount(BitCount{0, 2}); output != 1 {
		t.Errorf("Expected 1, got %#v", output)
	}

	if output := CalculateLeastCommonBitCount(BitCount{2, 2}); output != 0 {
		t.Errorf("Expected 0, got %#v", output)
	}
}

func TestCalculateRating(t *testing.T) {
	var inputConverted []int64

	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	for _, line := range input {
		value, err := strconv.ParseInt(line, 2, 16)

		if err != nil {
			t.Error(err)
		}

		inputConverted = append(inputConverted, value)
	}

	if output := CalculateRate(inputConverted, 5, CalculateLeastCommonBitCount); output != 10 {
		t.Errorf("Expected 10, got %#v", output)
	}

	if output := CalculateRate(inputConverted, 5, CalculateMostCommonBitCount); output != 23 {
		t.Errorf("Expected 23, got %#v", output)
	}
}
