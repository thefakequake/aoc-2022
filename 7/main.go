package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Size     int
	Name     string
	Children []*Node
}

func main() {
	dat, _ := os.ReadFile("./input.txt")
	inp := string(dat)
	
	var currentDir []string
	var readingFiles bool
	tree := &Node{}

	for _, line := range strings.Split(inp, "\n") {
		split := strings.Split(line, " ")
		if split[0] == "$" {
			readingFiles = false
			if split[1] == "cd" {
				if split[2] == "/" {
					currentDir = []string{}
				} else if split[2] == ".." {
					currentDir = currentDir[:len(currentDir)-1]
				} else {
					currentDir = append(currentDir, split[2])
				}
			} else {
				readingFiles = true
			}
		} else if readingFiles {
			node := tree
			fileSize := 0
			if len(currentDir) > 0 {
				node = findInTree(tree, currentDir)
			}
			if split[0] != "dir" {
				fileSize, _ = strconv.Atoi(split[0])
			}
			node.Children = append(node.Children, &Node{Name: split[1], Size: fileSize})
		}
	}

	partOne := 0
	partTwo := 0
	sums, _ := sumTree(tree)
	unused := 70000000 - sums[len(sums)-1]

	for _, sum := range sums {
		if sum <= 100000 {
			partOne += sum
		}
		if (unused + sum >= 30000000) && (partTwo == 0 || sum < partTwo)  {
			partTwo = sum
		}
	}

	fmt.Printf("Part 1: %d\n", partOne)
	fmt.Printf("Part 2: %d\n", partTwo)
}

func findInTree(tree *Node, path []string) *Node {
	for _, child := range tree.Children {
		if child.Name == path[0] {
			if len(path) == 1 {
				return child
			}
			return findInTree(child, path[1:])
		}
	}
	return nil
}

func sumTree(tree *Node) ([]int, int) {
	var currentTotal int
	dirTotals := []int{}
	for _, child := range tree.Children {
		currentTotal += child.Size
		if len(child.Children) > 0 {
			dirs, total := sumTree(child)
			currentTotal += total
			dirTotals = append(dirTotals, dirs...)
		}
	}
	if currentTotal > 0 {
		dirTotals = append(dirTotals, currentTotal)
	}

	return dirTotals, currentTotal
}
