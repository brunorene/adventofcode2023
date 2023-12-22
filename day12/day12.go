package day12

import (
	"bufio"
	"bytes"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/brunorene/adventofcode2023/common"
)

type Spring struct {
	Condition    string
	Distribution []int
	Length       int
}

func parse(repeat int) (parsed []Spring) {
	file, err := os.Open("day12/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == ','
		})

		distribution := make([]int, 0, repeat*(len(parts)-1))

		for i := 0; i < repeat; i++ {
			for _, part := range parts[1:] {
				num, err := strconv.Atoi(part)
				common.CheckError(err)

				distribution = append(distribution, num)
			}
		}

		matcher, _ := strings.CutSuffix(strings.Repeat(parts[0]+"?", repeat), "?")

		parsed = append(parsed, Spring{
			Condition:    matcher,
			Distribution: distribution,
			Length:       len(parts[0])*repeat + repeat - 1,
		})
	}

	return parsed
}

func (s Spring) unknownRangeEnd() (end int64) {
	unknownLen := len(strings.ReplaceAll(strings.ReplaceAll(s.Condition, "#", ""), ".", ""))

	return int64(math.Pow(2, float64(unknownLen)) - 1)
}

const (
	Space = iota
	OnSpring
)

func (s Spring) validateCandidate(empty int64) (result string) {
	var emptyIdx int

	var line bytes.Buffer

	var springLen int

	var currDist int

	state := Space

	for idx := 0; idx < s.Length; idx++ {
		if s.Condition[idx] == '?' {
			emptyIdx++

			if empty&int64(math.Pow(2, float64(emptyIdx))) == 0 {
				line.WriteRune('.')

				if state == OnSpring {
					currDist++
				}

				state = Space

				continue
			}

			line.WriteRune('#')

			state = OnSpring

			springLen++

			continue
		}

		line.WriteRune(rune(s.Condition[idx]))

		if s.Condition[idx] == '.' {
			state = Space
		} else {
			state = OnSpring

			springLen++
		}
	}

	return line.String()
}

func (s Spring) generator() []string {
	var candidates []string

	for empty := int64(0); empty <= s.unknownRangeEnd(); empty++ {
		if candidate := s.validateCandidate(empty); candidate != "" {
			candidates = append(candidates, candidate)
		}
	}

	return candidates
}

func totalCandidates(repeat int) int {
	springs := parse(repeat)

	var result int

	for _, spring := range springs {
		candidates := spring.generator()

		result += len(candidates)
	}

	return result
}

func Part1() int {
	return totalCandidates(1)
}

func Part2() int {
	return totalCandidates(5)
}
