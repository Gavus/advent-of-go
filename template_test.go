package main

import (
	"github.com/Gavus/advent-of-go/testutils/stringf"
	"strings"
	"testing"
)

func createInput() []string {
	str := `row1
row2
row3`

	return strings.Split(str, "\n")
}

func TestPart1(t *testing.T) {
	g, w := part1(createInput()), 0
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestPart2(t *testing.T) {
	g, w := part2(createInput()), 0
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
