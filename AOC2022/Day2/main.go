package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	// A for rock, 		score: 1
	// B for paper, 	score: 2
	// c for scissor,   score: 3
	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func partOne() int {
	var score int

	for _, value := range parseInput() {
		enemy := value[0]
		self := value[1]

		switch enemy {
		case "A":
			if self == "X" {
				score += 4
			} else if self == "Y" {
				score += 8
			} else {
				score += 3
			}

		case "B":
			if self == "X" {
				score += 1
			} else if self == "Y" {
				score += 5
			} else {
				score += 9
			}

		case "C":
			if self == "X" {
				score += 7
			} else if self == "Y" {
				score += 2
			} else {
				score += 6
			}
		}
	}

	return score
}

func partTwo() int {
	var score int

	for _, value := range parseInput() {
		enemy := value[0]
		condition := value[1]

		switch condition {
		case "X": // Lose condition
			if enemy == "A" {
				score += 3
			} else if enemy == "B" {
				score += 1
			} else {
				score += 2
			}

		case "Y":
			if enemy == "A" {
				score += 4
			} else if enemy == "B" {
				score += 5
			} else {
				score += 6
			}

		case "Z":
			if enemy == "A" {
				score += 8
			} else if enemy == "B" {
				score += 9
			} else {
				score += 7
			}

		}
	}

	return score
}

func parseInput() (ans [][]string) {
	for _, line := range strings.Split(input, "\r\n") {
		word := strings.Fields(line)
		ans = append(ans, word)
	}

	return
}
