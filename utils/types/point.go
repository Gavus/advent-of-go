package types

import (
	"fmt"
)

// Point in a 2 Dimensional space
type Point struct {
	X int
	Y int
}

func MakePoint(x int, y int) Point {
	return Point{x, y}
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
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
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
