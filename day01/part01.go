package day01

import (
	"aoc2025/util"
	"fmt"
)

func SolvePart01(sample bool) {
	fmt.Println("Day 01 - Part 01")
	input := util.LoadInput(1, sample)

	items := parseInput(input)
	currentPos := 50
	zeroCount := 0

	for _, item := range items {
		newPos := calculatePosPart01(currentPos, item)
		currentPos = newPos
		if currentPos == 0 {
			zeroCount++
		}
		fmt.Printf("Item: %+v, New Position: %d\n", item, currentPos)
	}

	fmt.Printf("Final Position: %d\n", currentPos)
	fmt.Printf("Zero Count: %d\n", zeroCount)
}

func calculatePosPart01(currentPos int, instruction Item) int {
	stepsRemaining := instruction.steps % 100

	newPos := 0

	switch instruction.direction {
	case 1: // Right
		newPos = currentPos + stepsRemaining
		if newPos > 99 {
			newPos = newPos - 100
		}
	case -1: // Left
		newPos = currentPos - stepsRemaining
		if newPos < 0 {
			newPos = 100 - (-newPos)
		}
	}

	return newPos
}
