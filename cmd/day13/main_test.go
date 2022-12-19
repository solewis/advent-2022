package main

import (
	"advent-2022/internal/parse"
	"reflect"
	"testing"
)

func TestPart1(t *testing.T) {
	input := parse.String("input.test.txt")
	result := part1(input)
	if result != 13 {
		t.Errorf("Expected 13 but was %d", result)
	}
}

func TestPart2(t *testing.T) {
	input := parse.String("input.test.txt")
	result := part2(input)
	if result != 140 {
		t.Errorf("Expected 140 but was %d", result)
	}
}

func TestSplit(t *testing.T) {
	t.Run("Parses ints", func(t *testing.T) {
		result := split("[3,9]")
		expected := []string{"3", "9"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Wrong split, wanted %v, got %v", expected, result)
		}
	})
	t.Run("Parses lists", func(t *testing.T) {
		result := split("[[3],[9]]")
		expected := []string{"[3]", "[9]"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Wrong split, wanted %v, got %v", expected, result)
		}
	})
	t.Run("Parses nested lists and ints", func(t *testing.T) {
		result := split("[[3],[[9,[3]]],3]")
		expected := []string{"[3]", "[[9,[3]]]", "3"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Wrong split, wanted %v, got %v", expected, result)
		}
	})
}
