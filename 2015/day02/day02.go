package main

import (
	"fmt"
	"github.com/Gavus/advent-of-code/utils/conv"
	"github.com/Gavus/advent-of-code/utils/files"
	"github.com/Gavus/advent-of-code/utils/log"
)

const (
	question1 = "How many total square feet of wrapping paper should they order?"
	question2 = "How many total feet of ribbon should they order?"
)

func main() {
	log.Init()
	input := file.GetInput("2015-12-02", "\n")

	fmt.Println(question1, wrappingPaperNeeded(input))
	fmt.Println(question2, ribbonNeeded(input))
}

func wrappingPaperNeeded(input []string) int {
	feet := 0
	presents := conv.ToPresents(input)

	for _, box := range presents {
		feet += box.SurfaceArea() + box.SmallestSideArea()
	}

	return feet
}

func ribbonNeeded(input []string) int {
	return 0
}
