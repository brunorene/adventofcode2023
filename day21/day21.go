package day21

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/brunorene/adventofcode2023/common"
)

type Coords struct {
	X, Y int64
}

type GridMap []string

func parse() (parsed GridMap, start Coords) {
	file, err := os.Open("day21/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var row int64

	for scanner.Scan() {
		line := scanner.Text()

		index := strings.Index(line, "S")

		if index >= 0 {
			start = Coords{int64(index), row}

			line = replace(line, '.', index)
		}

		parsed = append(parsed, line)

		row++
	}

	return parsed, start
}

func replace(str string, letter rune, index int) string {
	var buffer bytes.Buffer

	for idx, char := range str {
		if idx == index {
			buffer.WriteRune(letter)

			continue
		}

		buffer.WriteRune(char)
	}

	return buffer.String()
}

type CoordMap map[Coords]int64

func (cm *CoordMap) pop() (Coords, int64) {
	for k, v := range *cm {
		delete(*cm, k)

		return k, v
	}

	return Coords{}, -1
}

func (cm *CoordMap) peek() (Coords, int64) {
	for k, v := range *cm {
		return k, v
	}

	return Coords{}, -1
}

func (g GridMap) spreadCount(start Coords, target int64, infinite bool) int {
	steps := CoordMap{start: 0}

	var accum = make(CoordMap)

	for {
		currentCoords, currentDist := steps.pop()

		for _, diff := range [][2]int64{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nextX := currentCoords.X + diff[0]
			nextY := currentCoords.Y + diff[1]

			var mapPosition rune

			if infinite {
				repeatedX := nextX % int64(len(g[0]))
				repeatedY := nextY % int64(len(g))

				if repeatedX != 0 && nextX < 0 {
					repeatedX += int64(len(g[0]))
				}

				if repeatedY != 0 && nextY < 0 {
					repeatedY += int64(len(g))
				}

				mapPosition = rune(g[repeatedX][repeatedY])
			} else if nextX < 0 || nextX >= int64(len(g[0])) || nextY < 0 || nextY >= int64(len(g)) {
				continue
			} else {
				mapPosition = rune(g[nextX][nextY])
			}

			if mapPosition != '#' {
				accum[Coords{nextX, nextY}] = currentDist + 1
			}
		}

		if len(steps) == 0 {
			steps = accum

			accum = make(CoordMap)

			_, dist := steps.peek()

			if isRelevantDistance(dist) {
				fmt.Println(dist, len(steps), geometricCount((dist-65)/131))
				printStatsPerRegion(steps)
			}

			if dist == target {
				return len(steps)
			}
		}
	}
}

func isRelevantDistance(distance int64) bool {
	return (distance-65)%131 == 0
}

func printStatsPerRegion(steps CoordMap) {
	regions := make(map[Coords]int)

	for k := range steps {
		var beginX, beginY int64

		if k.X >= 0 {
			for {
				if k.X >= beginX && k.X < (beginX+131) {
					break
				}

				beginX += 131
			}
		} else {
			beginX = int64(-1)

			for {
				if k.X <= beginX && k.X > (beginX-131) {
					break
				}

				beginX -= 131
			}
		}

		if k.Y >= 0 {
			for {
				if k.Y >= beginY && k.Y < (beginY+131) {
					break
				}

				beginY += 131
			}
		} else {
			beginY = int64(-1)

			for {
				if k.Y <= beginY && k.Y > (beginY-131) {
					break
				}

				beginY -= 131
			}
		}

		regions[Coords{beginX, beginY}] += 1
	}

	mapX := make(map[int]struct{})
	mapY := make(map[int]struct{})

	var keysX, keysY []int

	for k := range regions {
		initLenX := len(mapX)
		initLenY := len(mapY)

		mapX[int(k.X)] = struct{}{}
		mapY[int(k.Y)] = struct{}{}

		if len(mapX) > initLenX {
			keysX = append(keysX, int(k.X))
		}

		if len(mapY) > initLenY {
			keysY = append(keysY, int(k.Y))
		}
	}

	sort.Ints(keysX)
	sort.Ints(keysY)

	for _, y := range keysY {
		for _, x := range keysX {
			fmt.Print(regions[Coords{int64(x), int64(y)}], "\t")
		}

		fmt.Println()
	}
}

func Part1() (sum int64) {
	grid, start := parse()

	return int64(grid.spreadCount(start, 64, false))
}

func Part2() (sum int64) {
	grid, start := parse()

	grid.spreadCount(start, 589, true)

	// 589 == 65 + (131 * 4)
	// 26501365 == 65 + (131 * 202300) N = 202300

	return geometricCount(202300)

}

func geometricCount(param int64) int64 {
	return 5612 + 5594 + 5589 + 5607 +
		param*(940+972+955+936) +
		(param-1)*(6516+6519+6498+6514) +
		param*param*7490 +
		(param-1)*(param-1)*7423
}
