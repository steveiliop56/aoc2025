package day05

import (
	"aoc2025/util"
	"fmt"
)

// Hey, just wanted to point out that in order to figure out the logic, I had to consult an LLM. Code is mine, logic not so much ;)

func SolvePart01(sample bool) {
	fmt.Println("Day 05 - Part 01")
	input := util.LoadInput(5, sample)

	avail, fresh := parseInput(input)

	count := 0

	for _, num := range avail {
		for _, interval := range fresh {
			if num >= interval.Start && num <= interval.End {
				count++
				break
			}
		}
	}

	fmt.Printf("Total fresh numbers: %d\n", count)
}
