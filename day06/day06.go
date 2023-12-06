package day06

import (
	"strconv"
)

type Race struct {
	Time     int64
	Distance int64
}

func getRaces() []Race {
	return []Race{
		{41, 214},
		{96, 1789},
		{88, 1127},
		{94, 1055},
	}
}

func wins(race Race) (count int64) {
	for button := int64(0); button <= race.Time; button++ {
		travel := button * (race.Time - button)

		if travel > race.Distance {
			count++
		}
	}

	return count
}

func Part1() string {
	mult := int64(1)

	for _, race := range getRaces() {
		mult *= wins(race)
	}

	return strconv.FormatInt(mult, 10)
}
func Part2() string {
	singleRace := Race{
		Time:     41968894,
		Distance: 214178911271055,
	}

	return strconv.FormatInt(wins(singleRace), 10)
}
