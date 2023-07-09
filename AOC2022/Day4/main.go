package main

import (
	_ "embed"
	"fmt"
	"strings"
	"util"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(partOne())
}

func partOne() int {
	var res int

	for index := range parseInput() {
		start1 := util.ToInt(parseInput()[index][0])
		end1 := util.ToInt(parseInput()[index][1])
		start2 := util.ToInt(parseInput()[index][2])
		end2 := util.ToInt(parseInput()[index][3])

		len1 := end1 - start1
		len2 := end2 - start2

		if len1 > len2 {
			if (start2 <= end1 && start2 >= start1) && (end2 <= end1 && end2 >= start1) {
				res++
			}
		} else {
			if (start1 <= end2 && start1 >= start2) && (end1 <= end2 && end1 >= start2) {
				res++
			}
		}
	}

	return res
}

func parseInput() (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		row := []string{}

		for _, assignments := range strings.Split(line, ",") {
			for _, temp := range strings.Split(assignments, "-") {
				row = append(row, temp)
			}
		}

		ans = append(ans, row)
	}

	return ans[:len(ans)-1]
}
