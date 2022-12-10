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

	split := strings.Split(strings.ReplaceAll(inp, "\r\n", "\n"), "\n")
	
	var startStacks [9][]string
	var commands []string

	for index, line := range split {
		if split[index+1] == "" {
			commands = split[index+2:]
			break
		}

		for i, j := 0, 1; j < len(line); i, j = i+1, j+4 {
			char := string(line[j])
			if strings.TrimSpace(char) != "" {
				startStacks[i] = append([]string{char}, startStacks[i]...)
			}
		}
	}

	for part := 0; part < 2; part++ {
		stacks := startStacks

		for _, cmd := range commands {
			parts := strings.Split(cmd, " ")
			amount, _ := strconv.Atoi(parts[1])
			from, _ := strconv.Atoi(parts[3])
			to, _ := strconv.Atoi(parts[5])
			from--
			to--

			fromStack := stacks[from]
			sliceIndex := len(fromStack) - amount
			crates := fromStack[sliceIndex:]
			stacks[from] = fromStack[:sliceIndex]

			if part == 0 {
				for i, j := 0, len(crates)-1; i < j; i, j = i+1, j-1 {
					crates[i], crates[j] = crates[j], crates[i]
				}
			}

			stacks[to] = append(stacks[to], crates...)
		}

		fmt.Printf("Part %d: ", part+1)
		for _, stack := range stacks {
			fmt.Print(stack[len(stack)-1])
		}
		fmt.Println()
	}
}
