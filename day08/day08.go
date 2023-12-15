package day08

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/brunorene/adventofcode2023/common"
)

const (
	Left  = 'L'
	Right = 'R'
)

type Node struct {
	Name     string
	Junction map[rune]string
}

type DesertMap struct {
	current    int
	Directions string
	Nodes      map[string]Node
}

func (d *DesertMap) nextDirection() rune {
	next := rune(d.Directions[d.current])
	d.current = (d.current + 1) % len(d.Directions)

	return next
}

func (d *DesertMap) stepsFromTo(start Node, endsWith string) int64 {
	var steps int64

	currentNode := start

	for {
		steps++

		nextDirection := d.nextDirection()
		nextNode := currentNode.Junction[nextDirection]

		if strings.HasSuffix(nextNode, endsWith) {
			return steps
		}

		currentNode = d.Nodes[nextNode]
	}
}

func parse() (parsed DesertMap) {
	file, err := os.Open("day08/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	directions := scanner.Text()

	parsed = DesertMap{
		Directions: directions,
		Nodes:      make(map[string]Node),
	}

	scanner.Scan()
	scanner.Text() // newline

	matcher := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

	for scanner.Scan() {
		line := scanner.Text()

		matches := matcher.FindStringSubmatch(line)

		parsed.Nodes[matches[1]] = Node{
			Name: matches[1],
			Junction: map[rune]string{
				Left:  matches[2],
				Right: matches[3],
			},
		}
	}

	return parsed
}

func Part1() int64 {
	dMap := parse()

	return dMap.stepsFromTo(dMap.Nodes["AAA"], "ZZZ")
}

func Part2() int64 {
	dMap := parse()

	var steps []int64

	for k, n := range dMap.Nodes {
		if k[2] == 'A' {
			steps = append(steps, dMap.stepsFromTo(n, "Z"))
		}
	}

	return LCM(steps[0], steps[1], steps[2:]...)
}

func GCD(valA, valB int64) int64 {
	for valB != 0 {
		t := valB
		valB = valA % valB
		valA = t
	}

	return valA
}

func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
