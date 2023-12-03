package day03

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"

	"github.com/brunorene/adventofcode2023/common"
)

type partType int

const (
	numberType partType = iota
	symbolType
	maxRow = 140
	maxCol
	gearCount = 2
)

type coord struct {
	x, y int
}

type number struct {
	value  int
	coords []coord
}

type symbol struct {
	value int32
	coord coord
}

type part interface {
	around() []coord
	myType() partType
	getValue() int64
	getKey() string
}

func (s symbol) getKey() string {
	return fmt.Sprintf("%+v", s)
}

func (s symbol) getValue() int64 {
	return int64(s.value)
}

func (s symbol) myType() partType {
	return symbolType
}

func (s symbol) around() (result []coord) {
	for _, diffX := range []int{-1, 0, 1} {
		for _, diffY := range []int{-1, 0, 1} {
			if diffX == 0 && diffY == 0 {
				continue
			}

			border := coord{s.coord.x + diffX, s.coord.y + diffY}
			if border.x < 0 || border.x >= maxCol || border.y < 0 || border.y >= maxRow {
				continue
			}

			result = append(result, border)
		}
	}

	return result
}

func (n number) getKey() string {
	return fmt.Sprintf("%+v", n)
}

func (n number) getValue() int64 {
	return int64(n.value)
}

func (n number) myType() partType {
	return numberType
}

//nolint:cyclop // sorry
func (n number) around() (result []coord) {
	for _, mainDigit := range n.coords {
		for _, diffX := range []int{-1, 0, 1} {
			for _, diffY := range []int{-1, 0, 1} {
				if diffX == 0 && diffY == 0 {
					continue
				}

				border := coord{mainDigit.x + diffX, mainDigit.y + diffY}
				if border.x < 0 || border.x >= maxCol || border.y < 0 || border.y >= maxRow {
					continue
				}

				var matches bool

				for _, digit := range append(n.coords, result...) {
					if border == digit {
						matches = true

						break
					}
				}

				if matches {
					continue
				}

				result = append(result, border)
			}
		}
	}

	return result
}

type engine struct {
	parts    []part
	location map[coord]part
}

func (e *engine) addNumber(numStr string, col, row int) string {
	if numStr != "" {
		num, err := strconv.Atoi(numStr)
		common.CheckError(err)

		var coords []coord

		for i := col - len(numStr); i < col; i++ {
			coords = append(coords, coord{i, row})
		}

		sort.Slice(coords, func(i, j int) bool {
			return coords[i].x < coords[j].x || (coords[i].x == coords[j].x && coords[i].y < coords[j].y)
		})

		item := number{value: num, coords: coords}

		e.parts = append(e.parts, item)

		for i := col - len(numStr); i < col; i++ {
			e.location[coord{i, row}] = item
		}

		return ""
	}

	return numStr
}

func parse() (parsed engine) {
	file, err := os.Open("day03/input")
	common.CheckError(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var yIdx int

	var numStr string

	parsed.location = make(map[coord]part)

	for scanner.Scan() {
		line := scanner.Text()

		for xIdx, val := range line {
			switch {
			case unicode.IsDigit(val):
				numStr += string(val)
			case val == '.':
				numStr = parsed.addNumber(numStr, xIdx, yIdx)
			default:
				numStr = parsed.addNumber(numStr, xIdx, yIdx)

				item := symbol{
					value: val,
					coord: coord{xIdx, yIdx},
				}

				parsed.parts = append(parsed.parts, item)
				parsed.location[item.coord] = item
			}
		}

		if numStr != "" {
			numStr = parsed.addNumber(numStr, maxCol, yIdx)
		}

		yIdx++
	}

	return parsed
}

func Part1() string {
	current := parse()

	var sum int64

	for Idx := range current.parts {
		if current.parts[Idx].myType() != numberType {
			continue
		}

		borders := current.parts[Idx].around()

		for _, border := range borders {
			part, exists := current.location[border]
			if exists && part.myType() == symbolType {
				sum += current.parts[Idx].getValue()

				break
			}
		}
	}

	return strconv.FormatInt(sum, 10)
}

func Part2() string {
	current := parse()

	var sum int64

	for Idx := range current.parts {
		if current.parts[Idx].myType() != symbolType || rune(current.parts[Idx].getValue()) != '*' {
			continue
		}

		borders := current.parts[Idx].around()

		adjacents := make(map[string]part)

		for _, border := range borders {
			part, exists := current.location[border]
			if exists && part.myType() == numberType {
				adjacents[part.getKey()] = part
			}
		}

		if len(adjacents) == gearCount {
			mult := int64(1)

			for _, part := range adjacents {
				mult *= part.getValue()
			}

			sum += mult
		}
	}

	return strconv.FormatInt(sum, 10)
}
