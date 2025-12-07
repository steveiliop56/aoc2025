package day06

import (
	"aoc2025/util"
	"fmt"
	"strconv"
	"strings"
)

func SolvePart01(sample bool) {
	fmt.Println("Day 06 - Part 01")
	input := util.LoadInput(6, sample)

	numCols, operations := parseInputPart01(input)

	var sum int

	for i, col := range numCols {
		op := operations[i]
		colSum := 0
		for _, num := range col {
			switch op {
			case 1:
				colSum += num
			case 2:
				if colSum == 0 {
					colSum = 1
				}
				colSum *= num
			}
		}
		sum += colSum
		fmt.Printf("Column %d, Operation %d, Column Sum: %d\n", i, op, colSum)
	}

	fmt.Println("Result:", sum)
}

func parseInputPart01(input string) ([][]int, []int) {
	var numbers [][]int
	var operations []int

	lines := strings.Split(input, "\n")
	operationsLn := lines[len(lines)-1]
	numbersLns := lines[:len(lines)-1]

	operationsSeq := strings.SplitSeq(operationsLn, " ")

	for operation := range operationsSeq {
		opStr := strings.TrimSpace(operation)
		if opStr == "" {
			continue
		}
		switch opStr {
		case "+":
			operations = append(operations, 1)
		case "*":
			operations = append(operations, 2)
		default:
			panic("unknown operation: " + opStr)
		}
	}

	var numRows [][]int

	for _, line := range numbersLns {
		numSeq := strings.SplitSeq(line, " ")
		var row []int
		for num := range numSeq {
			numStr := strings.TrimSpace(num)
			if numStr == "" {
				continue
			}
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic("invalid number: " + numStr)
			}
			row = append(row, num)
		}
		numRows = append(numRows, row)
	}

	for i := range len(numRows[0]) {
		var col []int
		for j := range numRows {
			col = append(col, numRows[j][i])
		}
		numbers = append(numbers, col)
	}

	return numbers, operations
}
