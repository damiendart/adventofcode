// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import "testing"

func TestCountDigits(t *testing.T) {
	input := []string{
		"fdgacbe cefdb cefbgd gcbe",
		"fcgedb cgb dgebacf gc",
		"cg cg fdcagb cbg",
		"efabcd cedba gadfec cb",
		"gecf egdcabf bgf bfgea",
		"gebdcfa ecba ca fadegcb",
		"cefg dcbef fcge gbcadfe",
		"ed bcgafe cdgba cbgef",
		"gbdfcae bgc cg cgb",
		"fgae cfgab fg bagce",
	}

	if output := countDigits(input); output != 26 {
		t.Errorf("Expected 26, got %#v", output)
	}
}
