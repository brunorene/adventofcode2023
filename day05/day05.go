package day05

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/brunorene/adventofcode2023/common"
)

const stepLen = 7

type valRange struct {
	sourceMin, sourceMax, destination int64
}

type step []valRange

type almanac struct {
	seeds []int64
	steps []step
}

func (s step) convert(source int64) int64 {
	for _, vRange := range s {
		if source >= vRange.sourceMin && source <= vRange.sourceMax {
			return vRange.destination + (source - vRange.sourceMin)
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

		parsed.steps[len(parsed.steps)-1] = append(parsed.steps[len(parsed.steps)-1], valRange{
			sourceMin:   nums[1],
			sourceMax:   nums[1] + nums[2] - 1,
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
