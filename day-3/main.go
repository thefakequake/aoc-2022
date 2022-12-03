package main

import (
	"fmt"
	"os"
	"strings"
)

const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	dat, _ := os.ReadFile("./input.txt")
	input := string(dat)

	split := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	partOne := 0
	for _, bp := range split {
		comp1 := bp[:len(bp)/2]
		comp2 := bp[len(bp)/2:]
		common := ""
		for _, c := range comp2 {
			char := string(c)
			if !strings.Contains(common, char) && strings.Contains(comp1, char) {
				common += char
				partOne += strings.Index(alpha, char) + 1
			}
		}

	}

	partTwo := 0
	for i := 0; i < len(split); i += 3 {
		group := split[i : i+3]
		max := ""
		for _, bp := range group {
			if len(bp) > len(max) {
				max = bp
			}
		}

		for _, c := range max {
			char := string(c)
			found := true
			for _, bp := range group {
				if !strings.Contains(bp, char) {
					found = false
				}
			}
			if found {
				partTwo += strings.Index(alpha, char) + 1
				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", partOne)
	fmt.Printf("Part 2: %d\n", partTwo)
}
