package main

import (
	"testing"
)

var (
	testMod     = 23 * 19 * 13 * 17
	testMonkeys = []*Monkey{
		{
			Items: []int{79, 98},
			Op: func(v int) int {
				return v * 19
			},
			Test: func(v int) int {
				if v%23 == 0 {
					return 2
				}
				return 3
			},
		},
		{
			Items: []int{54, 65, 75, 74},
			Op: func(v int) int {
				return v + 6
			},
			Test: func(v int) int {
				if v%19 == 0 {
					return 2
				}
				return 0
			},
		},
		{
			Items: []int{79, 60, 97},
			Op: func(v int) int {
				return v * v
			},
			Test: func(v int) int {
				if v%13 == 0 {
					return 1
				}
				return 3
			},
		},
		{
			Items: []int{74},
			Op: func(v int) int {
				return v + 3
			},
			Test: func(v int) int {
				if v%17 == 0 {
					return 0
				}
				return 1
			},
		},
	}
)

func TestPart1(t *testing.T) {
	result := part1(testMonkeys)
	if result != 10605 {
		t.Errorf("Expected 10605 but was %d", result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testMonkeys, testMod)
	if result != 2713310158 {
		t.Errorf("Expected 2713310158 but was %d", result)
	}
}
