package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func day1Part1() {
	file, err := os.Open("day1.txt")
	checkError(err)

	defer file.Close()

	firstRegex := regexp.MustCompile("^[a-z]*([0-9])")
	LastRegex := regexp.MustCompile("([0-9])[a-z]*$")

	var sum int

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit := firstRegex.FindStringSubmatch(line)[1]
		lastDigit := LastRegex.FindStringSubmatch(line)[1]

		number, err := strconv.Atoi(firstDigit + lastDigit)
		checkError(err)

		sum += number
	}

	fmt.Println(sum)
}

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

var numbers = []string{"1", "one", "2", "two", "3", "three", "4", "four", "5", "five", "6", "six", "7", "seven", "8", "eight", "9", "nine"}

func findFirst(line string) string {
	minIdx := len(line)
	var result string

	for _, number := range numbers {
		index := strings.Index(line, number)
		if index < minIdx && index >= 0 {
			minIdx = index
			result = number
		}
	}

	return result
}

func findLast(line string) string {
	maxIdx := -1
	var result string

	for _, number := range numbers {
		index := strings.LastIndex(line, number)
		if index > maxIdx && index >= 0 {
			maxIdx = index
			result = number
		}
	}

	return result
}

func day1Part2() {
	file, err := os.Open("day1.txt")
	checkError(err)

	defer file.Close()

	var sum int

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(convert(findFirst(line)) + convert(findLast(line)))
		checkError(err)

		fmt.Println(line, number)

		sum += number
	}

	fmt.Println(sum)
}
