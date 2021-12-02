// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"reflect"
	"testing"
)

func TestParseInstructions(t *testing.T) {
	expected := Position{15, 10}
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	if output := ParseInstructions(input); reflect.DeepEqual(output, expected) == false {
		t.Errorf("Expected %#v, got %#v", expected, output)
	}
}
