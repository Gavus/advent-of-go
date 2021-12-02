package conv

import (
	"github.com/Gavus/advent-of-go/testutils/stringf"
	"github.com/Gavus/advent-of-go/utils/types"
	"testing"
)

func TestToPresents(t *testing.T) {
	strings := []string{"1x2x3", "1x1x10"}
	g, err := ToPresents(strings)
	w := make([]types.Present, 2)
	w[0].Set(1, 2, 3)
	w[1].Set(1, 1, 10)

	for i, v := range g {
		if v != w[i] {
			t.Errorf(stringf.Mismatch, g, w)
		}
	}

	if err != nil {
		t.Errorf(stringf.Mismatch, err, nil)
	}
}

func TestToStrings(t *testing.T) {
	g := ToStrings([]byte("(("), "")
	w := []string{"(", "("}
	for i, v := range g {
		if v != w[i] {
			t.Errorf(stringf.Mismatch, v, w[i])
		}
	}

}
