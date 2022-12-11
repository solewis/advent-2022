package main

import (
	"advent-2022/internal/parse"
	"fmt"
)

func main() {
	input := parse.Lines("cmd/day8/input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	grid := parseGrid(input)
	countVisible := 0
	for rowI, row := range grid {
		for colI, col := range row {
			treesLeft := row[:colI]
			treesRight := row[colI+1:]
			var treesUp, treesDown []int
			for r := rowI - 1; r >= 0; r-- {
				treesUp = append(treesUp, grid[r][colI])
			}
			for r := rowI + 1; r < len(grid); r++ {
				treesDown = append(treesDown, grid[r][colI])
			}
			if isVisible(col, treesLeft, treesRight, treesUp, treesDown) {
				countVisible++
			}
		}
	}
	return countVisible
}

func part2(input []string) int {
	grid := parseGrid(input)
	highScore := 0
	for rowI, row := range grid {
		for colI, col := range row {
			treesLeft := reverse(row[:colI])
			treesRight := row[colI+1:]
			var treesUp, treesDown []int
			for r := rowI - 1; r >= 0; r-- {
				treesUp = append(treesUp, grid[r][colI])
			}
			for r := rowI + 1; r < len(grid); r++ {
				treesDown = append(treesDown, grid[r][colI])
			}
			score := scenicScore(col, treesLeft, treesRight, treesUp, treesDown)
			if score > highScore {
				highScore = score
			}
		}
	}
	return highScore
}

func reverse(slc []int) []int {
	var reversed []int
	for i := len(slc) - 1; i >= 0; i-- {
		reversed = append(reversed, slc[i])
	}
	return reversed
}

func scenicScore(treeHeight int, left, right, up, down []int) int {
	score := 1
	count := 0
	for _, line := range [][]int{left, right, up, down} {
		for _, t := range line {
			count++
			if t >= treeHeight {
				break
			}
		}
		score *= count
		count = 0
	}
	return score
}

func isVisible(treeHeight int, left, right, up, down []int) bool {
outer:
	for _, line := range [][]int{left, right, up, down} {
		for _, t := range line {
			if t >= treeHeight {
				continue outer
			}
		}
		return true
	}
	return false
}

func parseGrid(input []string) [][]int {
	var grid [][]int
	var row []int
	for _, line := range input {
		for _, r := range line {
			row = append(row, int(r-'0'))
		}
		grid = append(grid, row)
		row = nil
	}
	return grid
}
