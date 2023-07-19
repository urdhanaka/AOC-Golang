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

	m := make(map[int]int)
	r := make(map[int]int)

	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
	}

	return res
}
