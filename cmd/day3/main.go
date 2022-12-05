package main

import (
	"advent-2022/internal/parse"
	"fmt"
	"strings"
)

const abc = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	input := parse.Lines("cmd/day3/input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(rucksacks []string) int {
	prioritySum := 0
	for _, rucksack := range rucksacks {
		shared := findShared(rucksack)
		priority := strings.Index(abc, string(shared)) + 1
		prioritySum += priority
	}
	return prioritySum
}

func part2(rucksacks []string) int {
	prioritySum := 0
	for i := 0; i < len(rucksacks); i += 3 {
		r1 := rucksacks[i]
		r2 := rucksacks[i+1]
		r3 := rucksacks[i+2]
		badge := findBadge(r1, r2, r3)
		prioritySum += strings.Index(abc, string(badge)) + 1
	}
	return prioritySum
}

func findBadge(r1, r2, r3 string) rune {
	r1Map, r2Map := map[rune]bool{}, map[rune]bool{}
	for _, r := range r1 {
		r1Map[r] = true
	}
	for _, r := range r2 {
		r2Map[r] = true
	}
	for _, r := range r3 {
		if r1Map[r] && r2Map[r] {
			return r
		}
	}
	panic("badge not found among group")
}

func findShared(contents string) rune {
	compartment1 := contents[:len(contents)/2]
	compartment2 := contents[len(contents)/2:]
	c1Map := map[rune]bool{}
	for _, r := range compartment1 {
		c1Map[r] = true
	}
	for _, r := range compartment2 {
		if c1Map[r] {
			return r
		}
	}
	panic("did not find a shared item")
}
