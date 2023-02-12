package main

import (
	"advent-2022/internal/parse"
	"fmt"
	"math"
	"regexp"
)

func main() {
	input := parse.Lines("cmd/day15/input.txt")
	fmt.Printf("Part 1: %d\n", part1(input, 2000000))
	fmt.Printf("Part 2: %d\n", part2(input, 4000000))
}

func part1(input []string, y int) int {
	sensorData := parseSensorData(input)
	minX, maxX := math.MaxInt, -math.MaxInt
	for s, b := range sensorData {
		if s.x < minX {
			minX = s.x
		}
		if s.x > maxX {
			maxX = s.x
		}
		if b.x < minX {
			minX = b.x
		}
		if b.x > maxX {
			maxX = b.x
		}
	}
	width := maxX - minX
	result := 0
outer:
	for x := minX - width; x < maxX+width; x++ {
		p := Point{x, y}
		for s, b := range sensorData {
			if s == p || b == p {
				continue outer
			}
			distToSensor := p.manhattanDistance(s)
			sensorToBeacon := s.manhattanDistance(b)
			if distToSensor <= sensorToBeacon {
				result++
				continue outer
			}
		}
	}
	return result
}

func part2(input []string, size int) int {
	sensorData := parseSensorData(input)
	// The beacon location must be on the edge of one of the sensors covered area, so we just look at those points
	// For each point along one sensor border, check to see if no other sensor overlaps, if so we found the beacon
	for s, b := range sensorData {
		dis := s.manhattanDistance(b)
		leftX, rightX := s.x, s.x
		for y := s.y - dis - 1; y <= s.y+dis+1; y++ {
			if y >= 0 && y <= size {
				if leftX >= 0 && leftX <= size {
					p := Point{leftX, y}
					found := true
					for s1, b1 := range sensorData {
						if s1 == p || b1 == p {
							found = false
							break
						}
						distToSensor := p.manhattanDistance(s1)
						sensorToBeacon := s1.manhattanDistance(b1)
						if distToSensor <= sensorToBeacon {
							found = false
							break
						}
					}
					if found {
						return p.x*4000000 + p.y
					}
				}

				if rightX >= 0 && rightX <= size {
					p := Point{rightX, y}
					found := true
					for s1, b1 := range sensorData {
						if s1 == p || b1 == p {
							found = false
							break
						}
						distToSensor := p.manhattanDistance(s1)
						sensorToBeacon := s1.manhattanDistance(b1)
						if distToSensor <= sensorToBeacon {
							found = false
							break
						}
					}
					if found {
						return p.x*4000000 + p.y
					}
				}
			}
			if y < s.y {
				leftX--
				rightX++
			} else {
				leftX++
				rightX--
			}
		}
	}
	panic("not found")
}

type Point struct {
	x, y int
}

func (p Point) manhattanDistance(o Point) int {
	return abs(p.x-o.x) + abs(p.y-o.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func parseSensorData(input []string) map[Point]Point {
	data := map[Point]Point{}
	lineRegex := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)
	for _, line := range input {
		matches := lineRegex.FindStringSubmatch(line)
		sensor := Point{parse.Int(matches[1]), parse.Int(matches[2])}
		beacon := Point{parse.Int(matches[3]), parse.Int(matches[4])}
		data[sensor] = beacon
	}
	return data
}
