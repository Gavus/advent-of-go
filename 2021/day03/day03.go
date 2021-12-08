package main

import (
	"fmt"
	"github.com/Gavus/advent-of-go/utils/input"
	_ "github.com/Gavus/advent-of-go/utils/log"
	"github.com/Gavus/advent-of-go/utils/types/binary_strings"
)

const (
	question1 = "What is the power consumption of the submarine?"
	question2 = "What is the life support rating of the submarine?"
)

func main() {
	input := input.GetInput("2021-12-03", "\n")

	solution1 := part1(input)
	fmt.Println(question1, solution1)

	solution2 := part2(input)
	fmt.Println(question2, solution2)
}

func part1(input []string) uint {
	gamma, epsilon := bstrings.CalcGammaEpsilon(input)

	return gamma * epsilon
}

func part2(input []string) uint {
	ogr := bstrings.CalcOxygenGeneratorRating(input)
	co2sr := bstrings.CO2ScrubberRating(input)
	return ogr * co2sr
}
