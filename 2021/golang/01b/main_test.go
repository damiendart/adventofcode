// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"reflect"
	"testing"
)

func TestCountMeasurementWindowIncreases(t *testing.T) {
	input := []int{607, 618, 618, 617, 647, 716, 769, 792}

	if count := countMeasurementWindowIncreases(input); count != 5 {
		t.Error("Expected 5, got", count)
	}
}

func TestGetMeasurementWindows(t *testing.T) {
	expected := []int{607, 618, 618, 617, 647, 716, 769, 792}
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	if output := getMeasurementWindows(input); reflect.DeepEqual(output, expected) == false {
		t.Errorf("Expected %#v, got %#v", expected, output)
	}
}
