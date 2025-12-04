package day04

import (
	"fmt"
	"strings"
)

const BlankSpot = rune('.')
const PaperSpot = rune('@')

func parseInput(input string) [][]rune {
	var result [][]rune

	lines := strings.SplitSeq(input, "\n")

	for line := range lines {
		runes := []rune(line)
		result = append(result, runes)
	}

	return result
}

func recompileGrid(parsed [][]rune) {
	for _, line := range parsed {
		var sb strings.Builder

		for _, r := range line {
			sb.WriteRune(r)
		}

		fmt.Println(sb.String())
	}
}
