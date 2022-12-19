package main

import (
	"advent-2022/internal/parse"
	"testing"
)

func TestPart1(t *testing.T) {
	input := parse.Lines("input.test.txt")
	result := part1(input)
	if result != 24 {
		t.Errorf("Expected 24 but was %d", result)
	}
}

func TestPart2(t *testing.T) {
	input := parse.Lines("input.test.txt")
	result := part2(input)
	if result != 93 {
		t.Errorf("Expected 93 but was %d", result)
	}
}
