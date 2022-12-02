package main

import (
	"fmt"
	"os"
	"strings"
)

type Combo struct {
	Loss        string
	Draw       string
	Win       string
}

var (
	combos = map[string]Combo{
		"A": {"Z", "X", "Y"},
		"B": {"X", "Y", "Z"},
		"C": {"Y", "Z", "X"},
	}
)

func main() {
	dat, _ := os.ReadFile("./input.txt")
	input := string(dat)
	partOne := 0
	partTwo := 0

	for _, round := range strings.Split(input, "\n") {
		oppChoice := round[0:1]
		myChoice := round[2:3]

		combo := combos[oppChoice]

		partOne += strings.Index("XYZ", myChoice) + 1

		switch myChoice {
		case combo.Win:
			partOne += 6
		case combo.Draw:
			partOne += 3
		}

		var move string

		switch myChoice {
		case "X":
			move = combo.Loss
		case "Y":
			partTwo += 3
			move = combo.Draw
		case "Z":
			partTwo += 6
			move = combo.Win
		}

		partTwo += strings.Index("XYZ", move) + 1
	}

	fmt.Printf("Part 1: %d\n", partOne)
	fmt.Printf("Part 2: %d\n", partTwo)
}
