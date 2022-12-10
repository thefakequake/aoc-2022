package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	doneBy int
	amount int
}

type PixelMap [6][40]bool

func main() {
	dat, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(dat), "\n")

	var curIns *Instruction
	curLine := 0
	x := 1
	partOne := 0
	partTwo := PixelMap{}

	for c := 1; c <= 240; c++ {
		if (c-20)%40 == 0 {
			partOne += c * x
		}

		curColumn := (c - 1) % 40
		if x-1 <= curColumn && curColumn <= x+1 {
			partTwo[(c - 1) / 40][curColumn] = true
		}

		if curIns == nil && curLine < len(lines) {
			s := strings.Split(lines[curLine], " ")
			curLine++
			if s[0] == "addx" {
				amount, _ := strconv.Atoi(s[1])
				curIns = &Instruction{doneBy: c + 1, amount: amount}
			}
		} else if curIns != nil {
			if curIns.doneBy == c {
				x += curIns.amount
				curIns = nil
			}
		}
	}

	fmt.Printf("Part 1: %d\n", partOne)
	fmt.Printf("Part 2: \n%s\n", renderDisplay(partTwo))
}

func renderDisplay(pixels PixelMap) string {
	output := ""
	for _, row := range pixels {
		for _, pixel := range row {
			if pixel {
				output += "#"
			} else {
				output += "."
			}
		}
		output += "\n"
	}
	return output[:len(output)-1]
}
