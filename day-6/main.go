package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("./input.txt")
	inp := string(dat)

	var partOne int
	var partTwo int

	for i := 0; i < len(inp); i++ {
		if i >= 3 && partOne == 0 && checkUniqueChars(inp[i-3:i+1]) {
			partOne = i + 1
		}
		if i >= 13 && partTwo == 0 && checkUniqueChars(inp[i-13:i+1]) {
			partTwo = i + 1
		}
	}

	fmt.Printf("Part 1: %d\n", partOne)
	fmt.Printf("Part 2: %d\n", partTwo)
}

func checkUniqueChars(str string) bool {
	unique := ""
	for _, c := range str {
		char := string(c)
		if strings.Contains(unique, char) {
			return false
		}
		unique += char
	}
	return true
}
