package day02

import (
	"aoc2025/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolvePart02(sample bool) {
	fmt.Println("Day 02 - Part 02")
	input := util.LoadInput(2, sample)

	parsed := parseInput(input)

	invalidIDs := getInvalidIDsPart02(parsed)

	sum := 0

	for _, id := range invalidIDs {
		sum += id
	}

	fmt.Printf("Invalid IDs: %v\n", invalidIDs)
	fmt.Printf("Sum of invalid IDs: %d\n", sum)
}

func getInvalidIDsPart02(ids [][]int) []int {
	var invalidIDs []int

	for _, idRange := range ids {
		fmt.Printf("Checking ID range: %d-%d\n", idRange[0], idRange[len(idRange)-1])
		for _, num := range idRange {
			idStr := strconv.Itoa(num)

			if len(idStr) == 1 {
				continue
			}

			if len(idStr) == 2 {
				if idStr[0] == idStr[1] {
					invalidIDs = append(invalidIDs, num)
					fmt.Printf("Found invalid ID: %d\n", num)
				}
				continue
			}

			if len(idStr) == 3 {
				if idStr[0] == idStr[1] && idStr[1] == idStr[2] {
					invalidIDs = append(invalidIDs, num)
					fmt.Printf("Found invalid ID: %d\n", num)
				}
				continue
			}

			parts := strings.Split(idStr, "")
			slices.Reverse(parts)

			mid := len(parts) / 2
			left := len(parts) % 2

			leftPart := parts[:mid+left]

			slices.Reverse(leftPart)

			sub := strings.Join(leftPart, "")

			if strings.ReplaceAll(idStr, sub, "") == "" {
				invalidIDs = append(invalidIDs, num)
				fmt.Printf("Found invalid ID: %d\n", num)
				continue
			}

			for range len(leftPart) {
				subParts := strings.Split(sub, "")
				subParts = slices.Delete(subParts, 0, 1)
				sub = strings.Join(subParts, "")
				if strings.ReplaceAll(idStr, sub, "") == "" {
					invalidIDs = append(invalidIDs, num)
					fmt.Printf("Found invalid ID: %d\n", num)
					break
				}
			}
		}
	}

	return invalidIDs
}
