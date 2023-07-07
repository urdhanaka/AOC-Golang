package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func partOne() int {
	var res int

	for _, line := range parseInput() {
		m := make(map[rune]bool)

		sackOne := line[:len(line)/2]
		sackTwo := line[len(line)/2:]

		for _, char := range sackOne {
			m[char] = true
		}

		for _, char := range sackTwo {
			tmp := m[char]

			if tmp == true {
				if char <= 'z' && char >= 'a' {
					res += int(char) - 96
				} else if char <= 'Z' && char >= 'A' {
					res += int(char) - 64 + 26
				}
				break
			}
		}
	}

	return res
}

func partTwo() int {
	var res int

	for i := 0; i < len(parseInput())/3; i++ {
		m := make(map[rune]int)

		first := parseInput()[3*i]
		second := parseInput()[3*i+1]
		third := parseInput()[3*i+2]

		for _, char := range first {
			if m[char]+1 == 1 {
				m[char] += 1
			}
		}

		for _, char := range second {
			if m[char]+1 == 2 {
				m[char] += 1
			}
		}

		for _, char := range third {
			if m[char]+1 == 3 {
				if char <= 'z' && char >= 'a' {
					res += int(char) - 96
				} else if char <= 'Z' && char >= 'A' {
					res += int(char) - 64 + 26
				}
				break
			}
		}
	}

	return res
}

func parseInput() (ans []string) {
	for _, line := range strings.Split(input, "\r\n") {
		ans = append(ans, line)
	}

	return
}
