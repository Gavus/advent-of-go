package main

import (
	"fmt"
	"github.com/Gavus/advent-of-code/utils/conv"
	"github.com/Gavus/advent-of-code/utils/input"
	"github.com/Gavus/advent-of-code/utils/types"
	_ "github.com/Gavus/advent-of-code/utils/log"
)

const (
	question1 = "How many measurements are larger than the previous measurement?"
	question2 = "How many sums are larger than the previous sum?"
)

func main() {
	input := input.GetInput("2021-12-01", "\n")

	solution1 := part1(input)
	fmt.Println(question1, solution1)

	solution2 := part2(input)
	fmt.Println(question2, solution2)
}

func part1(input []string) int {
	increased := 0
	ints, _ := conv.ToInts(input)
	for i, v := range ints {
		if i == 0 {
			continue
		}

		if ints[i-1] < v {
			increased++
		}
	}

	return increased
}

func part2(input []string) int {
	increased := 0
	ints, _ := conv.ToInts(input)
	a := types.MakeRing(3)
	b := types.MakeRing(3)
	for i, v := range ints {
		if i == 0 {
			a.Put(v)
			continue
		}
		b.Put(v)
		if a.Sum() < b.Sum() && i > 2{
			increased++
		}
		a.Put(v)
	}
	return increased
}
