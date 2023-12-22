package day18

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/brunorene/adventofcode2023/common"
)

type Coords struct {
	X, Y int64
}

type Border struct {
	Coords
	Length int64
}

type DigPlan []Border

func parse(fromColour bool) (parsed DigPlan) {
	file, err := os.Open("day18/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var current Coords

	parsed = append(parsed, Border{Coords: Coords{}})

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		colour := strings.Trim(parts[2], "()#")

		meters, err := strconv.ParseInt(parts[1], 10, 32)
		common.CheckError(err)

		direction := parts[0]

		if fromColour {
			meters, err = strconv.ParseInt(colour[:len(colour)-1], 16, 32)
			common.CheckError(err)

			direction = string(colour[len(colour)-1])
		}

		directions := map[string]Coords{
			"U": {0, -1}, "3": {0, -1},
			"D": {0, 1}, "1": {0, 1},
			"L": {-1, 0}, "2": {-1, 0},
			"R": {1, 0}, "0": {1, 0},
		}

		parsed = append(parsed, Border{
			Coords: Coords{current.X + directions[direction].X*meters, current.Y + directions[direction].Y*meters},
			Length: meters,
		})

		current = parsed[len(parsed)-1].Coords
	}

	return parsed
}

func (dp DigPlan) Perimeter() (result int64) {
	for _, border := range dp {
		result += border.Length
	}

	return result
}

func (dp DigPlan) shoelaceArea() (area int64) {
	for idx := 0; idx < len(dp)-1; idx++ {
		area += dp[idx].X*dp[idx+1].Y - dp[idx].Y*dp[idx+1].X + dp[idx+1].Length
	}

	return area / 2
}

func (dp DigPlan) cubicMetersUsingPick() int64 {
	return dp.shoelaceArea() + 1
}

func Part1() int64 {
	return parse(false).cubicMetersUsingPick()
}

func Part2() int64 {
	return parse(true).cubicMetersUsingPick()
}
