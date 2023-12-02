package main

import (
	"log"

	"github.com/brunorene/adventofcode2023/day1"
	"github.com/brunorene/adventofcode2023/day2"
)

const (
	redMax   = int64(12)
	greenMax = int64(13)
	blueMax  = int64(14)
)

func main() {
	log.Println("day1 part1", day1.Part1())
	log.Println("day1 part2", day1.Part2())
	log.Println("day2 part1", day2.Part1(redMax, greenMax, blueMax))
	log.Println("day2 part2", day2.Part2())
}
