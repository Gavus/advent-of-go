package main

import (
	"github.com/Gavus/advent-of-go/testutils/stringf"
	"testing"
)

func createInput() []string {
	return []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010"}
}

func TestPart1(t *testing.T) {
	g, w := part1(createInput()), uint(198)
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestPart2(t *testing.T) {
	g, w := part2(createInput()), uint(230)
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
