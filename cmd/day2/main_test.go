package main

import "testing"

func TestTotalScorePart1(t *testing.T) {
	input := []string{"A Y", "B X", "C Z"}
	totalScore := calculateScorePart1(input)
	if totalScore != 15 {
		t.Errorf("Expected 15 but was %d", totalScore)
	}
}

func TestTotalScorePart2(t *testing.T) {
	input := []string{"A Y", "B X", "C Z"}
	totalScore := calculateScorePart2(input)
	if totalScore != 12 {
		t.Errorf("Expected 15 but was %d", totalScore)
	}
}
