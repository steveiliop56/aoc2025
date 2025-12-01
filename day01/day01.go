package day01

import (
	"strconv"
	"strings"
)

type Item struct {
	direction int // -1 for left, 1 for right
	steps     int
}

func parseInput(input string) []Item {
	var items []Item

	lines := strings.SplitSeq(input, "\n")

	for line := range lines {
		parts := strings.SplitN(line, "", 2)
		items = append(items, Item{
			direction: func() int {
				if parts[0] == "L" {
					return -1
				}
				return 1
			}(),
			steps: func() int {
				steps, err := strconv.Atoi(parts[1])
				if err != nil {
					panic(err)
				}
				return steps
			}(),
		})
	}

	return items
}
