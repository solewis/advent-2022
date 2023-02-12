package main

import (
	"advent-2022/internal/parse"
	"fmt"
)

func main() {
	input := parse.String("cmd/day17/input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	start := 1 + 3 + 2 + 1 + 2 + 1 + 3 + 2 + 2 + 1 + 3 + 2 + 2 + 1 + 3 + 3
	startLen := 16
	cycle := 10
	cycleLen := 26
	noStart := 1000000000000 - startLen
	cycles := noStart / cycleLen
	remainder := noStart % cycleLen
	fmt.Println("Cycles:", cycles)
	fmt.Println("Remainder:", remainder)
	cycleTotal := cycles * cycle
	remainderTotal := 4 + 1 + 2 + 3 + 1 + 1 + 3 + 2 + 2 + 2 + 3 + 4 + 1 + 2 + 1 + 2 + 1 + 2 + 1 + 2 + 1 + 3 + 2 + 1
	fmt.Println("Part 2:", start+cycleTotal+remainderTotal)
}

type Point struct {
	x, y int
}

type Rock struct {
	positions map[Point]bool
}

func (r Rock) MoveLR(dir string, chamber map[Point]bool) Rock {
	newPos := map[Point]bool{}
	var chg int
	switch dir {
	case ">":
		chg = 1
	case "<":
		chg = -1
	default:
		panic("invalid dir")
	}
	for pos := range r.positions {
		newPt := Point{pos.x + chg, pos.y}
		if newPt.x < 1 || newPt.x > 7 || chamber[newPt] {
			return r
		}
		newPos[newPt] = true
	}
	return Rock{newPos}
}

func (r Rock) MoveDown(chamber map[Point]bool) (Rock, bool) {
	newPos := map[Point]bool{}
	for pos := range r.positions {
		newPt := Point{pos.x, pos.y - 1}
		if chamber[newPt] {
			return r, false
		}
		newPos[newPt] = true
	}
	return Rock{newPos}, true
}

func CreateRock(rockType int, baseLevel int) Rock {
	switch rockType {
	case 0:
		return Rock{map[Point]bool{
			Point{3, baseLevel + 4}: true,
			Point{4, baseLevel + 4}: true,
			Point{5, baseLevel + 4}: true,
			Point{6, baseLevel + 4}: true,
		}}
	case 1:
		return Rock{map[Point]bool{
			Point{4, baseLevel + 4}: true,
			Point{3, baseLevel + 5}: true,
			Point{4, baseLevel + 5}: true,
			Point{5, baseLevel + 5}: true,
			Point{4, baseLevel + 6}: true,
		}}
	case 2:
		return Rock{map[Point]bool{
			Point{3, baseLevel + 4}: true,
			Point{4, baseLevel + 4}: true,
			Point{5, baseLevel + 4}: true,
			Point{5, baseLevel + 5}: true,
			Point{5, baseLevel + 6}: true,
		}}
	case 3:
		return Rock{map[Point]bool{
			Point{3, baseLevel + 4}: true,
			Point{3, baseLevel + 5}: true,
			Point{3, baseLevel + 6}: true,
			Point{3, baseLevel + 7}: true,
		}}
	case 4:
		return Rock{map[Point]bool{
			Point{3, baseLevel + 4}: true,
			Point{3, baseLevel + 5}: true,
			Point{4, baseLevel + 4}: true,
			Point{4, baseLevel + 5}: true,
		}}
	default:
		panic("invalid rock type")
	}
}

func part1(input string) int {
	return run(input, 2022)
}

func part2(input string) int {
	return run(input, 1000000000000)
}

func run(input string, rocks int) int {
	chamber := map[Point]bool{
		Point{1, 0}: true,
		Point{2, 0}: true,
		Point{3, 0}: true,
		Point{4, 0}: true,
		Point{5, 0}: true,
		Point{6, 0}: true,
		Point{7, 0}: true,
	}
	jetIndex := 0
	rockType := 0
	baseLevel := 0

	for i := 0; i < rocks; i++ {
		rock := CreateRock(rockType, baseLevel)
		for {
			rockNext := rock.MoveLR(string(input[jetIndex]), chamber)
			rockDown, moved := rockNext.MoveDown(chamber)
			jetIndex = (jetIndex + 1) % len(input)
			if !moved {
				rockType = (rockType + 1) % 5
				oldBaseLevel := baseLevel
				for pos := range rockDown.positions {
					if pos.y > baseLevel {
						baseLevel = pos.y
					}
					chamber[pos] = true
				}
				fmt.Println("Height change:", baseLevel-oldBaseLevel)
				break
			}
			rock = rockDown
		}
	}
	return baseLevel
}
