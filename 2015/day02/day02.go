package main

import (
	"fmt"
	"github.com/Gavus/advent-of-go/utils/conv"
	"github.com/Gavus/advent-of-go/utils/input"
	_ "github.com/Gavus/advent-of-go/utils/log"
)

const (
	question1 = "How many total square feet of wrapping paper should they order?"
	question2 = "How many total feet of ribbon should they order?"
)

func main() {
	input := input.GetInput("2015-12-02", "\n")

	fmt.Println(question1, part1(input))
	fmt.Println(question2, part2(input))
}

func part1(input []string) int {
	feet := 0
	presents, _ := conv.ToPresents(input)

	for _, present := range presents {
		feet += present.SurfaceArea() + present.SmallestSideArea()
	}

	return feet
}

func part2(input []string) int {
	feet := 0
	presents, _ := conv.ToPresents(input)

	for _, present := range presents {
		feet += present.RibbonLength()
	}

	return feet
}
