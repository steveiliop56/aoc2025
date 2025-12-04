package day04

import (
	"aoc2025/util"
	"fmt"
)

func SolvePart02(sample bool) {
	fmt.Println("Day 04 - Part 01")
	input := util.LoadInput(4, sample)

	parsed := parseInput(input)

	var sum int
	foundNew := true

	for foundNew {
		totalFound := 0
		for i, line := range parsed {
			purgeIDs := make([]int, 0)
			for j, r := range line {
				if checkGridPart01(parsed, line, i, j, r) {
					sum++
					totalFound++
					purgeIDs = append(purgeIDs, j)
					foundNew = true
				}
			}
			for _, purgeID := range purgeIDs {
				parsed[i][purgeID] = BlankSpot
			}
		}
		if totalFound == 0 {
			foundNew = false
		}
	}

	recompileGrid(parsed) // beautiful artwork
	fmt.Printf("Paper rolls that can be used: %d\n", sum)
}
