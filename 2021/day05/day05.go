package main

import (
	"fmt"
	"github.com/Gavus/advent-of-go/utils/input"
	_ "github.com/Gavus/advent-of-go/utils/log"
	"github.com/Gavus/advent-of-go/utils/types/line"
)

const (
	question1 = "At how many points do at least two lines overlap?"
	question2 = "unknown"
)

func main() {
	input := input.GetInput("2021-12-05", "\n")

	solution1 := part1(input)
	fmt.Println(question1, solution1)

	solution2 := part2(input)
	fmt.Println(question2, solution2)
}

func part1(input []string) int {
	lines := line.StringsToLines(input)
	lines.StraightToPoints()
	return 0
}

func part2(input []string) int {
	return 0
}
