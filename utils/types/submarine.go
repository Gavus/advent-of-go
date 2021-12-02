package types

import (
	"github.com/Gavus/advent-of-code/utils/log"
)

type Submarine struct {
	pos Point
	aim int
}

func (s Submarine) String() string {
	return s.pos.String()
}

type SubmarineInstruction struct {
	dir  Direction
	dist int
}

func MakeSubmarine(pos Point) Submarine {
	return Submarine{pos, 0}
}

func MakeSubmarineInstruction(dirStr string, dist int) SubmarineInstruction {
	var dir Direction
	switch dirStr {
	case "forward":
		dir = DirForward
	case "up":
		dir = DirUp
	case "down":
		dir = DirDown
	default:
		log.Warn.Print("Invalid direction: ", dirStr)
	}
	return SubmarineInstruction{dir, dist}
}

func (s Submarine) Mult() int {
	return s.pos.X * s.pos.Y
}

func (s *Submarine) Move(instr SubmarineInstruction) {
	switch instr.dir {
	case DirForward:
		s.pos.X += instr.dist
	case DirUp:
		s.pos.Y -= instr.dist
	case DirDown:
		s.pos.Y += instr.dist
	default:
		log.Warn.Print("Invalid submarine instruction: ", instr.dir)
	}
}

func (s *Submarine) Move2(instr SubmarineInstruction) {
	switch instr.dir {
	case DirForward:
		s.pos.X += instr.dist
		s.pos.Y += s.aim * instr.dist
	case DirUp:
		s.aim -= instr.dist
	case DirDown:
		s.aim += instr.dist
	default:
		log.Warn.Print("Invalid submarine instruction: ", instr.dir)
	}
}
