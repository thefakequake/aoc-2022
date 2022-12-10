package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("./input.txt")
	inp := string(dat)

	var trees [][]int
	for _, line := range strings.Split(inp, "\n") {
		var curLine []int
		for _, c := range line {
			num, _ := strconv.Atoi(string(c))
			curLine = append(curLine, num)
		}
		trees = append(trees, curLine)
	}

	partOne := 0
	partTwo := 0

	for i, row := range trees {
		for j, tree := range row {
			right := row[j+1:]
			left := []int{}
			for _, n := range row[:j] {
				left = append([]int{n}, left...)
			}
			up := []int{}
			for k := 0; k < i; k++ {
				up = append([]int{trees[k][j]}, up...)
			}
			down := []int{}
			for k := i + 1; k < len(trees); k++ {
				down = append(down, trees[k][j])
			}

			if (i == 0 || i == len(trees)-1 || j == 0 || j == len(row)-1) ||
				(allLess(tree, left) || allLess(tree, right) || allLess(tree, up) || allLess(tree, down)) {
				partOne += 1
			}

			scenicScore := calcViewDist(tree, left) * calcViewDist(tree, right) * calcViewDist(tree, up) * calcViewDist(tree, down)
			if scenicScore > partTwo {
				partTwo = scenicScore
			}
		}
	}

	fmt.Printf("Part 1: %d\n", partOne)
	fmt.Printf("Part 2: %d\n", partTwo)
}

func allLess(target int, arr []int) bool {
	for _, n := range arr {
		if n >= target {
			return false
		}
	}
	return true
}

func calcViewDist(target int, arr []int) int {
	for i, n := range arr {
		if n >= target || i == len(arr)-1 {
			return i + 1
		}
	}
	return 0
}
