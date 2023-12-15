package day04

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/brunorene/adventofcode2023/common"
)

var (
	errPipeNotFound        = errors.New("| not found")
	errDoubleColonNotFound = errors.New(": not found")
)

type card struct {
	count   int
	winning map[int64]struct{}
	actual  []int64
}

func (c card) points() int64 {
	double := int64(0)

	for _, num := range c.actual {
		_, exists := c.winning[num]
		if !exists {
			continue
		}

		if double == 0 {
			double = 1

			continue
		}

		double *= 2
	}

	return double
}

func (c card) winCount() int {
	count := 0

	for _, num := range c.actual {
		_, exists := c.winning[num]
		if !exists {
			continue
		}

		count++
	}

	return count
}

func parse() (parsed []card) {
	file, err := os.Open("day04/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	matcher := regexp.MustCompile(`\d+`)

	for scanner.Scan() {
		line := scanner.Text()

		winning, actual, found := strings.Cut(line, "|")

		if !found {
			common.CheckError(errPipeNotFound)
		}

		_, winning, found = strings.Cut(winning, ":")

		if !found {
			common.CheckError(errDoubleColonNotFound)
		}

		winningStrList := matcher.FindAllString(winning, -1)
		actualStrList := matcher.FindAllString(actual, -1)

		newItem := card{
			count:   1,
			actual:  make([]int64, 0, len(actualStrList)),
			winning: make(map[int64]struct{}, len(winningStrList)),
		}

		for _, numStr := range winningStrList {
			num, err := strconv.Atoi(numStr)
			common.CheckError(err)

			newItem.winning[int64(num)] = struct{}{}
		}

		for _, numStr := range actualStrList {
			num, err := strconv.Atoi(numStr)
			common.CheckError(err)

			newItem.actual = append(newItem.actual, int64(num))
		}

		parsed = append(parsed, newItem)
	}

	return parsed
}

func Part1() int64 {
	cards := parse()

	var sum int64

	for _, card := range cards {
		sum += card.points()
	}

	return sum
}

func Part2() int64 {
	cards := parse()

	for idx, current := range cards {
		count := current.winCount()

		for i := idx + 1; i < idx+1+count && i < len(cards); i++ {
			cards[i].count += current.count
		}
	}

	var sum int64

	for _, current := range cards {
		sum += int64(current.count)
	}

	return sum
}
