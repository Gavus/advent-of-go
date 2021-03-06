package main

import (
	"fmt"
	"github.com/Gavus/advent-of-go/utils/input"
	_ "github.com/Gavus/advent-of-go/utils/log"
)

const (
	whichFloorDoSantaGo = "To what floor do the instructions take Santa?"
	question2           = "What is the position of the character that causes Santa to first enter the basement?"
)

func main() {
	input := input.GetInput("2015-12-01", "")

	solution1, _ := part1(input, 0)
	fmt.Println(whichFloorDoSantaGo, solution1)

	solution2 := part2(input)
	fmt.Println(question2, solution2)
}

func part1(input []string, floorStop int) (int, int) {
	floor := 0
	index := 0

	for i, val := range input {
		if val == "(" {
			floor++
		} else if val == ")" {
			floor--
		}
		if floorStop != 0 && floor == floorStop {
			index = i
			break
		}
	}

	return floor, index
}

func part2(input []string) int {
	floor := 0
	floorstop := -1
	index := 1

	for {
		floorTmp, indexTmp := part1(input, floorstop)
		floor += floorTmp
		index += indexTmp

		if floor == floorstop {
			break
		}
	}

	return index
}
