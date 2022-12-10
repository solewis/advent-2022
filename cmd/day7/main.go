package main

import (
	"advent-2022/internal/parse"
	"fmt"
	"strings"
)

func main() {
	input := parse.Lines("cmd/day7/input.txt")
	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}

func Part1(input []string) int {
	directoryMap := ParseFileSystem(input)
	result := 0
	for _, dir := range directoryMap {
		size := calculateDirectorySize(dir)
		if size <= 100000 {
			result += size
		}
	}
	return result
}

func Part2(input []string) int {
	directoryMap := ParseFileSystem(input)
	rootSize := calculateDirectorySize(directoryMap["/"])
	spaceNeeded := 30000000 - (70000000 - rootSize)
	currentSizeToDelete := 70000000
	for _, dir := range directoryMap {
		size := calculateDirectorySize(dir)
		if size > spaceNeeded && size < currentSizeToDelete {
			currentSizeToDelete = size
		}
	}
	return currentSizeToDelete
}

type Directory struct {
	Name     string
	Parent   *Directory
	Children []*Directory
	Files    []File
}

type File struct {
	Name string
	Size int
}

func ParseFileSystem(output []string) map[string]*Directory {
	directoryMap := map[string]*Directory{}
	var currentDir *Directory
	for _, line := range output {
		if strings.HasPrefix(line, "$") {
			if line == "$ ls" {
				continue
			} else if line == "$ cd .." {
				currentDir = currentDir.Parent
			} else {
				var dirName string
				if currentDir == nil {
					dirName = strings.Split(line, " ")[2]
				} else {
					dirName = currentDir.Name + strings.Split(line, " ")[2]
				}
				newDir := &Directory{Name: dirName, Parent: currentDir}
				directoryMap[dirName] = newDir
				if currentDir != nil {
					currentDir.Children = append(currentDir.Children, newDir)
				}
				currentDir = newDir
			}
			continue
		}
		if !strings.HasPrefix(line, "dir") {
			parts := strings.Split(line, " ")
			currentDir.Files = append(currentDir.Files, File{
				Name: parts[1],
				Size: parse.Int(parts[0]),
			})
		}
	}
	return directoryMap
}

func calculateDirectorySize(dir *Directory) int {
	size := 0
	for _, file := range dir.Files {
		size += file.Size
	}
	for _, child := range dir.Children {
		size += calculateDirectorySize(child)
	}
	return size
}
