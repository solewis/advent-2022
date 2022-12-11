package main

import (
	"advent-2022/internal/parse"
	"fmt"
	"strings"
)

func main() {
	input := parse.Lines("cmd/day9/input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type Point struct {
	x, y int
}

// follow moves point p to follow point o until it is touching
func (p *Point) follow(o *Point) {
	// move x and y one step closer to the point being followed until it is touching
	for !p.touching(o) {
		if p.x > o.x {
			p.x--
		} else if p.x < o.x {
			p.x++
		}
		if p.y > o.y {
			p.y--
		} else if p.y < o.y {
			p.y++
		}
	}
}

// touching returns true if o is in the same point as p or one of the 8 surrounding points
func (p *Point) touching(o *Point) bool {
	return (o.x <= (p.x+1) && o.x >= (p.x-1)) && (o.y <= (p.y+1) && o.y >= (p.y-1))
}

func (p *Point) copy() Point {
	return Point{p.x, p.y}
}

func part1(input []string) int {
	return simulateRope(input, 2)
}

func part2(input []string) int {
	return simulateRope(input, 10)
}

func simulateRope(input []string, ropeLength int) int {
	var rope []*Point
	for i := 0; i < ropeLength; i++ {
		rope = append(rope, &Point{0, 0})
	}
	visited := map[Point]bool{Point{0, 0}: true}
	for _, step := range input {
		parts := strings.Split(step, " ")
		head := rope[0]
		tail := rope[len(rope)-1]
		for i := 0; i < parse.Int(parts[1]); i++ {
			switch parts[0] {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y--
			case "D":
				head.y++
			}
			for i, knot := range rope[1:] {
				knot.follow(rope[i])
			}
			visited[tail.copy()] = true
		}
	}
	return len(visited)
}
