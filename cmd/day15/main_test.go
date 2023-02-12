package main

import (
	"advent-2022/internal/parse"
	"testing"
)

func TestPart1(t *testing.T) {
	input := parse.Lines("input.test.txt")
	result := part1(input, 10)
	if result != 26 {
		t.Errorf("Expected 26 but was %d", result)
	}
}

func TestPart2(t *testing.T) {
	input := parse.Lines("input.test.txt")
	result := part2(input, 20)
	if result != 56000011 {
		t.Errorf("Expected 56000011, but was %d", result)
	}
}
