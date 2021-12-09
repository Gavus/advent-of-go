package main

import (
	"github.com/Gavus/advent-of-go/testutils/stringf"
	"strings"
	"testing"
)

func createInput() []string {
	return strings.Split("^v^v^v^v^v", "")
}

func TestPart1(t *testing.T) {
	g, w := part1(createInput()), 2
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestPart2(t *testing.T) {
	g, w := part2(createInput()), 10
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
