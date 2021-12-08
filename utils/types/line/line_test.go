package line

import (
	"github.com/Gavus/advent-of-go/testutils/stringf"
	"github.com/Gavus/advent-of-go/utils/types/point"
	"testing"
)

func TestStraightToPoints(t *testing.T) {
	line := Make(point.Make(0, 0), point.Make(5, 0))
	g := line.StraightToPoints()
	w := point.Points{point.Make(0, 0), point.Make(1, 0), point.Make(2, 0), point.Make(3, 0), point.Make(4, 0), point.Make(4, 0)}
	
	if g.Equal(w) {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
