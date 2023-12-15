package day14

import (
	"bufio"
	"hash/maphash"
	"os"
	"regexp"
	"slices"

	"github.com/brunorene/adventofcode2023/common"
)

const (
	Round = 'O'
	Cube  = '#'
	Empty = '.'
)

//nolint:gochecknoglobals // its a seed, it needs to be global
var Seed = maphash.MakeSeed()

type Terrain []string

func parse() (parsed Terrain) {
	file, err := os.Open("day14/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		parsed = append(parsed, line)
	}

	return parsed
}

func replaceAtIndex(input string, char rune, idx int) string {
	out := []rune(input)
	out[idx] = char

	return string(out)
}

func (t Terrain) rockRoll() {
	for col := range t[0] {
		northIdx := 0

		for row, line := range t {
			switch line[col] {
			case Round:
				if northIdx < row {
					t[northIdx] = replaceAtIndex(t[northIdx], Round, col)
					t[row] = replaceAtIndex(line, Empty, col)
				}

				northIdx++
			case Cube:
				northIdx = row + 1
			}
		}
	}
}

func (t Terrain) transposeClockwise() (output Terrain) {
	output = make(Terrain, len(t[0]))

	reversed := slices.Clone(t)

	slices.Reverse(reversed)

	for xIdx := range reversed[0] {
		for yIdx := range reversed {
			output[yIdx] += string(reversed[xIdx][yIdx])
		}
	}

	return output
}

func (t Terrain) cycle() Terrain {
	copyT := slices.Clone(t)

	for i := 0; i < 4; i++ {
		copyT.rockRoll()
		copyT = copyT.transposeClockwise()
	}

	return copyT
}

func (t Terrain) load() int {
	nonRound := regexp.MustCompile(`[.#]`)

	var sum int

	for i := 0; i < len(t); i++ {
		score := len(t) - i

		sum += len(nonRound.ReplaceAllString(t[i], "")) * score
	}

	return sum
}

func Part1() int {
	terrain := parse()

	terrain.rockRoll()

	return terrain.load()
}

func (t Terrain) Hash() uint64 {
	var all string

	for _, line := range t {
		all += line
	}

	return maphash.String(Seed, all)
}

func Part2() int {
	terrain := parse()

	hits := make(map[uint64][]int)

	var loads []int

	var startLoop int

	for i := 0; i < 500; i++ {
		terrain = terrain.cycle()

		hash := terrain.Hash()

		hits[hash] = append(hits[hash], i)
		loads = append(loads, terrain.load())

		if len(hits[hash]) > 1 {
			startLoop = hits[hash][0]

			break
		}
	}

	endloop := len(hits)

	maxCycles := 1000000000

	index := (startLoop - 1) + (maxCycles-startLoop)%(endloop-startLoop)

	return loads[index]
}
