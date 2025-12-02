package day02

import (
	"aoc2025/util"
	"fmt"
	"strconv"
)

func SolvePart01(sample bool) {
	fmt.Println("Day 02 - Part 01")
	input := util.LoadInput(2, sample)

	parsed := parseInput(input)
	invalidIDs := getInvalidIDsPart01(parsed)

	sum := 0

	for _, id := range invalidIDs {
		sum += id
	}

	fmt.Printf("Invalid IDs: %v\n", invalidIDs)
	fmt.Printf("Sum of invalid IDs: %d\n", sum)
}

func getInvalidIDsPart01(ids [][]int) []int {
	var invalidIDs []int

	for _, idRange := range ids {
		fmt.Printf("Checking ID range: %d-%d\n", idRange[0], idRange[len(idRange)-1])
		for _, num := range idRange {
			idStr := strconv.Itoa(num)

			if len(idStr)%2 != 0 {
				continue
			}

			mid := len(idStr) / 2
			firstHalf := idStr[:mid]
			secondHalf := idStr[mid:]

			if firstHalf == secondHalf {
				invalidIDs = append(invalidIDs, num)
				fmt.Printf("Found invalid ID: %d\n", num)
			}
		}
	}

	return invalidIDs
}
