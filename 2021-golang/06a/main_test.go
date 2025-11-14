// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"reflect"
	"testing"
)

func TestSimulateDays(t *testing.T) {
	expected := []int{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8}
	input := []int{3, 4, 3, 1, 2}

	if output := simulateDays(input, 18); reflect.DeepEqual(output, expected) == false {
		t.Errorf("Expected %#v, got %#v", expected, output)
	}

	if output := simulateDays(input, 80); len(output) == 5934 {
		t.Errorf("Expected 5934, got %#v", len(output))
	}
}
