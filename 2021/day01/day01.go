package main

import (
	"fmt"
	"github.com/Gavus/advent-of-go/utils/input"
	"github.com/Gavus/advent-of-go/utils/types/ints"
	"github.com/Gavus/advent-of-go/utils/types/queue_ring"
	_ "github.com/Gavus/advent-of-go/utils/log"
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
	ints, _ := ints.ToInts(input)
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
	ints, _ := ints.ToInts(input)
	a := queue_ring.Make(3)
	b := queue_ring.Make(3)
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
