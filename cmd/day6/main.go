package main

import (
	"advent-2022/internal/parse"
	"fmt"
	"strings"
)

func main() {
	input := parse.String("cmd/day6/input.txt")
	fmt.Printf("Part 1: %d\n", charsToFirstMarker(input, 4))
	fmt.Printf("Part 2: %d\n", charsToFirstMarker(input, 14))
}

func charsToFirstMarker(input string, markerSize int) int {
	for i := 0; i < len(input)-markerSize; i++ {
		markerCandidate := input[i : i+markerSize]
		if allUniqueChars(markerCandidate) {
			return i + markerSize
		}
	}
	panic("did not find marker")
}

func allUniqueChars(input string) bool {
	set := map[rune]bool{}
	for _, r := range input {
		if set[r] {
			return false
		}
		set[r] = true
	}
	return true
}

func charsToFirstMessageMarker2(input string) int {
	marker := ""
	for i, r := range input {
		rs := string(r)
		idx := strings.Index(marker, rs)
		if idx > -1 {
			marker = marker[idx+1:]
		}
		marker += rs
		if len(marker) == 14 {
			return i + 1
		}
	}
	panic("did not find marker")
}
