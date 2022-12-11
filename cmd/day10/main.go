package main

import (
	"advent-2022/internal/parse"
	"fmt"
	"strings"
)

func main() {
	input := parse.Lines("cmd/day10/input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	part2(input)
}

func part1(input []string) int {
	cycles := run(input)
	strengthSum := 0
	strengthSum += 20 * cycles[19]
	strengthSum += 60 * cycles[59]
	strengthSum += 100 * cycles[99]
	strengthSum += 140 * cycles[139]
	strengthSum += 180 * cycles[179]
	strengthSum += 220 * cycles[219]
	return strengthSum
}

func part2(input []string) {
	cycles := run(input)
	horizontalPos := 0
	for i, register := range cycles {
		horizontalPos = i % 40
		if horizontalPos == 0 {
			fmt.Println()
		}
		if abs(register-horizontalPos) <= 1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func run(input []string) []int {
	var cycles []int
	curValue := 1
	for _, inst := range input {
		if strings.HasPrefix(inst, "noop") {
			cycles = append(cycles, curValue)
		} else {
			parts := strings.Split(inst, " ")
			cycles = append(cycles, curValue)
			cycles = append(cycles, curValue)
			curValue += parse.Int(parts[1])
		}
	}
	return cycles
}
