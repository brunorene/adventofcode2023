package day06

type Race struct {
	Time     int64
	Distance int64
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

func Part1() int64 {
	mult := int64(1)

	for _, race := range []Race{
		{41, 214},
		{96, 1789},
		{88, 1127},
		{94, 1055},
	} {
		mult *= wins(race)
	}

	return mult
}

func Part2() int64 {
	singleRace := Race{
		Time:     41968894,
		Distance: 214178911271055,
	}

	return wins(singleRace)
}
