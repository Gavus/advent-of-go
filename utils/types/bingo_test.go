package types

import (
	"github.com/Gavus/advent-of-go/testutils/stringf"
	"github.com/Gavus/advent-of-go/utils/calc"
	"testing"
)

func createBingoBoard() BingoBoard {
	input := []int{
		14, 21, 17, 24, 4,
		10, 16, 15, 9, 19,
		18, 8, 23, 26, 20,
		22, 11, 13, 6, 5,
		2, 0, 12, 3, 7}
	return MakeBingoBoard(input)
}

func TestRow(t *testing.T) {
	bb := createBingoBoard()
	g := bb.Row(0)
	w := []int{14, 21, 17, 24, 4}
	if !calc.Equal(g, w) {
		t.Errorf(stringf.Mismatch, g, w)
	}

	g = bb.Row(3)
	w = []int{22, 11, 13, 6, 5}
	if !calc.Equal(g, w) {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestColumn(t *testing.T) {
	bb := createBingoBoard()
	g := bb.Column(0)
	w := []int{14, 10, 18, 22, 2}
	if !calc.Equal(g, w) {
		t.Errorf(stringf.Mismatch, g, w)
	}

	g = bb.Column(3)
	w = []int{24, 9, 26, 6, 3}
	if !calc.Equal(g, w) {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestBingoRow(t *testing.T) {
	bb := createBingoBoard()
	g := bb.BingoRow(bb.Row(0))
	w := false

	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}

	bb.Add(14)
	bb.Add(21)
	bb.Add(17)
	bb.Add(24)
	bb.Add(4)

	g = bb.BingoRow(bb.Row(0))
	w = true

	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestBingo(t *testing.T) {
	bb := createBingoBoard()
	g := bb.Bingo()
	w := false

	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}

	bb.Add(14)
	bb.Add(21)
	bb.Add(17)
	bb.Add(24)
	bb.Add(4)

	g = bb.Bingo()
	w = true

	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestSumUnmarked(t *testing.T) {
	bb := createBingoBoard()
	bb.Add(14)
	bb.Add(21)
	bb.Add(17)
	bb.Add(24)
	bb.Add(4)
	bb.Add(9)
	bb.Add(23)
	bb.Add(11)
	bb.Add(5)
	bb.Add(2)
	bb.Add(0)
	bb.Add(7)

	g := bb.sumUnmarked()
	w := 188

	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestScore(t *testing.T) {
	bb := createBingoBoard()
	bb.Add(14)
	bb.Add(21)
	bb.Add(17)
	bb.Add(24)
	bb.Add(4)
	bb.Add(9)
	bb.Add(23)
	bb.Add(11)
	bb.Add(5)
	bb.Add(2)
	bb.Add(0)
	bb.Add(7)

	g := bb.Score()
	w := 188*7

	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
