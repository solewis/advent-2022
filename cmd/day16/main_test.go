package main

import (
	"advent-2022/internal/parse"
	"testing"
)

func TestPart1(t *testing.T) {
	input := parse.Lines("input.test.txt")
	result := part1(input)
	if result != 1651 {
		t.Errorf("Expected 1651 but was %d", result)
	}
}
