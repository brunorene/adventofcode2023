package day11

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/brunorene/adventofcode2023/common"
)

type SpaceType rune

const (
	Galaxy     SpaceType = '#'
	SquareSize           = 140
)

type Coords struct {
	X, Y int64
}

type Galaxies []Coords

type Pair [2]Coords

type GalaxyPairs map[Pair]struct{}

func (c1 Coords) distance(c2 Coords) int64 {
	xDiff := float64(max(c1.X, c2.X) - min(c1.X, c2.X))
	yDiff := float64(max(c1.Y, c2.Y) - min(c1.Y, c2.Y))

	return int64(xDiff + yDiff)
}

func (g GalaxyPairs) Contains(pair Pair) bool {
	left := pair[0]
	right := pair[1]

	if left.X > right.X || (left.X == right.X && left.Y > right.Y) {
		pair = Pair{right, left}
	}

	_, exists := g[pair]
	return exists
}

func (g GalaxyPairs) Add(pair Pair) {
	left := pair[0]
	right := pair[1]

	if left.X > right.X || left.Y > right.Y {
		pair = Pair{right, left}
	}

	g[pair] = struct{}{}
}

func expandUniverse(multInc int64, input []Coords, coordGetter func(Coords) int64, coordIncrease func(Coords, int64) Coords) []Coords {
	emptyLines := make(map[int]struct{})

	for i := 0; i < SquareSize; i++ {
		emptyLines[i] = struct{}{}
	}

	for _, galaxy := range input {
		delete(emptyLines, int(coordGetter(galaxy)))
	}

	emptyList := make([]int, 0, len(emptyLines))

	for num := range emptyLines {
		emptyList = append(emptyList, num)
	}

	sort.Ints(emptyList)

	for idx := range input {
		insertIdx := sort.SearchInts(emptyList, int(coordGetter(input[idx])))

		input[idx] = coordIncrease(input[idx], int64(insertIdx)*multInc-int64(insertIdx))
	}

	return input
}

func parse(multInc int64) (parsed Galaxies) {
	file, err := os.Open("day11/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var row int64

	for scanner.Scan() {
		line := scanner.Text()

		for col, item := range line {
			if SpaceType(item) == Galaxy {
				parsed = append(parsed, Coords{int64(col), row})
			}
		}

		row++
	}

	parsed = expandUniverse(multInc, parsed,
		func(coords Coords) int64 { return coords.X }, func(coords Coords, inc int64) Coords {
			return Coords{X: coords.X + inc, Y: coords.Y}
		})

	parsed = expandUniverse(multInc, parsed,
		func(coords Coords) int64 { return coords.Y }, func(coords Coords, inc int64) Coords {
			return Coords{X: coords.X, Y: coords.Y + inc}
		})

	return parsed
}

func findDistanceSum(multInc int64) string {
	galaxies := parse(multInc)

	pairs := make(GalaxyPairs)

	var sum int64

	for _, gal1 := range galaxies {
		for _, gal2 := range galaxies {
			if gal1 == gal2 || pairs.Contains(Pair{gal1, gal2}) {
				continue
			}

			pairs.Add(Pair{gal1, gal2})

			sum += gal1.distance(gal2)
		}
	}

	return fmt.Sprintf("%d", sum)
}

func Part1() string {
	return findDistanceSum(2)
}

func Part2() string {
	return findDistanceSum(1000000)
}
