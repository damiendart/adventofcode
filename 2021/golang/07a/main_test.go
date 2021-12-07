// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import "testing"

func TestCalculateCostOfCheapestPosition(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	if output := calculateCostOfCheapestPosition(input); output != 37 {
		t.Errorf("Expected 37, got %#v", output)
	}
}
