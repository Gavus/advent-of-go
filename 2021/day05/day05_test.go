package main

import (
	"github.com/Gavus/advent-of-go/testutils/stringf"
	"strings"
	"testing"
)

func createInput() []string {
	str := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	return strings.Split(str, "\n")
}

func TestPart1(t *testing.T) {
	g, w := part1(createInput()), 5
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestPart2(t *testing.T) {
	g, w := part2(createInput()), 12
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
