package main

import (
	"advent-2022/internal/parse"
	"testing"
)

var (
	testStack1 = []string{"Z", "N"}
	testStack2 = []string{"M", "C", "D"}
	testStack3 = []string{"P"}
)

func TestPart1(t *testing.T) {
	input := parse.Lines("input.test.txt")
	stacks := [][]string{testStack1, testStack2, testStack3}
	result := part1(input, stacks)
	if result != "CMZ" {
		t.Errorf("Expected CMZ, but was %s", result)
	}
}

func TestPart2(t *testing.T) {
	input := parse.Lines("input.test.txt")
	stacks := [][]string{testStack1, testStack2, testStack3}
	result := part2(input, stacks)
	if result != "MCD" {
		t.Errorf("Expected MCD, but was %s", result)
	}
}
