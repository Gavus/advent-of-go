package types

type Direction int

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
