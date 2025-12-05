package day05

import (
	"aoc2025/util"
	"fmt"
	"slices"
)

// Hey, just wanted to point out that in order to figure out the logic, I had to consult an LLM. Code is mine, logic not so much ;)

func SolvePart02(sample bool) {
	fmt.Println("Day 05 - Part 02")
	input := util.LoadInput(5, sample)

	_, fresh := parseInput(input)

	slices.SortFunc(fresh, func(a, b Interval) int {
		return a.Start - b.Start
	})

	merged := []Interval{}

	for _, fr := range fresh {
		if len(merged) == 0 {
			merged = append(merged, fr)
			continue
		}
		if fr.Start <= merged[len(merged)-1].End+1 {
			fmt.Printf("Merging %v into %v\n", fr, merged[len(merged)-1])
			if fr.End > merged[len(merged)-1].End {
				merged[len(merged)-1].End = fr.End
			}
		} else {
			merged = append(merged, fr)
		}
	}

	total := 0

	for _, interval := range merged {
		total += (interval.End - interval.Start) + 1
	}

	fmt.Printf("Fresh %v\n", merged)

	fmt.Printf("Total fresh ranges: %d\n", total)
}
