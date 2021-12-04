package main

import (
	"fmt"
	"github.com/Gavus/advent-of-go/utils/input"
	"github.com/Gavus/advent-of-go/utils/conv"
	_ "github.com/Gavus/advent-of-go/utils/log"
)

const (
	question1 = "What will your final score be if you choose that board?"
	question2 = "unknown"
)

func main() {
	input := input.GetInput("2021-12-04", "\n")

	solution1 := part1(input)
	fmt.Println(question1, solution1)

	solution2 := part2(input)
	fmt.Println(question2, solution2)
}

func part1(input []string) int {
	bi, bbs := conv.ToBingo(input)

	for _, v := range bi {
		for _, bb := range bbs {
			bb.Add(v)
			fmt.Println(bb)
			bingo, _ := bb.Bingo()
			if bingo {
				return bb.SumUnmarked() * v
			}

		}
	}
	return 0
}

func part2(input []string) int {
	return 0
}
