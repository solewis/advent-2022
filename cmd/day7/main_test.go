package main

import (
	"advent-2022/internal/parse"
	"testing"
)

func TestPart1(t *testing.T) {
	input := parse.Lines("input.test.txt")
	result := Part1(input)
	if result != 95437 {
		t.Errorf("Expected 95437 but was %d", result)
	}
}

func TestPart2(t *testing.T) {
	input := parse.Lines("input.test.txt")
	result := Part2(input)
	if result != 24933642 {
		t.Errorf("Expected 24933642 but was %d", result)
	}
}
