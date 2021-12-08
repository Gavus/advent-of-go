package submarine

import (
	"fmt"
	"github.com/Gavus/advent-of-go/utils/log"
	"github.com/Gavus/advent-of-go/utils/types/direction"
	"github.com/Gavus/advent-of-go/utils/types/point"
)

type Submarine struct {
	pos point.Point
	aim int
}

func (s Submarine) String() string {
	return s.pos.String()
}

type SubmarineInstruction struct {
	dir  direction.Direction
	dist int
}

type SubmarineInstructions []SubmarineInstruction

func Make(pos point.Point) Submarine {
	return Submarine{pos, 0}
}

func MakeSubmarineInstruction(dirStr string, dist int) SubmarineInstruction {
	var dir direction.Direction
	switch dirStr {
	case "forward":
		dir = direction.DirForward
	case "up":
		dir = direction.DirUp
	case "down":
		dir = direction.DirDown
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
	case direction.DirForward:
		s.pos.X += instr.dist
	case direction.DirUp:
		s.pos.Y -= instr.dist
	case direction.DirDown:
		s.pos.Y += instr.dist
	default:
		log.Warn.Print("Invalid submarine instruction: ", instr.dir)
	}
}

func (s *Submarine) Move2(instr SubmarineInstruction) {
	switch instr.dir {
	case direction.DirForward:
		s.pos.X += instr.dist
		s.pos.Y += s.aim * instr.dist
	case direction.DirUp:
		s.aim -= instr.dist
	case direction.DirDown:
		s.aim += instr.dist
	default:
		log.Warn.Print("Invalid submarine instruction: ", instr.dir)
	}
}

func ToSubmarineInstructions(input []string) SubmarineInstructions {
	instr := SubmarineInstructions{}
	for _, line := range input {
		var dir string
		var dist int
		_, err := fmt.Sscanf(line, "%s %d", &dir, &dist)
		if err != nil {
			panic(err)
		}
		instr = append(instr, MakeSubmarineInstruction(dir, dist))
	}

	return instr
}
