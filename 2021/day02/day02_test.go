package main

import (
	"github.com/Gavus/advent-of-go/testutils/stringf"
	"testing"
)

func createInput() []string {
	return []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2"}
}

func TestPart1(t *testing.T) {
	g, w := part1(createInput()), 150
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestPart2(t *testing.T) {
	g, w := part2(createInput()), 900
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
