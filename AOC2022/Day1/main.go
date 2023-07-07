package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"util"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func partOne() int {
	calories := parseInput(input)
	totals := []int{}

	for _, items := range calories {
		totals = append(totals, util.SumIntArrays(items))
	}

	return util.MaxInt(totals...)
}

func partTwo() int {
	var totalThree int

	calories := parseInput(input)
	totals := []int{}

	for _, items := range calories {
		totals = append(totals, util.SumIntArrays(items))
	}

	sort.SliceStable(totals, func(i, j int) bool { return totals[i] > totals[j] })

	for i := 0; i < 3; i++ {
		totalThree += totals[i]
	}

	return totalThree
}

func parseInput(input string) (ans [][]int) {
	for _, group := range strings.Split(input, "\r\n\r\n") {
		row := []int{}

		for _, line := range strings.Split(group, "\r\n") {
			row = append(row, util.ToInt(line))
		}

		ans = append(ans, row)
	}

	return
}
