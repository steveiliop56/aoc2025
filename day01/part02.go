package day01

import (
	"aoc2025/util"
	"fmt"
)

func SolvePart02(sample bool) {
	fmt.Println("Day 01 - Part 02")
	input := util.LoadInput(1, sample)

	items := parseInput(input)
	currentPos := 50
	zeroCount := 0

	for _, item := range items {
		newPos := calculatePosPart01(currentPos, item)
		rotations := calculateRotations(currentPos, newPos, item)
		currentPos = newPos
		if currentPos == 0 {
			zeroCount++
		}
		zeroCount += rotations
		fmt.Printf("Item: %+v, New Position: %d, Rotations: %d, Zeroes: %d\n", item, currentPos, rotations, zeroCount)
	}

	fmt.Printf("Final Position: %d\n", currentPos)
	fmt.Printf("Zero Count: %d\n", zeroCount)
}

func calculateRotations(currentPos int, newPos int, instruction Item) int {
	rotations := instruction.steps / 100
	stepsLeft := instruction.steps % 100

	switch instruction.direction {
	case 1: // Right
		if currentPos+stepsLeft > 99 {
			rotations++
			if newPos == 0 {
				rotations--
			}
		}
	case -1: // Left
		if currentPos-stepsLeft < 0 {
			if currentPos != 0 {
				rotations++
			}
			if newPos == 0 {
				rotations--
			}
		}
	}

	return rotations
}
