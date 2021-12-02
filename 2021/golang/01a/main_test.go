// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import "testing"

func TestCountMeasurementIncreases(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	if count := CountMeasurementIncreases(input); count != 7 {
		t.Error("Expected 7, got", count)
	}
}
