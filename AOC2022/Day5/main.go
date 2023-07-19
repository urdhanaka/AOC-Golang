package main

import (
	_ "embed"

	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(partOne())
}

func partOne() int {
	var res int
	crane := [][]rune{
		{'c'}, // empty array for index 0
		{'m', 'j', 'c', 'b', 'f', 'r', 'l', 'h'},
		{'z', 'c', 'd'},
		{'h', 'j', 'f', 'c', 'n', 'g', 'w'},
		{'p', 'j', 'd', 'm', 't', 's', 'b'},
		{'n', 'c', 'd', 'r', 'j'},
		{'w', 'l', 'd', 'q', 'p', 'j', 'g', 'z'},
		{'p', 'z', 't', 'f', 'r', 'h'},
		{'l', 'v', 'm', 'g'},
		{'c', 'b', 'g', 'p', 'f', 'q', 'r', 'j'},
	}

	for i := 1; i < len(crane); i++ {
		fmt.Println(crane[i][len(crane[i])-1])
	}

	return res
}

func swap(crane *[][]rune, iter int, start int, end int) {
	for i := 1; i <= iter; i++ {
		if len((*crane)[start]) == 0 {
			break
		}

		(*crane)[end] = append((*crane)[end], (*crane)[start][len((*crane)[start])-1])
	}
}
