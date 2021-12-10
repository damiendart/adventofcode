// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import "testing"

func TestCalculateRiskSum(t *testing.T) {
	input := [][]int{
		[]int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		[]int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		[]int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		[]int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		[]int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	if output := calculateRiskSum(input); output != 15 {
		t.Errorf("Expected 15, got %#v", output)
	}
}
