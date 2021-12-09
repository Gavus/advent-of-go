package point

import (
	"fmt"
	"github.com/Gavus/advent-of-go/utils/types/direction"
)

// Point in a 2 Dimensional space
type Point struct {
	X int
	Y int
}

type Points []Point

func Make(x int, y int) Point {
	return Point{x, y}
}

func Makes(point ...Point) Points {
	points := Points{}
	for _, v := range point {
		points = append(points, v)
	}
	return points
}

// Move a Point to specified Direction.
func Move(from Point, dir direction.Direction) Point {
	pos := from
	switch dir {
	case direction.DirLeft:
		pos.X--
	case direction.DirRight:
		pos.X++
	case direction.DirUp:
		pos.Y--
	case direction.DirDown:
		pos.Y++
	}

	return pos
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

// Returns Points without any dublicates.
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

func (a Point) Less(b Point) bool {
	if a.X < b.X && a.Y <= b.Y {
		return true
	} else if a.X > b.X && a.Y < b.Y {
		return true
	} else if a.Y < b.Y {
		return true
	}
	return false
}

func (points Points) Len() int {
	return len(points)
}

func (points Points) Less(i, j int) bool {
	return points[i].Less(points[j])
}

func (points Points) Swap(i, j int) {
	points[i], points[j] = points[j], points[i]
}

func (p1 Points) Equal(p2 Points) bool {
	if len(p1) != len(p2) {
		return false
	}

	for i, v := range p1 {
		if v != p2[i] {
			return false
		}
	}
	return true
}

func (points Points) LargestXY() (int, int) {
	x, y := 0, 0
	for _, p := range points {
		if p.X > x {
			x = p.X
		}

		if p.Y > y {
			y = p.Y
		}
	}

	return x, y
}

// Returns Points equal arg p.
func (points Points) Contains(p Point) Points {
	pts := Points{}
	for _, v := range points {
		if v == p {
			pts = append(pts, p)
		}
	}
	return pts
}
