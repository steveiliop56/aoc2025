package day04

import (
	"aoc2025/util"
	"fmt"
)

func SolvePart01(sample bool) {
	fmt.Println("Day 04 - Part 01")
	input := util.LoadInput(4, sample)

	parsed := parseInput(input)

	var sum int

	for i, line := range parsed {
		for j, r := range line {
			if checkGridPart01(parsed, line, i, j, r) {
				sum++
			}
		}
	}

	fmt.Printf("Paper rolls that can be used: %d\n", sum)
}

func safeLookupPart01(line []rune, index int) bool {
	if index < 0 || index >= len(line) {
		return false
	}
	return line[index] == PaperSpot
}

func checkGridPart01(parsed [][]rune, line []rune, i, j int, r rune) bool {
	spots := 0

	if r != PaperSpot {
		return false
	}

	if safeLookupPart01(line, j-1) {
		spots++
	}
	if safeLookupPart01(line, j+1) {
		spots++
	}

	if i >= 1 {
		aboveLine := parsed[i-1]
		if safeLookupPart01(aboveLine, j) {
			spots++
		}
	}

	if i+1 < len(parsed) {
		belowLine := parsed[i+1]
		if safeLookupPart01(belowLine, j) {
			spots++
		}
	}

	if i >= 1 {
		aboveLine := parsed[i-1]
		if safeLookupPart01(aboveLine, j-1) {
			spots++
		}
		if safeLookupPart01(aboveLine, j+1) {
			spots++
		}
	}

	if i+1 < len(parsed) {
		belowLine := parsed[i+1]
		if safeLookupPart01(belowLine, j-1) {
			spots++
		}
		if safeLookupPart01(belowLine, j+1) {
			spots++
		}
	}

	if spots < 4 {
		return true
	}

	return false
}
