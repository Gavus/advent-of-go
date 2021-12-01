package main

import (
	"fmt"
	"github.com/Gavus/advent-of-code/utils/input"
	"github.com/Gavus/advent-of-code/utils/conv"
	_ "github.com/Gavus/advent-of-code/utils/log"
	"strings"
)

const (
	question1 = "unknown"
	question2 = "unknown"
)

func main() {
	input := input.GetInput("2015-12-04", "")

	solution1 := part1(input)
	fmt.Println(question1, solution1)

	solution2 := part2(input)
	fmt.Println(question2, solution2)
}

func part1(input []string) string {
	inputStr := strings.Join(input, "")
	for i := 0;; i++ {
		md5 := conv.Md5Sum(inputStr, i)
		fmt.Println(md5)
		if strings.HasPrefix(md5, "00001") {
			return md5
		}
	}
}

func part2(input []string) int {
	return 0
}
