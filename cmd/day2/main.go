package main

import (
	"advent-2022/internal/parse"
	"fmt"
)

var (
	scoresPart1 = map[string]int{
		"A X": 4, // 1 + 3
		"A Y": 8, // 2 + 6
		"A Z": 3, // 3 + 0
		"B X": 1, // 1 + 0
		"B Y": 5, // 2 + 3
		"B Z": 9, // 3 + 6
		"C X": 7, // 1 + 6
		"C Y": 2, // 2 + 0
		"C Z": 6, // 3 + 3
	}
	scoresPart2 = map[string]int{
		"A X": 3, // 0 + 3
		"A Y": 4, // 3 + 1
		"A Z": 8, // 6 + 2
		"B X": 1, // 0 + 1
		"B Y": 5, // 3 + 2
		"B Z": 9, // 6 + 3
		"C X": 2, // 0 + 2
		"C Y": 6, // 3 + 3
		"C Z": 7, // 6 + 1
	}
)

func main() {
	input := parse.Lines("cmd/day2/input.txt")
	fmt.Printf("Part 1: %d\n", calculateScorePart1(input))
	fmt.Printf("Part 2: %d\n", calculateScorePart2(input))
}

func calculateScorePart1(games []string) int {
	score := 0
	for _, game := range games {
		score += scoresPart1[game]
	}
	return score
}

func calculateScorePart2(games []string) int {
	score := 0
	for _, game := range games {
		score += scoresPart2[game]
	}
	return score
}
