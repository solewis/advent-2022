package main

import (
	"advent-2022/internal/parse"
	"fmt"
	"sort"
)

func main() {
	input := parse.Lines("cmd/day1/input.txt")
	elfCalories := parseElfCalories(input)

	fmt.Printf("Part 1: %d\n", maxCalories(elfCalories))
	fmt.Printf("Part 2: %d\n", caloriesOfTopThree(elfCalories))
}

func parseElfCalories(lines []string) []int {
	var calories []int
	current := 0
	for _, line := range lines {
		if line == "" {
			calories = append(calories, current)
			current = 0
			continue
		}
		current += parse.Int(line)
	}
	return calories
}

func maxCalories(elfCalories []int) int {
	max := 0
	for _, elfCals := range elfCalories {
		if elfCals > max {
			max = elfCals
		}
	}
	return max
}

func caloriesOfTopThree(elfCalories []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))
	return elfCalories[0] + elfCalories[1] + elfCalories[2]
}
