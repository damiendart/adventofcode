// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import "testing"

func TestCalculateFlashes(t *testing.T) {
	input := [][]int{
		[]int{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		[]int{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		[]int{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		[]int{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		[]int{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		[]int{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		[]int{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		[]int{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		[]int{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		[]int{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}

	if output := calculateFlashes(input, 10); output != 204 {
		t.Errorf("Expected 204, got %#v", output)
	}

	input = [][]int{
		[]int{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		[]int{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		[]int{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		[]int{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		[]int{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		[]int{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		[]int{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		[]int{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		[]int{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		[]int{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}

	if output := calculateFlashes(input, 100); output != 1656 {
		t.Errorf("Expected 1656, got %#v", output)
	}

}
