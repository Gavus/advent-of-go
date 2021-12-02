package main

import (
	"fmt"
	"github.com/Gavus/advent-of-go/utils/conv"
	"github.com/Gavus/advent-of-go/utils/input"
	_ "github.com/Gavus/advent-of-go/utils/log"
	"github.com/Gavus/advent-of-go/utils/types"
)

const (
	question1 = "What do you get if you multiply your final horizontal position by your final depth?"
	question2 = "What do you get if you multiply your final horizontal position by your final depth?"
)

func main() {
	input := input.GetInput("2021-12-02", "\n")

	solution1 := part1(input)
	fmt.Println(question1, solution1)

	solution2 := part2(input)
	fmt.Println(question2, solution2)
}

func part1(input []string) int {
	submarine := types.MakeSubmarine(types.MakePoint(0, 0))
	instructions := conv.ToSubmarineInstructions(input)

	for _, instr := range instructions {
		submarine.Move(instr)
	}

	return submarine.Mult()
}

func part2(input []string) int {
	submarine := types.MakeSubmarine(types.MakePoint(0, 0))
	instructions := conv.ToSubmarineInstructions(input)

	for _, instr := range instructions {
		submarine.Move2(instr)
	}

	return submarine.Mult()
}
