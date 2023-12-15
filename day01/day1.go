package day01

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/brunorene/adventofcode2023/common"
)

func Part1() int64 {
	file, err := os.Open("day01/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	firstRegex := regexp.MustCompile(`^[a-z]*(\d)`)
	LastRegex := regexp.MustCompile(`(\d)[a-z]*$`)

	var sum int64

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		firstDigit := firstRegex.FindStringSubmatch(line)[1]
		lastDigit := LastRegex.FindStringSubmatch(line)[1]

		number, err := strconv.Atoi(firstDigit + lastDigit)
		common.CheckError(err)

		sum += int64(number)
	}

	return sum
}

//nolint:cyclop // complexity???
func convert(number string) string {
	switch number {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return number
	}
}

func findFirst(line string) string {
	minIdx := len(line)

	var result string

	for _, number := range Numbers() {
		index := strings.Index(line, number)
		if index < minIdx && index >= 0 {
			minIdx = index
			result = number
		}
	}

	return result
}

func Numbers() []string {
	return []string{
		"1", "one", "2", "two",
		"3", "three", "4", "four",
		"5", "five", "6", "six",
		"7", "seven", "8", "eight",
		"9", "nine",
	}
}

func findLast(line string) string {
	maxIdx := -1

	var result string

	for _, number := range Numbers() {
		index := strings.LastIndex(line, number)
		if index > maxIdx && index >= 0 {
			maxIdx = index
			result = number
		}
	}

	return result
}

func Part2() int64 {
	file, err := os.Open("day01/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var sum int64

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(convert(findFirst(line)) + convert(findLast(line)))
		common.CheckError(err)

		sum += int64(number)
	}

	return sum
}
