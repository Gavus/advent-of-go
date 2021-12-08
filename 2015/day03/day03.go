package main

import (
	"fmt"
	"github.com/Gavus/advent-of-go/utils/input"
	"github.com/Gavus/advent-of-go/utils/log"
	"github.com/Gavus/advent-of-go/utils/types/direction"
	"github.com/Gavus/advent-of-go/utils/types/point"
)

const (
	question1 = "How many houses receive at least one present?"
	question2 = "This year, how many houses receive at least one present?"
)

func main() {
	input := input.GetInput("2015-12-03", "")

	solution1 := part1(input)
	fmt.Println(question1, solution1)

	solution2 := part2(input)
	fmt.Println(question2, solution2)
}

func part1(input []string) int {
	dirs, err := direction.ToDirections(input)
	visitedPlaces := point.Points{}
	santaPos := point.Make(0, 0)
	if err != nil {
		log.Err.Fatal(err)
	}

	for _, d := range dirs {
		tmp := point.Move(santaPos, d)
		visitedPlaces = append(visitedPlaces, tmp)
		santaPos = tmp
	}

	visitedPlaces = point.Unique(visitedPlaces)

	return len(visitedPlaces)
}

func part2(input []string) int {
	santa := point.Make(0, 0)
	roboSanta := santa
	dirs, err := direction.ToDirections(input)
	visitedPlaces := point.Points{}
	if err != nil {
		log.Err.Fatal(err)
	}

	for i, d := range dirs {
		if i%2 == 0 {
			tmp := point.Move(roboSanta, d)
			visitedPlaces = append(visitedPlaces, tmp)
			roboSanta = tmp
		} else {
			tmp := point.Move(santa, d)
			visitedPlaces = append(visitedPlaces, tmp)
			santa = tmp
		}
	}

	visitedPlaces = point.Unique(visitedPlaces)

	return len(visitedPlaces)
}
