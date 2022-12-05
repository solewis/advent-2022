package main

import (
	"advent-2022/internal/parse"
	"fmt"
	"regexp"
)

var (
	stack1 = []string{"J", "H", "P", "M", "S", "F", "N", "V"}
	stack2 = []string{"S", "R", "L", "M", "J", "D", "Q"}
	stack3 = []string{"N", "Q", "D", "H", "C", "S", "W", "B"}
	stack4 = []string{"R", "S", "C", "L"}
	stack5 = []string{"M", "V", "T", "P", "F", "B"}
	stack6 = []string{"T", "R", "Q", "N", "C"}
	stack7 = []string{"G", "V", "R"}
	stack8 = []string{"C", "Z", "S", "P", "D", "L", "R"}
	stack9 = []string{"D", "S", "J", "V", "G", "P", "B", "F"}
)

func main() {
	input := parse.Lines("cmd/day5/input.txt")
	stacks := [][]string{stack1, stack2, stack3, stack4, stack5, stack6, stack7, stack8, stack9}
	// Can only run one at a time because it modifies the input stacks
	//fmt.Printf("Part 1: %s\n", part1(input, stacks))
	fmt.Printf("Part 2: %s\n", part2(input, stacks))
}

func part1(procedure []string, stacks [][]string) string {
	lineRegex := regexp.MustCompile(`^move (\d+) from (\d) to (\d)$`)
	for _, move := range procedure {
		matches := lineRegex.FindStringSubmatch(move)
		count := parse.Int(matches[1])
		from := parse.Int(matches[2]) - 1
		to := parse.Int(matches[3]) - 1
		for i := 0; i < count; i++ {
			index := len(stacks[from]) - 1
			elem := stacks[from][index]
			stacks[from] = stacks[from][:index]
			stacks[to] = append(stacks[to], elem)
		}
	}
	result := ""
	for _, s := range stacks {
		result += s[len(s)-1]
	}
	return result
}

func part2(procedure []string, stacks [][]string) string {
	lineRegex := regexp.MustCompile(`^move (\d+) from (\d) to (\d)$`)
	for _, move := range procedure {
		matches := lineRegex.FindStringSubmatch(move)
		count := parse.Int(matches[1])
		from := parse.Int(matches[2]) - 1
		to := parse.Int(matches[3]) - 1
		index := len(stacks[from]) - count
		elems := stacks[from][index:]
		stacks[from] = stacks[from][:index]
		stacks[to] = append(stacks[to], elems...)
	}
	result := ""
	for _, s := range stacks {
		result += s[len(s)-1]
	}
	return result
}
