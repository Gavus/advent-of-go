package types

type Direction int

const (
	DirLeft = iota
	DirRight
	DirUp
	DirDown
)

func (d Direction) String() string {
	switch(d) {
	case DirLeft:
		return "<"
	case DirRight:
		return ">"
	case DirDown:
		return "v"
	case DirUp:
		return "^"
	default:
		return "error"
	}
}
