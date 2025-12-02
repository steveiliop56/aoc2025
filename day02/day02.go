package day02

import (
	"strconv"
	"strings"
)

func parseInput(input string) [][]int {
	var result [][]int

	parts := strings.SplitSeq(input, ",")

	for part := range parts {
		start, end, ok := strings.Cut(part, "-")
		if !ok {
			panic("invalid input")
		}
		startNum, err := strconv.Atoi(start)
		if err != nil {
			panic(err)
		}
		endNum, err := strconv.Atoi(end)
		if err != nil {
			panic(err)
		}
		numRange := []int{startNum}
		for i := startNum + 1; i <= endNum; i++ {
			numRange = append(numRange, i)
		}
		result = append(result, numRange)
	}

	return result
}
