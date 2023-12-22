package day17

import (
	"bufio"
	"os"
	"sort"
	"strconv"

	"github.com/brunorene/adventofcode2023/common"
)

type Direction struct {
	DX, DY int
}

//nolint:gochecknoglobals // it is a var, but they do not change
var (
	Right         = Direction{1, 0}
	Left          = Direction{-1, 0}
	Up            = Direction{0, -1}
	Down          = Direction{0, 1}
	None          = Direction{0, 0}
	AllDirections = []Direction{Left, Right, Up, Down}
)

type Layout [][]int

func (l Layout) Get(coords Coords) int {
	return l[coords.Y][coords.X]
}

func (l Layout) Goal() Coords {
	return Coords{len(l[0]) - 1, len(l) - 1}
}

type Coords struct {
	X, Y int
}

type Node struct {
	Coords
	Direction
	DirCount int
}
type QueueNode struct {
	Cost int
	Node
}

type PriorityQueue []QueueNode

func (p *PriorityQueue) Push(qnode QueueNode) {
	*p = append(*p, qnode)

	sort.Slice(*p, func(i, j int) bool {
		return (*p)[i].Cost < (*p)[j].Cost
	})
}

func (p *PriorityQueue) Pop() QueueNode {
	item := (*p)[0]
	*p = (*p)[1:]

	return item
}

func parse() (parsed Layout) {
	file, err := os.Open("day17/input_test")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		sliceLine := make([]int, 0, len(line))

		for _, char := range line {
			num, err := strconv.ParseInt(string(char), 10, 32)
			common.CheckError(err)

			sliceLine = append(sliceLine, int(num))
		}

		parsed = append(parsed, sliceLine)
	}

	return parsed
}

func Directions(direction Direction, straight, maxStraight int) (result []Direction) {
	if direction == None {
		return []Direction{Down, Right}
	}

	if straight < maxStraight {
		result = append(result, direction)
	}

	for _, nextDir := range AllDirections {
		if (direction.DX+nextDir.DX == 0 && direction.DY+nextDir.DY == 0) || direction == nextDir {
			continue
		}

		result = append(result, nextDir)
	}

	return result
}

func (l Layout) Neighbours(node Node, maxStraight int) (result []Node) {
	for _, nextDir := range Directions(node.Direction, node.DirCount, maxStraight) {
		nextCoords := Coords{node.X + nextDir.DX, node.Y + nextDir.DY}

		if nextCoords.X < 0 || nextCoords.Y < 0 || nextCoords.X >= len(l[0]) || nextCoords.Y >= len(l) {
			continue
		}

		nextDirCount := 1
		if nextDir == node.Direction {
			nextDirCount++
		}

		result = append(result, Node{
			Coords:    nextCoords,
			Direction: nextDir,
			DirCount:  nextDirCount,
		})
	}

	return result
}

func (l Layout) dijkstra(minStraight, maxStraight int) int {
	visited := make(map[Node]struct{})

	pQueue := PriorityQueue{}

	pQueue.Push(QueueNode{
		Cost: 0,
		Node: Node{
			Coords:    Coords{0, 0},
			Direction: None,
			DirCount:  0,
		},
	})

	for len(pQueue) > 0 {
		node := pQueue.Pop()

		if node.DirCount >= minStraight && node.Coords == l.Goal() {
			return node.Cost
		}

		if _, exists := visited[node.Node]; exists {
			continue
		}

		visited[node.Node] = struct{}{}

		neighbours := l.Neighbours(node.Node, maxStraight)

		for _, nextNode := range neighbours {
			nextCost := node.Cost + l.Get(nextNode.Coords)

			pQueue.Push(QueueNode{Cost: nextCost, Node: nextNode})
		}
	}

	return 0
}

func Part1() int {
	return parse().dijkstra(1, 3)
}
