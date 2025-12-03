package day03

import (
	"aoc2025/util"
	"fmt"
	"slices"
	"strconv"
)

func SolvePart01(sample bool) {
	fmt.Println("Day 03 - Part 01")
	input := util.LoadInput(3, sample)

	parsed := parseInput(input)
	fmt.Printf("Parsed Input: %v\n", parsed)

	var total int

	for _, row := range parsed {
		fmt.Printf("Processing Row: %v\n", row)

		joltage1 := getLargestJoltagePart01(row, []int{})

		// stupid hack
		if joltage1.Index == len(row)-1 {
			joltage1 = getLargestJoltagePart01(row, []int{joltage1.Index})
		}

		ignoreIndexes := []int{}

		for i := range joltage1.Index + 1 {
			ignoreIndexes = append(ignoreIndexes, i)
		}

		fmt.Printf("	Ignore Indexes for 2nd Joltage: %v\n", ignoreIndexes)

		joltage2 := getLargestJoltagePart01(row, ignoreIndexes)

		fmt.Printf("	Largest Joltage 1: %v\n", joltage1)
		fmt.Printf("	Largest Joltage 2: %v\n", joltage2)

		num := joltageToNumPart01([]Joltage{joltage1, joltage2})
		fmt.Printf("	Combined Number: %d\n", num)

		total += num
	}

	fmt.Printf("Total Sum: %d\n", total)
}

func getLargestJoltagePart01(row []int, ignoreIndexes []int) Joltage {
	var res Joltage
	var prev = 0

	for i, num := range row {
		if num > prev && !slices.Contains(ignoreIndexes, i) {
			res.Value = num
			res.Index = i
			prev = num
		}
	}

	return res
}

func joltageToNumPart01(j []Joltage) int {
	if len(j) != 2 {
		panic("joltageToNumPart01 expects exactly 2 Joltage values")
	}

	parts := make([]int, len(j))

	for i, joltage := range j {
		parts[i] = joltage.Value
	}

	if j[0].Index > j[1].Index {
		slices.Reverse(parts)
	}

	num01 := strconv.Itoa(parts[0])
	num02 := strconv.Itoa(parts[1])

	numStr := num01 + num02
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}

	return num
}
