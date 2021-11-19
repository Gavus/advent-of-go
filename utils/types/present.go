package types

import (
	"fmt"
	"sort"
)

// Present with its dimensions.
type Present struct {
	h int
	l int
	w int
}

// Convert Present into a string.
func (p Present) String() string {
	return fmt.Sprintf("l: %d, w: %d, h: %d", p.l, p.w, p.h)
}

// Set Present's dimensions.
func (p *Present) Set(l int, w int, h int) {
	p.h = h
	p.l = l
	p.w = w
}

// Calculate Present's surface area.
func (p Present) SurfaceArea() int {
	return 2*p.l*p.w + 2*p.w*p.h + 2*p.h*p.l
}

// Calculate area of Present's smallest side.
func (p Present) SmallestSideArea() int {
	sideAreas := []int{p.l * p.w, p.l * p.h, p.w * p.h}
	sort.Ints(sideAreas)
	return sideAreas[0]
}

// Calculate ribbon length needed for Present.
func (p Present) RibbonLength() int {
	length := 0
	sides := []int{p.l, p.w, p.h}
	sort.Ints(sides)
	a, b, c := sides[0], sides[1], sides[2]

	length = 2*a + 2*b + a*b*c

	return length
}
