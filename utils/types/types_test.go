package types

import (
	"github.com/Gavus/advent-of-go/testutils/stringf"
	"testing"
)

func TestPresentString(t *testing.T) {
	g, w := (Present{}.String()), "[0 0 0]"
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}

	g, w = (Present{1, 2, 3}.String()), "[1 2 3]"
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestPresentSet(t *testing.T) {
	g, w := Present{}, Present{1, 2, 3}
	g.Set(1, 2, 3)
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestPresentSurfaceArea(t *testing.T) {
	g, w := (Present{2, 3, 4}.SurfaceArea()), 52
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}

	g, w = (Present{1, 1, 10}.SurfaceArea()), 42
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestPresentSmallestSideArea(t *testing.T) {
	g, w := (Present{2, 3, 4}.SmallestSideArea()), 6
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}

	g, w = (Present{1, 1, 10}.SmallestSideArea()), 1
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}

func TestPresentRibbonLength(t *testing.T) {
	g, w := (Present{2, 3, 4}.RibbonLength()), 34
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}

	g, w = (Present{1, 1, 10}.RibbonLength()), 14
	if g != w {
		t.Errorf(stringf.Mismatch, g, w)
	}
}
