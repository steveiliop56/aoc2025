package day03

import (
	"strconv"
	"strings"
)

type Joltage struct {
	Value  int
	Index  int
	Weight int
}

func parseInput(input string) [][]int {
	var joltages [][]int
	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")

	for line := range lines {
		var row []int
		parts := strings.SplitSeq(line, "")
		for part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}
		joltages = append(joltages, row)
	}

	return joltages
}
