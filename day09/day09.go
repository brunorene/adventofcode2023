package day09

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/brunorene/adventofcode2023/common"
)

func process(handler func([]int64) int64) (sum int64) {
	file, err := os.Open("day09/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " ")

		var measure []int64

		for _, val := range values {
			num, err := strconv.ParseInt(val, 10, 64)
			common.CheckError(err)

			measure = append(measure, num)
		}

		sum += handler(measure)
	}

	return sum
}

func sumAll(values []int64) (result int64) {
	for _, value := range values {
		result += value
	}

	return result
}

func createPyramid(values []int64) (totals [][]int64) {
	totals = [][]int64{values}

	for sumAll(totals[len(totals)-1]) != 0 {
		var results []int64

		for i := 1; i < len(totals[len(totals)-1]); i++ {
			results = append(results, totals[len(totals)-1][i]-totals[len(totals)-1][i-1])
		}

		totals = append(totals, results)
	}

	return totals
}

func processMeasure(values []int64) int64 {
	pyramid := createPyramid(values)

	for i := len(pyramid) - 1; i > 0; i-- {
		pyramid[i-1] = append(pyramid[i-1],
			pyramid[i][len(pyramid[i])-1]+pyramid[i-1][len(pyramid[i-1])-1])
	}

	return pyramid[0][len(pyramid[0])-1]
}

func processBackwardMeasure(values []int64) int64 {
	pyramid := createPyramid(values)

	for i := len(pyramid) - 1; i > 0; i-- {
		pyramid[i-1] = append(
			[]int64{pyramid[i-1][0] - pyramid[i][0]}, pyramid[i-1]...)
	}

	return pyramid[0][0]
}

func Part1() string {
	return strconv.FormatInt(process(processMeasure), 10)
}

func Part2() string {
	return strconv.FormatInt(process(processBackwardMeasure), 10)
}
