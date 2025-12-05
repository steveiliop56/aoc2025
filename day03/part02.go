package day03

import (
	"aoc2025/util"
	"fmt"
	"strconv"
	"strings"
)

// Hey, just wanted to point out that in order to figure out the logic, I had to consult an LLM. Code is mine, logic not so much ;)

func SolvePart02(sample bool) {
	fmt.Println("Day 03 - Part 02")
	input := util.LoadInput(3, sample)

	parsed := parseInput(input)
	fmt.Printf("Parsed Input: %v\n", parsed)

	var total int

	for _, row := range parsed {
		fmt.Printf("Processing Row: %v\n", row)

		largest := getLargestJoltagePart02(row, 12)

		fmt.Printf("	Largest 12 Joltages: %v\n", largest)

		num := joltageToNumPart02(largest)

		fmt.Printf("	Combined Number: %d\n", num)

		total += num
	}

	fmt.Printf("Total Sum: %d\n", total)
}

func getLargestJoltagePart02(row []int, length int) []int {
	res := []int{}
	removals := len(row) - length

	for i, num := range row {
		if i == 0 {
			res = append(res, num)
			continue
		}
		if num > res[len(res)-1] && removals > 0 && len(res) > 0 {
			for range len(res) {
				if removals <= 0 {
					break
				}
				if num <= res[len(res)-1] {
					break
				}
				res = res[:len(res)-1]
				removals--
			}
			res = append(res, num)
			continue
		}
		res = append(res, num)
	}

	res = res[:length]

	return res
}

func joltageToNumPart02(row []int) int {
	rowStr := make([]string, 12)

	for _, num := range row {
		rowStr = append(rowStr, strconv.Itoa(num))
	}

	num, err := strconv.Atoi(strings.Join(rowStr, ""))

	if err != nil {
		panic(err)
	}

	return num
}
