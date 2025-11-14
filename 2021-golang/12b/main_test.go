// This file was written by Damien Dart, <damiendart@pobox.com>. This is
// free and unencumbered software released into the public domain. For
// more information, please refer to the accompanying "UNLICENCE file.

package main

import "testing"

func TestCalculatePaths(t *testing.T) {
	input := []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}

	if output := calculatePaths(input); len(output) != 36 {
		t.Errorf("Expected 36, got %#v", len(output))
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

	if output := calculatePaths(input); len(output) != 103 {
		t.Errorf("Expected 103, got %#v", len(output))
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

	if output := calculatePaths(input); len(output) != 3509 {
		t.Errorf("Expected 3509, got %#v", len(output))
	}
}
