package present

import (
	"fmt"
	"sort"
)

// Present with its dimensions.
type Present struct {
	l int
	w int
	h int
}

func Make(l, w, h int) Present {
	return Present{l, w, h}
}

type Presents []Present

// Convert Present into a string.
func (p Present) String() string {
	return fmt.Sprintf("[%d %d %d]", p.l, p.w, p.h)
}

// Calculate Present's surface area.
func (p Present) SurfaceArea() int {
	return 2*p.l*p.w + 2*p.w*p.h + 2*p.h*p.l
}

// Calculate area of Present's smallest side.
func (p Present) SmallestSideArea() int {
	sideAreas := []int{p.l * p.w, p.l * p.h, p.w * p.h}
	sort.Ints(sideAreas)
	return sideAreas[0]
}

// Calculate ribbon length needed for Present.
func (p Present) RibbonLength() int {
	length := 0
	sides := []int{p.l, p.w, p.h}
	sort.Ints(sides)
	a, b, c := sides[0], sides[1], sides[2]

	length = 2*a + 2*b + a*b*c

	return length
}

// Convert strings into Presents.
func ToPresents(input []string) (Presents, error) {
	presents := Presents{}

	for _, value := range input {
		var l, w, h int
		_, err := fmt.Sscanf(value, "%dx%dx%d", &l, &w, &h)
		if err != nil {
			return nil, err
		}
		presents = append(presents, Make(l, w, h))
	}

	return presents, nil
}
