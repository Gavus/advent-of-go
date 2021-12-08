package bingo

import (
	"errors"
	"fmt"
	"strings"
	"strconv"
	"github.com/Gavus/advent-of-go/utils/types/ints"
)

const (
	bingoSize = 5
)

type BingoInput ints.Ints

type BingoBoard struct {
	size    int
	entries ints.Ints
	hits    ints.Ints
}

func Make(input ints.Ints) BingoBoard {
	bb := BingoBoard{bingoSize, ints.Ints{}, ints.Ints{}}

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
			if b.hits.Contains(v) {
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
func (b BingoBoard) Row(i int) ints.Ints {
	return b.entries[b.size*i : b.size*i+b.size]
}

// Get Column.
func (b BingoBoard) Column(i int) ints.Ints {
	col := []int{}
	for j := 0; j < b.size; j++ {
		col = append(col, b.entries[b.size*j+i])
	}

	return col
}

// Check if column or row got a bingo.
func (b BingoBoard) BingoRow(col ints.Ints) bool {
	for _, v := range col {
		if !b.hits.Contains(v) {
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
	vals := ints.Ints{}
	for _, v := range b.entries {
		if !b.hits.Contains(v) {
			vals = append(vals, v)
		}
	}

	return vals.Sum()
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


// Convert strings to BingoInput and BingoBoards.
func ToBingo(input []string) (BingoInput, []BingoBoard) {
	bi := BingoInput{}
	bbs := []BingoBoard{}

	// BingoInput
	for _, v := range strings.Split(input[0], ",") {
		x, err := strconv.Atoi(v)
		if err != nil {
			continue
		}

		bi = append(bi, x)
	}

	// BingoBoard
	row := []string{}
	for _, v := range input[1:] {
		if v == "" {
			if len(row) != 0 {
				entries, err := ints.ToInts(row)
				if err != nil {
					panic(err)
				}
				bbs = append(bbs, Make(entries))
				row = []string{}
			}
			continue
		}

		for _, v := range strings.Split(v, " ") {
			x := strings.TrimSpace(v)
			if x != "" {
				row = append(row, x)
			}
		}
	}

	return bi, bbs
}
