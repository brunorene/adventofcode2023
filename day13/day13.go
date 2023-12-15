package day13

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/brunorene/adventofcode2023/common"
)

type Bitset struct {
	Num uint64
	Len int
}

func NewBitset(binary string) Bitset {
	binary = strings.ReplaceAll(binary, "#", "1")
	binary = strings.ReplaceAll(binary, ".", "0")

	num, err := strconv.ParseUint(binary, 2, 64)
	common.CheckError(err)

	return Bitset{
		Num: num,
		Len: len(binary),
	}
}

func (b Bitset) String() string {
	binary := strconv.FormatUint(b.Num, 2)

	return strings.Repeat("0", b.Len-len(binary)) + binary
}

type Note struct {
	Horizontal []Bitset
	Vertical   []Bitset
}

func rotate(rows []Bitset) (cols []Bitset) {
	binary := rows[0].String()

	for colIdx := range binary {
		var current string

		for rowIdx := range rows {
			current += string(rows[rowIdx].String()[colIdx])
		}

		cols = append(cols, NewBitset(current))
	}

	return cols
}

func parse() (parsed []Note) {
	file, err := os.Open("day13/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var current Note

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			current.Vertical = rotate(current.Horizontal)

			parsed = append(parsed, current)

			current = Note{}

			continue
		}

		current.Horizontal = append(current.Horizontal, NewBitset(line))
	}

	current.Vertical = rotate(current.Horizontal)

	parsed = append(parsed, current)

	return parsed
}

func reflectionIndexes(lines []Bitset, withSmudge bool) (result []int) {
	for i := 0; i < len(lines)-1; i++ {
		xorResult := lines[i].Num ^ lines[i+1].Num

		if lines[i].Num == lines[i+1].Num || (withSmudge && (xorResult&(xorResult-1)) == 0) {
			result = append(result, i)
		}
	}

	return result
}

func reflectionMatch(lines []Bitset, withSmudge bool) (result int) {
	reflectionArr := reflectionIndexes(lines, withSmudge)

	for _, reflectionIdx := range reflectionArr {
		left := slices.Clone(lines[:reflectionIdx+1])
		right := slices.Clone(lines[reflectionIdx+1:])

		slices.Reverse(left)

		smudgeFound := !withSmudge

		reflectionOk := true

		for idx := 0; idx < min(len(left), len(right)); idx++ {
			xorResult := right[idx].Num ^ left[idx].Num
			if !smudgeFound && xorResult > 0 && (xorResult&(xorResult-1)) == 0 {
				smudgeFound = true

				continue
			}

			if right[idx].Num != left[idx].Num {
				reflectionOk = false

				break
			}
		}

		if reflectionOk && smudgeFound {
			return len(left)
		}
	}

	return 0
}

func (n Note) summarize(withSmudge bool) (horz, vert int) {
	return reflectionMatch(n.Horizontal, withSmudge), reflectionMatch(n.Vertical, withSmudge)
}

func summarizeAllNotes(withSmudge bool) int {
	notes := parse()

	var sum int

	for _, note := range notes {
		horz, vert := note.summarize(withSmudge)

		sum += vert + 100*horz
	}

	return sum
}

func Part1() int {
	return summarizeAllNotes(false)
}

func Part2() int {
	return summarizeAllNotes(true)
}
