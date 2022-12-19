package main

import (
	"advent-2022/internal/parse"
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := parse.String("cmd/day13/input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input string) int {
	packets := parsePackets(input)
	sum := 0
	for index, i := 1, 0; i < len(packets)-1; i, index = i+2, index+1 {
		ppResult := compare(packets[i], packets[i+1])
		if ppResult == 1 {
			sum += index
		}
	}
	return sum
}

func part2(input string) int {
	packets := parsePackets(input)
	packets = append(packets, []string{"[[2]]", "[[6]]"}...)
	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) > 0
	})
	div1, div2 := 0, 0
	for index, packet := range packets {
		if packet == "[[2]]" {
			div1 = index + 1
		}
		if packet == "[[6]]" {
			div2 = index + 1
		}
	}
	return div1 * div2
}

func isList(val string) bool {
	return strings.HasPrefix(val, "[")
}

func split(val string) []string {
	val = val[1 : len(val)-1] //strip outer []
	var result []string
	openParenCount := 0
	nextVal := ""
	for _, r := range val {
		if r == '[' {
			openParenCount++
			nextVal += string(r)
		} else if r == ']' {
			openParenCount--
			nextVal += string(r)
		} else if r == ',' {
			if openParenCount == 0 {
				result = append(result, nextVal)
				nextVal = ""
			} else {
				nextVal += string(r)
			}
		} else {
			nextVal += string(r)
		}
	}
	if nextVal != "" {
		result = append(result, nextVal)
	}
	return result
}

// return -1 for incorrect order, 1 for correct, 0 for unsolved
func compare(left, right string) int {
	if !isList(left) && !isList(right) {
		newLeft := parse.Int(left)
		newRight := parse.Int(right)
		if newLeft < newRight {
			return 1
		} else if newRight < newLeft {
			return -1
		} else {
			return 0
		}
	} else if isList(left) && isList(right) {
		leftSplit := split(left)
		rightSplit := split(right)
		i := 0
		for {
			if i >= len(leftSplit) || i >= len(rightSplit) {
				if len(leftSplit) < len(rightSplit) {
					return 1
				} else if len(rightSplit) < len(leftSplit) {
					return -1
				} else {
					return 0
				}
			}
			newLeft := leftSplit[i]
			newRight := rightSplit[i]
			compareResult := compare(newLeft, newRight)
			if compareResult != 0 {
				return compareResult
			}
			i++
		}
	} else {
		newLeft := left
		newRight := right
		if isList(right) {
			newLeft = "[" + newLeft + "]"
		} else {
			newRight = "[" + newRight + "]"
		}
		result := compare(newLeft, newRight)
		if result != 0 {
			return result
		}
	}
	return 0
}

func parsePackets(input string) []string {
	var packets []string
	pairs := strings.Split(input, "\n\n")
	for _, pair := range pairs {
		parts := strings.Split(pair, "\n")
		packets = append(packets, parts[0])
		packets = append(packets, parts[1])
	}
	return packets
}
