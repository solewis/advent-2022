package main

import (
	"advent-2022/internal/parse"
	"testing"
)

func TestPart1(t *testing.T) {
	input := parse.String("input.test.txt")
	result := part1(input)
	if result != 3068 {
		t.Errorf("Expected 3068 but was %d", result)
	}
}

func TestPart2(t *testing.T) {
	input := parse.String("input.test.txt")
	result := part2(input)
	if result != 1514285714288 {
		t.Errorf("Expected 1514285714288 but was %d", result)
	}
}
