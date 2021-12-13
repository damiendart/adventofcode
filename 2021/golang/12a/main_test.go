// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import (
	"reflect"
	"testing"
)

func TestCalculatePaths(t *testing.T) {
	expected := [][]string{
		[]string{"start", "A", "c", "A", "b", "A", "end"},
		[]string{"start", "A", "c", "A", "b", "end"},
		[]string{"start", "A", "c", "A", "end"},
		[]string{"start", "A", "b", "A", "c", "A", "end"},
		[]string{"start", "A", "b", "A", "end"},
		[]string{"start", "A", "b", "end"},
		[]string{"start", "A", "end"},
		[]string{"start", "b", "A", "c", "A", "end"},
		[]string{"start", "b", "A", "end"},
		[]string{"start", "b", "end"},
	}
	input := []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}

	if output := calculatePaths(input); reflect.DeepEqual(output, expected) == false {
		t.Errorf("Expected %#v, got %#v", expected, output)
	}

	input = []string{
		"dc-end",
		"HN-start",
		"start-kj",
		"dc-start",
		"dc-HN",
		"LN-dc",
		"HN-end",
		"kj-sa",
		"kj-HN",
		"kj-dc",
	}

	if output := calculatePaths(input); len(output) != 19 {
		t.Errorf("Expected 19, got %#v", len(output))
	}

	input = []string{
		"fs-end",
		"he-DX",
		"fs-he",
		"start-DX",
		"pj-DX",
		"end-zg",
		"zg-sl",
		"zg-pj",
		"pj-he",
		"RW-he",
		"fs-DX",
		"pj-RW",
		"zg-RW",
		"start-pj",
		"he-WI",
		"zg-he",
		"pj-fs",
		"start-RW",
	}

	if output := calculatePaths(input); len(output) != 226 {
		t.Errorf("Expected 226, got %#v", len(output))
	}
}
