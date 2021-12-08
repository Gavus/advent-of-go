package main

import (
	"fmt"
	"github.com/Gavus/advent-of-go/utils/input"
	"github.com/Gavus/advent-of-go/utils/types/bingo"
	_ "github.com/Gavus/advent-of-go/utils/log"
)

const (
	question1 = "What will your final score be if you choose that board?"
	question2 = "Once it wins, what would its final score be?"
)

func main() {
	input := input.GetInput("2021-12-04", "\n")

	solution1 := part1(input)
	fmt.Println(question1, solution1)

	solution2 := part2(input)
	fmt.Println(question2, solution2)
}

func part1(input []string) int {
	bi, bbs := bingo.ToBingo(input)

	for _, v := range bi {
		for i := range bbs {
			bbs[i].Add(v)
			if bbs[i].Bingo() {
				return bbs[i].Score()
			}
		}
	}
	return 0
}

func part2(input []string) int {
	bi, bbs := bingo.ToBingo(input)

	for _, v := range bi {
		bingoCount := 0
		for _, bb := range bbs {
			if bb.Bingo() {
				bingoCount++
			}
		}

		for i := range bbs {
			before := bbs[i].Bingo()
			bbs[i].Add(v)
			if bbs[i].Bingo() {
				if !before && bingoCount == len(bbs)-1 {
					return bbs[i].Score()
				}
			}
		}
	}
	return 0
}
