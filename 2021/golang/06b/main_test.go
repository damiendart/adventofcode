// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import "testing"

func TestSimulateDays(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	if output := simulateDays(input, 18); output != 26 {
		t.Errorf("Expected 26, got %#v", output)
	}

	if output := simulateDays(input, 256); output != 26984457539 {
		t.Errorf("Expected 26984457539, got %#v", output)
	}
}
