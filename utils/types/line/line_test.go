package line

import (
	"github.com/Gavus/advent-of-go/testutils/stringf"
	"github.com/Gavus/advent-of-go/utils/types/point"
	"testing"
)

func TestStraightToPointsLeftToRight(t *testing.T) {
	// 0,0 -> 5,0
	line := Make(point.Make(0, 0), point.Make(5, 0))
	g := line.StraightToPoints()
	w := point.Makes(point.Make(0, 0), point.Make(1, 0), point.Make(2, 0), point.Make(3, 0), point.Make(4, 0), point.Make(5, 0))

	if !g.Equal(w) {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
func TestStraightToPointsRightToLeft(t *testing.T) {
	// 5,0 -> 0,0
	line := Make(point.Make(5, 0), point.Make(0, 0))
	g := line.StraightToPoints()
	w := point.Makes(point.Make(5, 0), point.Make(4, 0), point.Make(3, 0), point.Make(2, 0), point.Make(1, 0), point.Make(0, 0))

	if !g.Equal(w) {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
func TestStraightToPointsUpToDown(t *testing.T) {
	// 0,0 -> 0,5
	line := Make(point.Make(0, 0), point.Make(0, 5))
	g := line.StraightToPoints()
	w := point.Makes(point.Make(0, 0), point.Make(0, 1), point.Make(0, 2), point.Make(0, 3), point.Make(0, 4), point.Make(0, 5))

	if !g.Equal(w) {
		t.Errorf(stringf.Mismatch, g, w)
	}

}
func TestStraightToPointsDownToUp(t *testing.T) {
	// 0,5 -> 0,0
	line := Make(point.Make(0, 5), point.Make(0, 0))
	g := line.StraightToPoints()
	w := point.Makes(point.Make(0, 5), point.Make(0, 4), point.Make(0, 3), point.Make(0, 2), point.Make(0, 1), point.Make(0, 0))

	if !g.Equal(w) {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestLinesStraightToPoints(t *testing.T) {
	// 0,0 -> 0,2 and -5,0 -> -2, 0.
	lines := Makes(Make(point.Make(0, 0), point.Make(0, 2)), Make(point.Make(-5, 0), point.Make(-2, 0)))
	g := lines.StraightToPoints()
	w := point.Makes(point.Make(0, 0), point.Make(0, 1), point.Make(0, 2), point.Make(-5, 0), point.Make(-4, 0), point.Make(-3, 0), point.Make(-2, 0))

	if !g.Equal(w) {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestLinesDraw(t *testing.T) {
	lines := Makes(Make(point.Make(3, 0), point.Make(3, 3)), Make(point.Make(3, 2), point.Make(0, 2)))
	g := Draw(lines.StraightToPoints())
	w := "...1\n...1\n1112\n...1\n"

	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestLinesSumStraightOverlaps(t *testing.T) {
	lines := Makes(Make(point.Make(3, 0), point.Make(3, 3)), Make(point.Make(3, 2), point.Make(0, 2)))
	g := SumOverlaps(lines.StraightToPoints())
	w := 1

	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestLineToPointsLeftToRightDown(t *testing.T) {
	line := Make(point.Make(1,1), point.Make(3,3))
	g := line.ToPoints()
	w := point.Makes(point.Make(1,1), point.Make(2,2), point.Make(3,3))

	if !g.Equal(w) {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestLineToPointsLeftToRightUp(t *testing.T) {
	line := Make(point.Make(1,3), point.Make(3,1))
	g := line.ToPoints()
	w := point.Makes(point.Make(1,3), point.Make(2,2), point.Make(3,1))

	if !g.Equal(w) {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestLineToPointsRightToLeftDown(t *testing.T) {
	line := Make(point.Make(3,1), point.Make(1,3))
	g := line.ToPoints()
	w := point.Makes(point.Make(3,1), point.Make(2,2), point.Make(1,3))

	if !g.Equal(w) {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestLineToPointsRightToLeftUp(t *testing.T) {
	line := Make(point.Make(3,3), point.Make(1,1))
	g := line.ToPoints()
	w := point.Makes(point.Make(3,3), point.Make(2,2), point.Make(1,1))

	if !g.Equal(w) {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
