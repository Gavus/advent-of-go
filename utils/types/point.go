package types

import (
	"fmt"
)

// Point in a 2 Dimensional space
type Point struct {
	X int
	Y int
}

// Move a Point to specified Direction.
func Move(from Point, direction Direction) Point {
	pos := from
	switch(direction) {
	case DirLeft:
		pos.X--
	case DirRight:
		pos.X++
	case DirUp:
		pos.Y--
	case DirDown:
		pos.Y++
	}

	return pos
}

func (p Point) String() string {
	var str string
	fmt.Sscanf(str, "(%d,%d)", p.X, p.Y)
	return str

}

func Unique(points []Point) []Point {
	occurred := map[Point]bool{}
	result := []Point{}

	for e := range points {
		if occurred[points[e]] != true {
			occurred[points[e]] = true
			result = append(result, points[e])
		}
	}
	return result

}
