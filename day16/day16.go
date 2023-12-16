package day16

import (
	"bufio"
	"os"

	"github.com/brunorene/adventofcode2023/common"
)

type PieceType rune
type Direction int

const (
	Right Direction = iota
	Left
	Up
	Down
	None
	HorzSplitter PieceType = '-'
	VertSplitter PieceType = '|'
	SWNEMirror   PieceType = '/'
	SENWMirror   PieceType = '\\'
)

type BeamID int

type Coords struct {
	X, Y int
}

type Contraption struct {
	Layout    []string
	Energized map[Direction]map[Coords]struct{}
}

func parse() (parsed Contraption) {
	file, err := os.Open("day16/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		parsed.Layout = append(parsed.Layout, line)
	}

	parsed.Energized = make(map[Direction]map[Coords]struct{})

	return parsed
}

func (c *Contraption) energize(direction Direction, coords Coords) {
	if coords.X < 0 || coords.X >= len(c.Layout[0]) || coords.Y < 0 || coords.Y >= len(c.Layout) {
		return
	}

	if _, exists := c.Energized[direction]; !exists {
		c.Energized[direction] = make(map[Coords]struct{})
	}

	if _, exists := c.Energized[direction][coords]; exists {
		return
	}

	c.Energized[direction][coords] = struct{}{}

	switch PieceType(c.Layout[coords.Y][coords.X]) {
	case HorzSplitter:
		switch direction {
		case Right:
			c.energize(Right, Coords{coords.X + 1, coords.Y})
		case Left:
			c.energize(Left, Coords{coords.X - 1, coords.Y})
		case Down:
			c.energize(Left, Coords{coords.X - 1, coords.Y})
			c.energize(Right, Coords{coords.X + 1, coords.Y})
		case Up:
			c.energize(Left, Coords{coords.X - 1, coords.Y})
			c.energize(Right, Coords{coords.X + 1, coords.Y})
		default:
		}
	case VertSplitter:
		switch direction {
		case Right:
			c.energize(Up, Coords{coords.X, coords.Y - 1})
			c.energize(Down, Coords{coords.X, coords.Y + 1})
		case Left:
			c.energize(Up, Coords{coords.X, coords.Y - 1})
			c.energize(Down, Coords{coords.X, coords.Y + 1})
		case Down:
			c.energize(Down, Coords{coords.X, coords.Y + 1})
		case Up:
			c.energize(Up, Coords{coords.X, coords.Y - 1})
		default:
		}
	case SENWMirror:
		switch direction {
		case Right:
			c.energize(Down, Coords{coords.X, coords.Y + 1})
		case Left:
			c.energize(Up, Coords{coords.X, coords.Y - 1})
		case Down:
			c.energize(Right, Coords{coords.X + 1, coords.Y})
		case Up:
			c.energize(Left, Coords{coords.X - 1, coords.Y})
		default:
		}
	case SWNEMirror:
		switch direction {
		case Right:
			c.energize(Up, Coords{coords.X, coords.Y - 1})
		case Left:
			c.energize(Down, Coords{coords.X, coords.Y + 1})
		case Down:
			c.energize(Left, Coords{coords.X - 1, coords.Y})
		case Up:
			c.energize(Right, Coords{coords.X + 1, coords.Y})
		default:
		}
	default:
		switch direction {
		case Right:
			c.energize(Right, Coords{coords.X + 1, coords.Y})
		case Left:
			c.energize(Left, Coords{coords.X - 1, coords.Y})
		case Down:
			c.energize(Down, Coords{coords.X, coords.Y + 1})
		case Up:
			c.energize(Up, Coords{coords.X, coords.Y - 1})
		default:
		}
	}
}

func (c *Contraption) energizedCount() (sum int) {
	accum := make(map[Coords]struct{})

	for _, list := range c.Energized {
		for coords := range list {
			accum[coords] = struct{}{}
		}
	}

	return len(accum)
}

func Part1() int {
	contraption := parse()

	contraption.energize(Right, Coords{0, 0})

	return contraption.energizedCount()
}

func Part2() int {
	contraption := parse()

	var maxEnergized int

	for y, line := range contraption.Layout {
		for x := range line {
			direction := None

			if x == 0 {
				direction = Left
			}

			if x == len(line)-1 {
				direction = Right
			}

			if y == 0 {
				direction = Down
			}

			if y == len(line)-1 {
				direction = Up
			}

			if direction == None {
				continue
			}

			current := calcEnergized(direction, Coords{x, y})

			if current > maxEnergized {
				maxEnergized = current
			}
		}
	}

	return maxEnergized
}

func calcEnergized(direction Direction, coords Coords) int {
	contraption := parse()

	contraption.energize(direction, coords)

	return contraption.energizedCount()
}
