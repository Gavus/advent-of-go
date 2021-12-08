package direction

import (
	"errors"
)

type Direction int
type Directions []Direction

const (
	DirLeft = iota
	DirRight
	DirUp
	DirDown
	DirForward
)

func (d Direction) String() string {
	switch d {
	case DirLeft:
		return "<"
	case DirRight:
		return ">"
	case DirDown:
		return "v"
	case DirUp:
		return "^"
	case DirForward:
		return "f"
	default:
		return "error"
	}
}

// Convert strings to Directions.
func ToDirections(input []string) (Directions, error) {
	dirs := Directions{}
	for _, v := range input {
		switch v {
		case "<":
			dirs = append(dirs, DirLeft)
		case ">":
			dirs = append(dirs, DirRight)
		case "^":
			dirs = append(dirs, DirUp)
		case "v":
			dirs = append(dirs, DirDown)
		default:
			return nil, errors.New("Invalid character found in input")
		}
	}
	return dirs, nil
}
