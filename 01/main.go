package main

import (
	"os"
	"strings"
	"fmt"
	"strconv"
	"sort"
)

func main() {
	dat, _ := os.ReadFile("./input.txt")
	input := string(dat)
	elves := []int{}

	for _, elf := range strings.Split(strings.ReplaceAll(input, "\n\r", "\n"), "\n\n") {
		calories := strings.Split(elf, "\n")
		total := 0
		for _, amount := range calories {
			num, _ := strconv.Atoi(strings.TrimSpace(amount))
			total += num
		}
		elves = append(elves, total)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	
	fmt.Printf("Part 1: %d\n", elves[0])
	fmt.Printf("Part 2: %d\n", elves[0] + elves[1] + elves[2])
}
