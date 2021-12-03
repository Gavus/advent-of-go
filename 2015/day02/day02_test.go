package main

import (
	"github.com/Gavus/advent-of-go/testutils/stringf"
	"testing"
)

func createInput() []string {
	return []string{"2x3x4", "1x1x10"}
}

func TestPart1(t *testing.T) {
	g, w := part1(createInput()), 58+43
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestPart2(t *testing.T) {
	g, w := part2(createInput()), 34+14
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
