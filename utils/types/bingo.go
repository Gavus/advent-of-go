package types

import (
	"errors"
	"fmt"
	"github.com/Gavus/advent-of-go/utils/calc"
)

const (
	bingoSize = 5
)

type BingoInput []int

type BingoBoard struct {
	size    int
	entries BingoInput
	hits    []int
}

func MakeBingoBoard(input BingoInput) BingoBoard {
	bb := BingoBoard{bingoSize, BingoInput{}, []int{}}

	if len(input) != bingoSize*bingoSize {
		panic(errors.New("Wrong input length given"))
	}

	for i := 0; i < bb.size; i++ {
		for j := 0; j < bb.size; j++ {
			bb.entries = append(bb.entries, input[i*bingoSize+j])
		}
	}

	return bb
}

// Convert BingoBoard to string.
func (b BingoBoard) String() string {
	str := ""

	for i := 0; i < b.size; i++ {
		for _, v := range b.Row(i) {
			var format string
			if calc.Contains(b.hits, v) {
				format = "*%2d*"
			} else {
				format = " %d "
			}
			str += fmt.Sprintf(format, v)
		}
		str += "\n"
	}
	return str
}

// Add BingoInput to BingoBoard.
func (b *BingoBoard) Add(num int) {
	b.hits = append(b.hits, num)
}

// Get Row.
func (b BingoBoard) Row(i int) []int {
	return b.entries[b.size*i : b.size*i+b.size]
}

// Get Column.
func (b BingoBoard) Column(i int) []int {
	col := []int{}
	for j := 0; j < b.size; j++ {
		col = append(col, b.entries[b.size*j+i])
	}

	return col
}

// Check if column or row got a bingo.
func (b BingoBoard) BingoRow(col []int) bool {
	for _, v := range col {
		if !calc.Contains(b.hits, v) {
			return false
		}
	}

	return true

}

// Check if BingoBoard has any bingo.
func (b BingoBoard) Bingo() bool {
	for i := 0; i < b.size; i++ {
		row := b.Row(i)
		if b.BingoRow(row) {
			return true
		}

		col := b.Column(i)
		if b.BingoRow(col) {
			return true
		}
	}

	return false
}

func (b BingoBoard) sumUnmarked() int {
	vals := []int{}
	for _, v := range b.entries {
		if !calc.Contains(b.hits, v) {
			vals = append(vals, v)
		}
	}

	return calc.Sum(vals)
}

func (b BingoBoard) Score() int {
	return b.sumUnmarked() * b.hits[len(b.hits)-1]
}

func RemoveIndex(bbs []BingoBoard, index int) []BingoBoard {
	if index < 0 || index >= len(bbs) || len(bbs) == 1 {
		return []BingoBoard{}
	} else if index == len(bbs)-1 {
		return bbs[:index]

	} else {
		return append(bbs[:index], bbs[index+1:]...)
	}
}
