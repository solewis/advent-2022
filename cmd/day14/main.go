package main

import (
	"advent-2022/internal/parse"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := parse.Lines("cmd/day14/input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	walls := parseWalls(input)
	sand := simulate(walls)
	printCave(walls, sand)
	return len(sand)
}

func part2(input []string) int {
	walls := parseWalls(input)
	sand := simulateWithFloor(walls)
	return len(sand)
}

type Point struct {
	x, y int
}

func (p Point) stepTowards(o Point) Point {
	next := Point{p.x, p.y}
	if p.x < o.x {
		next.x++
	}
	if p.x > o.x {
		next.x--
	}
	if p.y < o.y {
		next.y++
	}
	if p.y > o.y {
		next.y--
	}
	return next
}

func parseWalls(input []string) map[Point]bool {
	walls := map[Point]bool{}
	for _, line := range input {
		points := strings.Split(line, " -> ")
		index := 0
		from := points[index]
		fromCoords := strings.Split(from, ",")
		fromPoint := Point{parse.Int(fromCoords[0]), parse.Int(fromCoords[1])}
		walls[fromPoint] = true
		for {
			to := points[index+1]
			toCoords := strings.Split(to, ",")
			toPoint := Point{parse.Int(toCoords[0]), parse.Int(toCoords[1])}
			for fromPoint != toPoint {
				fromPoint = fromPoint.stepTowards(toPoint)
				walls[fromPoint] = true
			}
			index++
			if index > len(points)-2 {
				break
			}
		}
	}
	return walls
}

func simulate(walls map[Point]bool) map[Point]bool {
	sand := map[Point]bool{}
	maxY := 0
	for p := range walls {
		if p.y > maxY {
			maxY = p.y
		}
	}
	sandUnit := Point{500, 0}
	for {
		if sandUnit.y > maxY {
			return sand
		}
		down := Point{sandUnit.x, sandUnit.y + 1}
		if !walls[down] && !sand[down] {
			sandUnit = down
			continue
		}
		downLeft := Point{sandUnit.x - 1, sandUnit.y + 1}
		if !walls[downLeft] && !sand[downLeft] {
			sandUnit = downLeft
			continue
		}
		downRight := Point{sandUnit.x + 1, sandUnit.y + 1}
		if !walls[downRight] && !sand[downRight] {
			sandUnit = downRight
			continue
		}
		sand[sandUnit] = true
		sandUnit = Point{500, 0}
	}
}

func simulateWithFloor(walls map[Point]bool) map[Point]bool {
	sand := map[Point]bool{}
	maxY := 0
	for p := range walls {
		if p.y > maxY {
			maxY = p.y
		}
	}
	floor := maxY + 2
	source := Point{500, 0}
	sandUnit := source
	for {
		down := Point{sandUnit.x, sandUnit.y + 1}
		if !walls[down] && !sand[down] && !(down.y == floor) {
			sandUnit = down
			continue
		}
		downLeft := Point{sandUnit.x - 1, sandUnit.y + 1}
		if !walls[downLeft] && !sand[downLeft] && !(downLeft.y == floor) {
			sandUnit = downLeft
			continue
		}
		downRight := Point{sandUnit.x + 1, sandUnit.y + 1}
		if !walls[downRight] && !sand[downRight] && !(downRight.y == floor) {
			sandUnit = downRight
			continue
		}
		sand[sandUnit] = true
		if sandUnit == source {
			return sand
		}
		sandUnit = source
	}
}

func printCave(walls map[Point]bool, sand map[Point]bool) {
	minX, maxX, minY, maxY := math.MaxInt, -math.MaxInt, math.MaxInt, -math.MaxInt
	for p := range walls {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	for p := range sand {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if walls[Point{x, y}] {
				fmt.Print("#")
			} else if sand[Point{x, y}] {
				fmt.Print("o")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
