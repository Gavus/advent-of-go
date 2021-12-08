package line

import (
	"fmt"
	_ "sort"
	"github.com/Gavus/advent-of-go/utils/types/point"
)

type Line struct {
	start point.Point
	end   point.Point
}

type Lines []Line

func Make(start point.Point, end point.Point) Line {
	return Line{start, end}
}

func StringsToLines(input []string) Lines {
	lines := Lines{}
	for _, row := range input {
		var a, b, c, d int
		_, err := fmt.Sscanf(row, "%d,%d -> %d,%d", &a, &b, &c, &d)
		if err != nil {
			panic(err)
		}
		lines = append(lines, Make(point.Make(a, b), point.Make(c, d)))
	}

	return lines
}

func (l Line) String() string {
	return fmt.Sprintf("%s -> %s", l.start, l.end)
}

func (l Line) StraightToPoints() point.Points {
	points := point.Points{}
	if l.start.X == l.end.X {
		for x := l.start.X; x <= l.end.X; x++ {
			points = append(points, point.Make(x, l.start.Y))
		}
		for x := l.start.X; x >= l.end.X; x-- {
			points = append(points, point.Make(x, l.start.Y))
		}
	} else if l.start.Y == l.end.Y {
		for y := l.start.Y; y <= l.end.Y; y++ {
			points = append(points, point.Make(l.start.X, y))
		}
		for y := l.start.Y; y >= l.end.Y; y-- {
			points = append(points, point.Make(l.start.X, y))
		}
	}
	return points
}

func (lines Lines) StraightToPoints() point.Points {
	points := point.Points{}

	for _, line := range lines {
		points = append(points, line.StraightToPoints()...)
	}
	return points
}

func (lines Lines) Draw() string {
	str := ""

	return str
}
