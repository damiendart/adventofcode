package main

import "testing"

func TestCountMeasurementIncreases(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	if count := CountMeasurementIncreases(input); count != 7 {
		t.Error("Expected 7, got", count)
	}
}
