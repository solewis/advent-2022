package parse

import (
	"os"
	"strconv"
	"strings"
)

func Lines(filename string) []string {
	dat, err := os.ReadFile(filename)
	check(err)
	lines := strings.Split(string(dat), "\n")
	return lines[:len(lines)-1] // strip final line which will be blank
}

func String(filename string) string {
	dat, err := os.ReadFile(filename)
	check(err)
	return string(dat)[:len(dat)-1]
}

func Ints(filename, separator string) []int {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var ints []int
	for _, l := range strings.Split(string(dat), separator) {
		ints = append(ints, Int(l))
	}
	return ints
}

func Int(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
