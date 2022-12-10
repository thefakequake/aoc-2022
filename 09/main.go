package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

func main() {
	dat, _ := os.ReadFile("./input.txt")
	cmds := strings.Split(string(dat), "\n")

	fmt.Printf("Part 1: %d\n", simRope(1, cmds))
	fmt.Printf("Part 2: %d\n", simRope(9, cmds))
}

func simRope(numKnots int, cmds []string) int {
	knots := []*Position{}
	for i := 0; i < numKnots; i++ {
		knots = append(knots, &Position{0, 0})
	}

	headPos := &Position{}
	unqTailPos := []Position{{0, 0}}

	for _, cmd := range cmds {
		s := strings.Split(cmd, " ")
		direction := s[0]
		amount, _ := strconv.Atoi(s[1])

	MoveRope:
		for i := 0; i < amount; i++ {
			switch direction {
			case "R":
				headPos.X++
			case "L":
				headPos.X--
			case "U":
				headPos.Y++
			case "D":
				headPos.Y--
			}

			for i, knotPos := range knots {
				lastKnot := headPos
				if i > 0 {
					lastKnot = knots[i-1]
				}

				xDiff := float64(lastKnot.X - knotPos.X)
				yDiff := float64(lastKnot.Y - knotPos.Y)

				if math.Abs(xDiff) < 2 && math.Abs(yDiff) < 2 {
					continue
				}

				if xDiff > 0 {
					knotPos.X++
				} else if xDiff < 0 {
					knotPos.X--
				}
				if yDiff > 0 {
					knotPos.Y++
				} else if yDiff < 0 {
					knotPos.Y--
				}

				if i+1 == numKnots {
					for _, pos := range unqTailPos {
						if pos.X == knotPos.X && pos.Y == knotPos.Y {
							continue MoveRope
						}
					}
					unqTailPos = append(unqTailPos, *knotPos)
				}
			}
		}
	}

	return len(unqTailPos)
}
