package date

import (
	"github.com/Gavus/advent-of-code/testutils/stringf"
	"testing"
	"time"
)

func TestDateParseDate(t *testing.T) {
	g, w := ParseDate("2021-11-20"), time.Unix(1637366400, 0).UTC()
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestDateTimeToDayYearString(t *testing.T) {
	g1, g2 := ToDayYearString(time.Unix(1637366400, 0).UTC())
	w1, w2 := "2021", "20"
	if g1 != w1 {
		t.Errorf(stringf.Mismatch, g1, w1)
	}

	if g2 != w2 {
		t.Errorf(stringf.Mismatch, g2, w2)
	}
}
