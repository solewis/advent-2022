package main

import (
	"advent-2022/internal/parse"
	"regexp"
	"strings"
)

func main() {

}

func part1(input []string) int {
	valves := parseValves(input)
	return simulate(valves, map[string]Valve{}, valves["AA"], 0, 0)
}

type Valve struct {
	name     string
	rate     int
	children []string
}

type State struct {
	valve       string
	minutesLeft int
}

var (
	cache = map[State]int{}
)

func simulate(valves map[string]Valve, flows map[string]Valve, current Valve, acc int, minute int) int {
	if minute == 30 {
		return acc
	}
	if best, exists := cache[State{valve: current.name, minutesLeft: 30 - minute}]; exists {
		return best
	}
	addedFlow := 0
	newFlows := map[string]Valve{}
	for _, valve := range flows {
		addedFlow += valve.rate
		newFlows[valve.name] = valve
	}
	bestPath := 0
	// this path opens the current valve and moves on
	if _, exists := flows[current.name]; !exists && current.rate > 0 {
		newFlows[current.name] = current
		path := simulate(valves, newFlows, current, acc+addedFlow, minute+1)
		if path > bestPath {
			bestPath = path
		}
	}
	for _, child := range current.children {
		if _, exists := flows[child]; !exists { //don't go back to open one
			path := simulate(valves, newFlows, valves[child], acc+addedFlow, minute+1)
			if path > bestPath {
				bestPath = path
			}
		}
	}
	cache[State{valve: current.name, minutesLeft: 30 - minute}] = bestPath
	return bestPath
}

func parseValves(input []string) map[string]Valve {
	valves := map[string]Valve{}
	lineRegex := regexp.MustCompile(`^Valve (\w+) has flow rate=(\d+); tunnel(s?) lead(s?) to valve(s?) (.+)$`)
	for _, line := range input {
		matches := lineRegex.FindStringSubmatch(line)
		name := matches[1]
		rate := parse.Int(matches[2])
		children := strings.Split(matches[6], ", ")
		valves[name] = Valve{
			name:     name,
			rate:     rate,
			children: children,
		}
	}
	return valves
}
