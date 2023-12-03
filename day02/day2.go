package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/brunorene/adventofcode2023/common"
)

type Move struct {
	red   int64
	green int64
	blue  int64
}

type Game []Move

type Input []Game

func (m Move) possible(redMax, greenMax, blueMax int64) bool {
	if m.red <= redMax && m.green <= greenMax && m.blue <= blueMax {
		return true
	}

	return false
}

func parse() (out Input) {
	file, err := os.Open("day02/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.Split(line, ": ")[1]

		plays := strings.Split(line, "; ")

		var game Game

		for _, play := range plays {
			var move Move

			parts := strings.Split(play, ", ")
			for _, part := range parts {
				countAndColor := strings.Split(part, " ")

				switch countAndColor[1] {
				case "red":
					move.red, err = strconv.ParseInt(countAndColor[0], 10, 64)
					common.CheckError(err)
				case "green":
					move.green, err = strconv.ParseInt(countAndColor[0], 10, 64)
					common.CheckError(err)
				case "blue":
					move.blue, err = strconv.ParseInt(countAndColor[0], 10, 64)
					common.CheckError(err)
				}
			}

			game = append(game, move)
		}

		out = append(out, game)
	}

	return out
}

func Part1(redMax, greenMax, blueMax int64) string {
	var sum int64

	input := parse()

	idx := int64(-1)

	for {
	outerFor:
		idx++

		if idx >= int64(len(input)) {
			break
		}

		for _, move := range input[idx] {
			if !move.possible(redMax, greenMax, blueMax) {
				goto outerFor
			}
		}

		sum += idx + 1
	}

	return strconv.FormatInt(sum, 10)
}

func Part2() string {
	var sum int64

	input := parse()

	for _, game := range input {
		var maxRed, maxGreen, maxBlue int64
		for _, move := range game {
			maxRed = max(move.red, maxRed)
			maxGreen = max(move.green, maxGreen)
			maxBlue = max(move.blue, maxBlue)
		}

		sum += maxRed * maxGreen * maxBlue
	}

	return strconv.FormatInt(sum, 10)
}
