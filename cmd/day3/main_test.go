package main

import (
	"advent-2022/internal/parse"
	"testing"
)

func TestPart1(t *testing.T) {
	input := parse.Lines("input.test.txt")
	result := part1(input)
	if result != 157 {
		t.Errorf("Expected 157 but was %d", result)
	}
}

func TestPart2(t *testing.T) {
	input := parse.Lines("input.test.txt")
	result := part2(input)
	if result != 70 {
		t.Errorf("Expected 70 but was %d", result)
	}
}
