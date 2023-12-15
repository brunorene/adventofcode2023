package day07

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/brunorene/adventofcode2023/common"
)

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
	Cards1 = "23456789TJQKA"
	Cards2 = "J23456789TQKA"
)

type HandType int

type Hand string

type Play struct {
	Hand Hand
	Bid  int64
}

func (h Hand) counter() map[rune]int {
	result := make(map[rune]int)

	for _, card := range h {
		result[card] += 1
	}

	return result
}

func getMax(content map[rune]int) int {
	maxCount := 0

	for _, count := range content {
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}

func GetHandType(hand Hand) HandType {
	result := hand.counter()

	switch len(result) {
	case 1:
		return FiveOfAKind
	case 2:
		if getMax(result) == 4 {
			return FourOfAKind
		}

		return FullHouse
	case 3:
		if getMax(result) == 3 {
			return ThreeOfAKind
		}

		return TwoPair
	case 4:
		return OnePair
	default:
		return HighCard
	}
}

func GetHandTypeWithJoker(hand Hand) HandType {
	if !strings.ContainsRune(string(hand), 'J') {
		return GetHandType(hand)
	}

	maxType := GetHandType(hand)

	for _, card := range Cards2 {
		newHand := Hand(strings.ReplaceAll(string(hand), "J", string(card)))

		currentType := GetHandType(newHand)

		if currentType > maxType {
			maxType = currentType
		}
	}

	return maxType
}

func SortPlays(input []Play, cardScore string, handTypeGetter func(Hand) HandType) []Play {
	sort.Slice(input, func(i, j int) bool {
		hand1 := input[i].Hand
		hand2 := input[j].Hand

		type1 := handTypeGetter(hand1)
		type2 := handTypeGetter(hand2)

		if type1 < type2 {
			return true
		}

		if type1 == type2 {
			for i := 0; i < 5; i++ {
				value1 := strings.IndexRune(cardScore, rune(hand1[i]))
				value2 := strings.IndexRune(cardScore, rune(hand2[i]))

				if value1 < value2 {
					return true
				}

				if value1 > value2 {
					return false
				}
			}
		}

		return false
	})

	return input
}

func parse() (parsed []Play) {
	file, err := os.Open("day07/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	parsed = make([]Play, 0, 1000)

	for scanner.Scan() {
		line := scanner.Text()

		pair := strings.Split(line, " ")

		bid, err := strconv.ParseInt(pair[1], 10, 64)
		common.CheckError(err)

		parsed = append(parsed, Play{Hand(pair[0]), bid})
	}

	return parsed
}

func Part1() int64 {
	parsed := parse()

	sorted := SortPlays(parsed, Cards1, GetHandType)

	var sum int64

	for idx, play := range sorted {
		sum += int64(idx+1) * play.Bid
	}

	return sum
}

func Part2() int64 {
	parsed := parse()

	sorted := SortPlays(parsed, Cards2, GetHandTypeWithJoker)

	var sum int64

	for idx, play := range sorted {
		sum += int64(idx+1) * play.Bid
	}

	return sum
}
