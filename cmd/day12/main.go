package main

import (
	"advent-2022/internal/parse"
	"fmt"
	"math"
)

func main() {
	input := parse.Lines("cmd/day12/input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type Point struct {
	row, col int
}

func (p Point) adjacent() []Point {
	return []Point{
		{p.row - 1, p.col},
		{p.row, p.col - 1},
		{p.row, p.col + 1},
		{p.row + 1, p.col},
	}
}

func part1(input []string) int {
	var start, end Point
	var grid [][]rune
	for row, line := range input {
		var gridRow []rune
		for col, val := range line {
			if val == 'S' {
				start = Point{row, col}
				val = 'a'
			}
			if val == 'E' {
				end = Point{row, col}
				val = 'z'
			}
			gridRow = append(gridRow, val)
		}
		grid = append(grid, gridRow)
	}
	return shortestPath(start, end, grid)
}

func part2(input []string) int {
	var starts []Point
	var end Point
	var grid [][]rune
	for row, line := range input {
		var gridRow []rune
		for col, val := range line {
			if val == 'S' || val == 'a' {
				starts = append(starts, Point{row, col})
				val = 'a'
			}
			if val == 'E' {
				end = Point{row, col}
				val = 'z'
			}
			gridRow = append(gridRow, val)
		}
		grid = append(grid, gridRow)
	}
	shortestDistance := math.MaxInt
	for _, start := range starts {
		dist := shortestPath(start, end, grid)
		if dist < shortestDistance {
			shortestDistance = dist
		}
	}
	return shortestDistance
}

// bfs for shortest path
func shortestPath(start, end Point, grid [][]rune) int {
	distance := map[Point]int{start: 0}
	var queue []Point
	queue = append(queue, start)
	for len(queue) > 0 {
		curPos := queue[0]
		curDistance := distance[curPos]
		queue = queue[1:]
		if curPos == end {
			return curDistance
		}
		for _, nextPos := range curPos.adjacent() {
			if allowedStep(curPos, nextPos, grid, distance) {
				queue = append(queue, nextPos)
				distance[nextPos] = curDistance + 1
			}
		}
	}
	return math.MaxInt
}

func allowedStep(pos, next Point, grid [][]rune, distance map[Point]int) bool {
	_, visited := distance[next]
	return !visited && isInGrid(next, grid) &&
		grid[next.row][next.col]-grid[pos.row][pos.col] <= 1
}

func isInGrid(pos Point, grid [][]rune) bool {
	if pos.row < 0 || pos.row >= len(grid) {
		return false
	}
	if pos.col < 0 || pos.col >= len(grid[0]) {
		return false
	}
	return true
}
