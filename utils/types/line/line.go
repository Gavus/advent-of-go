package line

import (
	"fmt"
	"github.com/Gavus/advent-of-go/utils/types/point"
	_ "sort"
)

type Line struct {
	start point.Point
	end   point.Point
}

type Lines []Line

func Make(start point.Point, end point.Point) Line {
	return Line{start, end}
}

func Makes(line ...Line) Lines {
	lines := Lines{}
	for _, v := range line {
		lines = append(lines, v)
	}
	return lines
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
	if l.start.Y == l.end.Y {
		for x := l.start.X; x <= l.end.X; x++ {
			points = append(points, point.Make(x, l.start.Y))
		}
		for x := l.start.X; x >= l.end.X; x-- {
			points = append(points, point.Make(x, l.start.Y))
		}
	} else if l.start.X == l.end.X {
		for y := l.start.Y; y <= l.end.Y; y++ {
			points = append(points, point.Make(l.start.X, y))
		}
		for y := l.start.Y; y >= l.end.Y; y-- {
			points = append(points, point.Make(l.start.X, y))
		}
	}
	return points
}

func (l Line) ToPoints() point.Points {
	points := point.Points{}

	if l.start.Y == l.end.Y || l.start.X == l.end.X {
		return l.StraightToPoints()
	} else {
		xdiff := l.end.X - l.start.X
		ydiff := l.end.Y - l.start.Y
		if xdiff > 0 && ydiff > 0 {
			for i := 0; i <= xdiff; i++ {
				points = append(points, point.Make(l.start.X+i, l.start.Y+i))
			}
		} else if xdiff > 0 && ydiff < 0 {
			for i := 0; i <= xdiff; i++ {
				points = append(points, point.Make(l.start.X+i, l.start.Y-i))
			}
		} else if xdiff < 0 && ydiff > 0 {
			for i := 0; i >= xdiff; i-- {
				points = append(points, point.Make(l.start.X+i, l.start.Y-i))
			}
		} else if xdiff < 0 && ydiff < 0 {
			for i := 0; i >= xdiff; i-- {
				points = append(points, point.Make(l.start.X+i, l.start.Y+i))
			}
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

func (lines Lines) ToPoints() point.Points {
	points := point.Points{}

	for _, line := range lines {
		points = append(points, line.ToPoints()...)
	}
	return points
}

func Draw(points point.Points) string {
	str := ""
	xend, yend := points.LargestXY()

	for y := 0; y <= yend; y++ {
		for x := 0; x <= xend; x++ {
			count := len(points.Contains(point.Make(x, y)))
			if count > 0 {
				str += fmt.Sprintf("%d", count)
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}

func SumOverlaps(p point.Points) int {
	str := Draw(p)
	sum := 0
	for _, v := range str {
		if v > '1' && v <= '9' {
			sum++
		}
	}
	return sum
}
