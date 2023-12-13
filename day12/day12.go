package day12

import (
	"bufio"
	"hash/maphash"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/brunorene/adventofcode2023/common"
)

type Spring struct {
	Condition    Matcher
	Distribution []int
	Length       int
}

type Matcher string

var (
	memoizeStore sync.Map
	hashSeed     = maphash.MakeSeed()
)

func newMemoizeKey(starter [][]int, spaceCount int) uint64 {
	source := make([]byte, 0, len(starter)+2)

	source = append(source, byte(spaceCount))

	for idx := range starter {
		for _, num := range starter[idx] {
			source = append(source, byte(num))
		}
	}

	return maphash.Bytes(hashSeed, source)
}

func (m Matcher) matches(str string) (result bool) {
	if len(m) != len(str) {
		return false
	}

	for idx, item := range m {
		if item != '?' && str[idx] != m[idx] {
			return false
		}
	}

	return true
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
			Condition:    Matcher(matcher),
			Distribution: distribution,
			Length:       len(parts[0])*repeat + repeat - 1,
		})
	}

	return parsed
}

func sum(slice []int) (result int) {
	for _, num := range slice {
		result += num
	}

	return result
}

func combinations(starter [][]int, spaceCount int) (result [][]int) {
	if len(starter) == 0 {
		return [][]int{{}}
	}

	if len(starter) == 1 {
		for _, num := range starter[0] {
			if num == spaceCount {
				return [][]int{{num}}
			}
		}

		return [][]int{{}}
	}

	innerStarter := make([][]int, len(starter))

	for arrIdx, arr := range starter {
		for numIdx, num := range arr {
			if num == spaceCount && numIdx < len(starter[arrIdx])-1 {
				innerStarter[arrIdx] = starter[arrIdx][:numIdx+1]

				break
			} else {
				innerStarter[arrIdx] = starter[arrIdx]
			}
		}
	}

	// key := newMemoizeKey(innerStarter, spaceCount)
	//
	// if cache, exists := memoizeStore.Load(key); exists {
	// 	return cache.([][]int)
	// }

	for _, num := range starter[0] {
		suffixes := combinations(innerStarter[1:], spaceCount-num)

		for _, suffix := range suffixes {
			if num+sum(suffix) != spaceCount {
				continue
			}

			result = append(result, slices.Insert(suffix, 0, num))
		}
	}

	// memoizeStore.Store(key, result)

	return result
}

func createCandidate(distribution, spaces []int) (result string) {
	for idx := range distribution {
		result += strings.Repeat(".", spaces[idx]) + strings.Repeat("#", distribution[idx])
	}

	return result + strings.Repeat(".", spaces[len(spaces)-1])
}

func generateSeq(max int, withZero bool) (seq []int) {
	if withZero {
		seq = append(seq, 0)
	}

	for i := 1; i <= max; i++ {
		seq = append(seq, i)
	}

	return seq
}

func (s Spring) generator() []string {
	springCount := sum(s.Distribution)

	spaceCount := s.Length - springCount

	maxSpace := spaceCount - (len(s.Distribution) - 2)

	spaces := make([][]int, 0, len(s.Distribution)-1)

	for i := 0; i <= len(s.Distribution); i++ {
		spaces = append(spaces, generateSeq(maxSpace, i == 0 || i == len(s.Distribution)))
	}

	var candidates []string

	for _, comb := range combinations(spaces, spaceCount) {
		candidate := createCandidate(s.Distribution, comb)

		if s.Condition.matches(candidate) {
			candidates = append(candidates, candidate)
		}
	}

	return candidates
}

func totalCandidates(repeat int) string {
	springs := parse(repeat)

	var result int

	for _, spring := range springs {
		candidates := spring.generator()

		result += len(candidates)
	}

	return strconv.FormatInt(int64(result), 10)
}

func Part1() string {
	return totalCandidates(1)
}

func Part2() string {
	return totalCandidates(5)
}
