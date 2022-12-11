package main

import (
	"advent-2022/internal/parse"
	"testing"
)

func TestCycles(t *testing.T) {
	input := parse.Lines("input.test.txt")
	cycles := run(input)
	if cycles[19] != 21 {
		t.Errorf("expected 20th cycle to be 21, but was %d", cycles[19])
	}
	if cycles[59] != 19 {
		t.Errorf("expected 60th cycle to be 19, but was %d", cycles[59])
	}
	if cycles[99] != 18 {
		t.Errorf("expected 100th cycle to be 18, but was %d", cycles[99])
	}
	if cycles[139] != 21 {
		t.Errorf("expected 140th cycle to be 21, but was %d", cycles[139])
	}
	if cycles[179] != 16 {
		t.Errorf("expected 180th cycle to be 16, but was %d", cycles[179])
	}
	if cycles[219] != 18 {
		t.Errorf("expected 220th cycle to be 18, but was %d", cycles[219])
	}
}

func TestPart1(t *testing.T) {
	input := parse.Lines("input.test.txt")
	result := part1(input)
	if result != 13140 {
		t.Errorf("Expected 13140 but was %d", result)
	}
}
