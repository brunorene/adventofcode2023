package day05

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/brunorene/adventofcode2023/common"
)

const (
	stepLen        = 7
	ignoreFirstIdx = 2
)

type ValRange struct {
	SourceStart, SourceEnd, destination int64
}

type step []ValRange

type almanac struct {
	seeds []int64
	steps []step
}

func (s step) convert(source int64) int64 {
	for _, vRange := range s {
		if source >= vRange.SourceStart && source <= vRange.SourceEnd {
			return vRange.destination + (source - vRange.SourceStart)
		}
	}

	return source
}

func parse() (parsed almanac) {
	file, err := os.Open("day05/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	seedsStr := strings.Split(strings.ReplaceAll(scanner.Text(), "seeds: ", ""), " ")

	for _, seedStr := range seedsStr {
		conv, err := strconv.ParseInt(seedStr, 10, 64)
		common.CheckError(err)

		parsed.seeds = append(parsed.seeds, conv)
	}

	parsed.steps = make([]step, 0, stepLen)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			scanner.Scan()
			scanner.Text()

			parsed.steps = append(parsed.steps, step{})

			continue
		}

		parts := strings.Split(line, " ")

		var nums []int64

		for _, part := range parts {
			num, err := strconv.ParseInt(part, 10, 64)
			common.CheckError(err)

			nums = append(nums, num)
		}

		parsed.steps[len(parsed.steps)-1] = append(parsed.steps[len(parsed.steps)-1], ValRange{
			SourceStart: nums[1],
			SourceEnd:   nums[1] + nums[2] - 1,
			destination: nums[0],
		})
	}

	return parsed
}

func (a almanac) seedToLocation(seed int64) int64 {
	current := seed

	for idx := range a.steps {
		current = a.steps[idx].convert(current)
	}

	return current
}

func Part1() string {
	parsed := parse()

	minimum := int64(math.MaxInt64)

	for _, seed := range parsed.seeds {
		location := parsed.seedToLocation(seed)

		if location < minimum {
			minimum = location
		}
	}

	return strconv.FormatInt(minimum, 10)
}

func SplitRanges(start, end int64, ranges []ValRange) (result []ValRange) {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].SourceStart < ranges[j].SourceStart
	})

	for idx := range ranges {
		if start > end {
			break
		}

		if start > ranges[idx].SourceEnd {
			continue
		}

		// before
		if start < ranges[idx].SourceStart {
			if end < ranges[idx].SourceStart {
				break
			}

			result = append(result, ValRange{SourceStart: start, SourceEnd: ranges[idx].SourceStart - 1})

			start = ranges[idx].SourceStart
		}

		// during
		if start >= ranges[idx].SourceStart {
			if end <= ranges[idx].SourceEnd {
				break
			}

			result = append(result, ValRange{SourceStart: start, SourceEnd: ranges[idx].SourceEnd})

			start = ranges[idx].SourceEnd + 1
		}
	}

	if start <= end {
		result = append(result, ValRange{SourceStart: start, SourceEnd: end})
	}

	return result
}

func MergeRanges(input []ValRange) []ValRange {
	if len(input) <= 1 {
		return input
	}

	sort.Slice(input, func(i, j int) bool {
		return input[i].SourceStart < input[j].SourceStart
	})

	idx := 0

	for idx < len(input)-1 {
		if input[idx].SourceEnd == input[idx+1].SourceStart-1 {
			inner := make([]ValRange, 0, len(input)-1)
			inner = append(inner, input[0:idx]...)
			inner = append(inner, ValRange{SourceStart: input[idx].SourceStart, SourceEnd: input[idx+1].SourceEnd})

			for i := idx + ignoreFirstIdx; i < len(input); i++ {
				inner = append(inner, input[i])
			}

			input = inner

			continue
		}

		idx++
	}

	return input
}

func (a almanac) seedsToLocations(start, end int64) (result []ValRange) {
	currentRanges := []ValRange{{SourceStart: start, SourceEnd: end}}

	for idx := range a.steps {
		var rangesOut []ValRange

		for _, currentRangeIn := range currentRanges {
			rangesIn := SplitRanges(currentRangeIn.SourceStart, currentRangeIn.SourceEnd, a.steps[idx])

			for _, currentRangeOut := range rangesIn {
				rangesOut = append(rangesOut, ValRange{
					SourceStart: a.steps[idx].convert(currentRangeOut.SourceStart),
					SourceEnd:   a.steps[idx].convert(currentRangeOut.SourceEnd),
				})
			}
		}

		currentRanges = MergeRanges(rangesOut)
	}

	return currentRanges
}

func Part2() string {
	parsed := parse()

	minimum := int64(math.MaxInt64)

	for i := 0; i < len(parsed.seeds); i += 2 {
		locations := parsed.seedsToLocations(parsed.seeds[i], parsed.seeds[i]+parsed.seeds[i+1]-1)

		for _, location := range locations {
			if location.SourceStart < minimum {
				minimum = location.SourceStart
			}
		}
	}

	return strconv.FormatInt(minimum, 10)
}
