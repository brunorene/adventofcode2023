package main

import (
	"log"

	"github.com/brunorene/adventofcode2023/day01"
	"github.com/brunorene/adventofcode2023/day02"
	"github.com/brunorene/adventofcode2023/day03"
)

const (
	redMax   = int64(12)
	greenMax = int64(13)
	blueMax  = int64(14)
)

func main() {
	log.Println("day1 part1", day01.Part1())
	log.Println("day1 part2", day01.Part2())
	log.Println("day2 part1", day02.Part1(redMax, greenMax, blueMax))
	log.Println("day2 part2", day02.Part2())
	log.Println("day3 part1", day03.Part1())
	log.Println("day3 part2", day03.Part2())
}
