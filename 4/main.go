package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("input.txt")
	inp := string(dat)

	partOne := 0
	partTwo := 0

	for _, line := range strings.Split(inp, "\n") {
		var nums []int
		for _, value := range strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, "-", " "), ",", " "), " ") {
			n, _ := strconv.Atoi(value)
			nums = append(nums, n)
		}

		if nums[0] <= nums[2] && nums[1] >= nums[3] || nums[2] <= nums[0] && nums[3] >= nums[1] {
			partOne += 1
		}
		if !(nums[1] < nums[2] || nums[0] > nums[3]) {
			partTwo += 1
		}
	}

	fmt.Printf("Part 1: %d\n", partOne)
	fmt.Printf("Part 2: %d\n", partTwo)
}
