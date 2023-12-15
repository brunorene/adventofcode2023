package day10

import (
	"bufio"
	"os"
	"slices"
	"sort"

	"github.com/brunorene/adventofcode2023/common"
)

type PipeType rune

const (
	NorthSouth PipeType = '|'
	NorthEast  PipeType = 'L'
	NorthWest  PipeType = 'J'
	SouthWest  PipeType = '7'
	SouthEast  PipeType = 'F'
	EastWest   PipeType = '-'
	Start      PipeType = 'S'
	MaxCol              = 140
	MaxRow              = 140
)

type Pipe struct {
	PipeType                 PipeType
	North, South, East, West *Coords
}

type Coords struct {
	X, Y int
}

type Area struct {
	Pipes map[Coords]Pipe
	Start Coords
}

func (p Pipe) ends() (left, right Coords) {
	first := true

	for _, coords := range []*Coords{p.North, p.South, p.East, p.West} {
		if coords != nil {
			if first {
				left = *coords

				first = false

				continue
			}

			right = *coords

			return
		}
	}

	return
}

func createPipe(name PipeType, row, col int) Pipe {
	var north, south, east, west *Coords

	if slices.Contains([]PipeType{NorthSouth, NorthEast, NorthWest, Start}, name) && row > 0 {
		north = &Coords{X: col, Y: row - 1}
	}

	if slices.Contains([]PipeType{NorthSouth, SouthWest, SouthEast, Start}, name) && row < MaxRow-1 {
		south = &Coords{X: col, Y: row + 1}
	}

	if slices.Contains([]PipeType{EastWest, NorthEast, SouthEast, Start}, name) && col < MaxCol-1 {
		east = &Coords{X: col + 1, Y: row}
	}

	if slices.Contains([]PipeType{EastWest, NorthWest, SouthWest, Start}, name) && col > 0 {
		west = &Coords{X: col - 1, Y: row}
	}

	return Pipe{
		PipeType: name,
		North:    north,
		South:    south,
		East:     east,
		West:     west,
	}
}

func startNewName(north, south, east, west *Coords) PipeType {
	if north != nil {
		if south != nil {
			return NorthSouth
		}

		if east != nil {
			return NorthEast
		}

		if west != nil {
			return NorthWest
		}
	}

	if south != nil {
		if east != nil {
			return SouthEast
		}

		if west != nil {
			return SouthWest
		}
	}

	return EastWest
}

func refreshStart(area Area, start Pipe) Pipe {
	var north, south, east, west *Coords

	if slices.Contains([]PipeType{SouthEast, SouthWest, NorthSouth}, area.Pipes[*start.North].PipeType) {
		north = start.North
	}

	if slices.Contains([]PipeType{NorthEast, NorthWest, NorthSouth}, area.Pipes[*start.South].PipeType) {
		south = start.South
	}

	if slices.Contains([]PipeType{SouthEast, NorthEast, EastWest}, area.Pipes[*start.West].PipeType) {
		west = start.West
	}

	if slices.Contains([]PipeType{SouthWest, NorthWest, EastWest}, area.Pipes[*start.East].PipeType) {
		east = start.East
	}

	return Pipe{
		PipeType: startNewName(north, south, east, west),
		North:    north,
		South:    south,
		East:     east,
		West:     west,
	}
}

func parse() (parsed Area) {
	file, err := os.Open("day10/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	parsed = Area{
		Pipes: make(map[Coords]Pipe),
	}

	var row int

	for scanner.Scan() {
		line := scanner.Text()

		for col, pipeType := range line {
			if PipeType(pipeType) == Start {
				parsed.Start = Coords{col, row}
			}

			parsed.Pipes[Coords{col, row}] = createPipe(PipeType(pipeType), row, col)
		}

		row++
	}

	parsed.Pipes[parsed.Start] = refreshStart(parsed, parsed.Pipes[parsed.Start])

	return parsed
}

func (a *Area) GetDistance() (distance int64) {
	beforeLeft := a.Start
	beforeRight := a.Start
	left, right := a.Pipes[a.Start].ends()

	distance = 1

	for left != right {
		distance++

		nextLeft1, nextLeft2 := a.Pipes[left].ends()
		nextRight1, nextRight2 := a.Pipes[right].ends()

		if nextLeft1 == beforeLeft {
			beforeLeft = left
			left = nextLeft2
		} else {
			beforeLeft = left
			left = nextLeft1
		}

		if nextRight1 == beforeRight {
			beforeRight = right
			right = nextRight2
		} else {
			beforeRight = right
			right = nextRight1
		}
	}

	return distance
}

func Part1() int64 {
	area := parse()

	return area.GetDistance()
}

func addSorted(list sort.IntSlice, item int) sort.IntSlice {
	idx := list.Search(item)

	list = append(list, 0)
	copy(list[idx+1:], list[idx:])

	list[idx] = item

	return list
}

func (a *Area) BelongToLoop() (result map[int]sort.IntSlice) {
	result = make(map[int]sort.IntSlice)

	beforeLeft := a.Start
	beforeRight := a.Start
	left, right := a.Pipes[a.Start].ends()

	result[a.Start.Y] = addSorted(result[a.Start.Y], a.Start.X)

	for left != right {
		result[left.Y] = addSorted(result[left.Y], left.X)
		result[right.Y] = addSorted(result[right.Y], right.X)

		nextLeft1, nextLeft2 := a.Pipes[left].ends()
		nextRight1, nextRight2 := a.Pipes[right].ends()

		if nextLeft1 == beforeLeft {
			beforeLeft = left
			left = nextLeft2
		} else {
			beforeLeft = left
			left = nextLeft1
		}

		if nextRight1 == beforeRight {
			beforeRight = right
			right = nextRight2
		} else {
			beforeRight = right
			right = nextRight1
		}
	}

	result[right.Y] = addSorted(result[right.Y], right.X)

	return result
}

func (a *Area) insideLoop(loop map[int]sort.IntSlice) (result []Coords) {
	for coord := range a.Pipes {
		idx := loop[coord.Y].Search(coord.X)

		// part of loop
		if idx < len(loop[coord.Y]) && loop[coord.Y][idx] == coord.X {
			continue
		}

		var northFacingCount int

		for i := 0; i < idx; i++ {
			if slices.Contains([]PipeType{NorthSouth, NorthEast, NorthWest},
				a.Pipes[Coords{loop[coord.Y][i], coord.Y}].PipeType) {
				northFacingCount++
			}
		}

		// outside - north facing loop parts count is even
		if northFacingCount%2 == 0 {
			continue
		}

		result = append(result, coord)
	}

	return result
}

func Part2() int {
	area := parse()

	return len(area.insideLoop(area.BelongToLoop()))
}
