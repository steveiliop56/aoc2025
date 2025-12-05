package day05

import (
	"slices"
	"strconv"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

func parseInput(input string) ([]int, []Interval) {
	var avail []int
	var fresh []Interval

	rFresh, rAvail, ok := strings.Cut(input, "\n\n")

	if !ok {
		panic("invalid input")
	}

	for r := range strings.SplitSeq(rAvail, "\n") {
		num, err := strconv.Atoi(r)

		if err != nil {
			panic(err)
		}

		if !slices.Contains(avail, num) {
			avail = append(avail, num)
		}
	}

	for r := range strings.SplitSeq(rFresh, "\n") {
		start, end, ok := strings.Cut(r, "-")
		if !ok {
			panic("invalid range")
		}

		s, err := strconv.Atoi(start)

		if err != nil {
			panic(err)
		}

		e, err := strconv.Atoi(end)

		if err != nil {
			panic(err)
		}

		fresh = append(fresh, Interval{Start: s, End: e})
	}

	return avail, fresh
}
