package main

import (
	"advent-2022/internal/parse"
	"testing"
)

func TestTouching(t *testing.T) {
	p := Point{0, 0}
	if p.touching(&Point{0, 0}) != true {
		t.Errorf("Same point not touching")
	}
	if p.touching(&Point{-1, 0}) != true {
		t.Errorf("Point to left not touching")
	}
	if p.touching(&Point{1, 0}) != true {
		t.Errorf("Point to right not touching")
	}
	if p.touching(&Point{0, 1}) != true {
		t.Errorf("Point up not touching")
	}
	if p.touching(&Point{0, -1}) != true {
		t.Errorf("Point down not touching")
	}
	if p.touching(&Point{-1, -1}) != true {
		t.Errorf("Point diagonal not touching")
	}
	if p.touching(&Point{-2, 0}) != false {
		t.Errorf("Shouldn't be touching")
	}
}

func TestPart1(t *testing.T) {
	input := parse.Lines("input.test.txt")
	result := part1(input)
	if result != 13 {
		t.Errorf("Expected 13 but was %d", result)
	}
}

func TestPart2(t *testing.T) {
	t.Run("input1", func(t *testing.T) {
		input := parse.Lines("input.test.txt")
		result := part2(input)
		if result != 1 {
			t.Errorf("Expected 1 but was %d", result)
		}
	})
	t.Run("input2", func(t *testing.T) {
		input := parse.Lines("input.test2.txt")
		result := part2(input)
		if result != 36 {
			t.Errorf("Expected 36 but was %d", result)
		}
	})
}
