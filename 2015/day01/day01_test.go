package main

import (
	"github.com/Gavus/advent-of-go/testutils/stringf"
	"strings"
	"testing"
)

func createInput() []string {
	return strings.Split("))(((((", "")
}

func TestPart1(t *testing.T) {
	g, _ := part1(createInput(), 0)
	w := 3
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestPart2(t *testing.T) {
	g, w := part2(createInput()), 1
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
