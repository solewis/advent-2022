package main

import (
	"advent-2022/internal/parse"
	"fmt"
	"regexp"
)

func main() {
	input := parse.Lines("cmd/day4/input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	lineRegex := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)
	countOverlap := 0
	for _, row := range input {
		matches := lineRegex.FindStringSubmatch(row)
		pair1Start := parse.Int(matches[1])
		pair1End := parse.Int(matches[2])
		pair2Start := parse.Int(matches[3])
		pair2End := parse.Int(matches[4])
		if pair1Start >= pair2Start && pair1End <= pair2End {
			countOverlap++
		} else if pair2Start >= pair1Start && pair2End <= pair1End {
			countOverlap++
		}
	}
	return countOverlap
}

func part2(input []string) int {
	lineRegex := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)
	countOverlap := 0
	for _, row := range input {
		matches := lineRegex.FindStringSubmatch(row)
		pair1Start := parse.Int(matches[1])
		pair1End := parse.Int(matches[2])
		pair2Start := parse.Int(matches[3])
		pair2End := parse.Int(matches[4])
		if (pair1Start >= pair2Start && pair1Start <= pair2End) || (pair1End >= pair2Start && pair1End <= pair2End) {
			countOverlap++
		} else if (pair2Start >= pair1Start && pair2Start <= pair1End) || (pair2End >= pair1Start && pair2End <= pair1End) {
			countOverlap++
		}
	}
	return countOverlap
}
