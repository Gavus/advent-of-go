package main

import (
	"github.com/Gavus/advent-of-go/testutils/stringf"
	"testing"
)

func createInput() []string {
	return []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263"}
}

func TestPart1(t *testing.T) {
	g, w := part1(createInput()), 7
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestPart2(t *testing.T) {
	g, w := part2(createInput()), 5
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
