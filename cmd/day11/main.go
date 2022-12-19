package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Printf("Part 1: %d\n", part1(monkeys))
	// part 1 and 2 can't both run due to modifying the input
	fmt.Printf("Part 2: %d\n", part2(monkeys, inputMod))
}

func part1(monkeys []*Monkey) int {
	simulate(monkeys, 20, 0)
	var counts []int
	for _, m := range monkeys {
		counts = append(counts, m.ItemCount)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	return counts[0] * counts[1]
}

func part2(monkeys []*Monkey, mod int) int {
	simulate(monkeys, 10000, mod)
	var counts []int
	for _, m := range monkeys {
		counts = append(counts, m.ItemCount)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	return counts[0] * counts[1]
}

type Monkey struct {
	Items     []int
	ItemCount int
	Op        func(v int) int
	Test      func(v int) int
}

func simulate(monkeys []*Monkey, rounds int, mod int) {
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			monkey.ItemCount += len(monkey.Items)
			for _, item := range monkey.Items {
				worry := monkey.Op(item)
				if mod == 0 {
					worry /= 3
				} else {
					worry %= mod
				}
				next := monkey.Test(worry)
				monkeys[next].Items = append(monkeys[next].Items, worry)
			}
			monkey.Items = nil
		}
	}
}

//func parseMonkeys(input []string) []Monkey {
//	var monkeys []Monkey
//	cur := Monkey{}
//	for _, line := range input {
//		if line == "" {
//			monkeys = append(monkeys, cur)
//			cur = Monkey{}
//			continue
//		}
//		if strings.HasPrefix(line, "Monkey") {
//			continue
//		}
//		if strings.HasPrefix(line, "  Starting items") {
//			parts := strings.Split(line, ": ")
//			itemStrs := strings.Split(parts[1], ", ")
//			var items []int
//			for _, i := range itemStrs {
//				items = append(items, parse.Int(i))
//			}
//			cur.Items = items
//		}
//		if strings.HasPrefix(line, "  Operation:") {
//			parts := strings.Split(line, " = ")
//			opStrs := strings.Split(parts[1], " ")
//
//
//			}
//		}
//	}
//}

var (
	inputMod = 2 * 7 * 13 * 5 * 3 * 19 * 11 * 17
	monkeys  = []*Monkey{
		{
			Items:     []int{85, 79, 63, 72},
			ItemCount: 0,
			Op: func(v int) int {
				return v * 17
			},
			Test: func(v int) int {
				if v%2 == 0 {
					return 2
				}
				return 6
			},
		},
		{
			Items:     []int{53, 94, 65, 81, 93, 73, 57, 92},
			ItemCount: 0,
			Op: func(v int) int {
				return v * v
			},
			Test: func(v int) int {
				if v%7 == 0 {
					return 0
				}
				return 2
			},
		},
		{
			Items:     []int{62, 63},
			ItemCount: 0,
			Op: func(v int) int {
				return v + 7
			},
			Test: func(v int) int {
				if v%13 == 0 {
					return 7
				}
				return 6
			},
		},
		{
			Items:     []int{57, 92, 56},
			ItemCount: 0,
			Op: func(v int) int {
				return v + 4
			},
			Test: func(v int) int {
				if v%5 == 0 {
					return 4
				}
				return 5
			},
		},
		{
			Items:     []int{67},
			ItemCount: 0,
			Op: func(v int) int {
				return v + 5
			},
			Test: func(v int) int {
				if v%3 == 0 {
					return 1
				}
				return 5
			},
		},
		{
			Items:     []int{85, 56, 66, 72, 57, 99},
			ItemCount: 0,
			Op: func(v int) int {
				return v + 6
			},
			Test: func(v int) int {
				if v%19 == 0 {
					return 1
				}
				return 0
			},
		},
		{
			Items:     []int{86, 65, 98, 97, 69},
			ItemCount: 0,
			Op: func(v int) int {
				return v * 13
			},
			Test: func(v int) int {
				if v%11 == 0 {
					return 3
				}
				return 7
			},
		},
		{
			Items:     []int{87, 68, 92, 66, 91, 50, 68},
			ItemCount: 0,
			Op: func(v int) int {
				return v + 2
			},
			Test: func(v int) int {
				if v%17 == 0 {
					return 4
				}
				return 3
			},
		},
	}
)
