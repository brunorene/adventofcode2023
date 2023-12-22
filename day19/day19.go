package day19

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/brunorene/adventofcode2023/common"
)

type (
	Xmas     string
	Operator string
	Outcome  string
)

const (
	KeyX        Xmas     = "x"
	KeyM        Xmas     = "m"
	KeyA        Xmas     = "a"
	KeyS        Xmas     = "s"
	KeyNone     Xmas     = ""
	GreaterThan Operator = ">"
	LessThan    Operator = "<"
	NoneOp      Operator = ""
	Accept      Outcome  = "Accept"
	Reject      Outcome  = "Reject"
	Link        Outcome  = "Link"
	Begin                = "in"
)

type Limit struct {
	Key        Xmas
	MinExclMax [2]int64
}

type Rule struct {
	Key      Xmas
	Operator Operator
	Limit    int64
	Outcome  Outcome
	Link     string
}

type Part map[Xmas]int64

type System struct {
	Rules map[string][]Rule
	Parts []Part
}

var (
	RuleMatcher = regexp.MustCompile(`([xmas])([><])(\d+):(\w+)`)
	PartMatcher = regexp.MustCompile(`\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)
)

// mk{s<2413:A,x<3330:A,x>3427:A,A}.
func processRules(line string) (name string, rules []Rule) {
	parts := strings.Split(line, "{")

	name = parts[0]

	ruleParts := strings.Split(strings.Trim(parts[1], "}"), ",")

	for _, rulePart := range ruleParts {
		components := RuleMatcher.FindStringSubmatch(rulePart)

		if components == nil {
			switch rulePart {
			case "A":
				rules = append(rules, Rule{
					Key:      KeyNone,
					Operator: NoneOp,
					Outcome:  Accept,
					Link:     "",
				})
			case "R":
				rules = append(rules, Rule{
					Key:      KeyNone,
					Operator: NoneOp,
					Outcome:  Reject,
					Link:     "",
				})
			default:
				rules = append(rules, Rule{
					Key:      KeyNone,
					Operator: NoneOp,
					Outcome:  Link,
					Link:     rulePart,
				})
			}

			continue
		}

		limit, err := strconv.ParseInt(components[3], 10, 64)
		common.CheckError(err)

		newRule := Rule{
			Key:      Xmas(rune(components[1][0])),
			Operator: Operator(rune(components[2][0])),
			Limit:    limit,
		}

		switch components[4] {
		case "A":
			newRule.Outcome = Accept
		case "R":
			newRule.Outcome = Reject
		default:
			newRule.Outcome = Link
			newRule.Link = components[4]
		}

		rules = append(rules, newRule)
	}

	return name, rules
}

func processPart(line string) (result Part) {
	parts := PartMatcher.FindStringSubmatch(line)

	if len(parts) < 5 {
		log.Fatalf("unexpected part: %s", line)
	}

	result = make(Part)

	for idx, key := range []Xmas{KeyX, KeyM, KeyA, KeyS} {
		value, err := strconv.ParseInt(parts[idx+1], 10, 64)
		common.CheckError(err)

		result[key] = value
	}

	return result
}

func parse() (parsed System) {
	file, err := os.Open("day19/input")
	common.CheckError(err)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	processingRules := true

	parsed.Rules = make(map[string][]Rule)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			processingRules = false

			continue
		}

		if processingRules {
			name, rules := processRules(line)

			parsed.Rules[name] = rules

			continue
		}

		parsed.Parts = append(parsed.Parts, processPart(line))
	}

	return parsed
}

func (r Rule) isValid(part Part) bool {
	if r.Operator == LessThan {
		return part[r.Key] < r.Limit
	}

	if r.Operator == GreaterThan {
		return part[r.Key] > r.Limit
	}

	log.Fatalf("unexpected operator in %v", r)

	return false
}

func getOutcome(part Part, rules map[string][]Rule) Outcome {
	currentRule := Begin

	for {
		rule := rules[currentRule]

		for _, rulePart := range rule {
			if rulePart.Key == KeyNone || rulePart.isValid(part) {
				if slices.Contains([]Outcome{Accept, Reject}, rulePart.Outcome) {
					return rulePart.Outcome
				}

				currentRule = rulePart.Link

				break
			}
		}
	}
}

func Part1() (sum int64) {
	system := parse()

	for _, part := range system.Parts {
		outcome := getOutcome(part, system.Rules)

		if outcome == Accept {
			sum += part[KeyX] + part[KeyM] + part[KeyA] + part[KeyS]

			continue
		}

		if outcome != Reject {
			log.Fatalf("part %v unknown outcome", part)
		}
	}

	return sum
}

type Node struct {
	Key           Xmas
	LessThan      int64
	Outcome       Outcome
	LessThanTrue  *Node
	LessThanFalse *Node
}

func (s System) createTree(name string, index int) (root *Node) {
	rule := s.Rules[name][index]

	setVertex := func(node *Node) {
		root.LessThanTrue = node
		root.LessThanFalse = s.createTree(name, index+1)
	}

	switch rule.Operator {
	case NoneOp:
		if rule.Outcome != Link {
			return &Node{Key: rule.Key, Outcome: rule.Outcome}
		}

		return s.createTree(rule.Link, 0)
	case LessThan:
		root = &Node{Key: rule.Key, LessThan: rule.Limit}
	case GreaterThan:
		root = &Node{Key: rule.Key, LessThan: rule.Limit + 1}

		setVertex = func(node *Node) {
			root.LessThanFalse = node
			root.LessThanTrue = s.createTree(name, index+1)
		}
	}

	switch rule.Outcome {
	case Accept:
		setVertex(&Node{Outcome: Accept})
	case Reject:
		setVertex(&Node{Outcome: Reject})
	case Link:
		setVertex(s.createTree(rule.Link, 0))
	default:
		log.Fatalf("unknown outcome: %v", rule.Outcome)
	}

	return root
}

func (o Outcome) startPart() map[Xmas][2]int64 {
	ending := [2]int64{0, 0}

	if o == Accept {
		ending = [2]int64{0, 1}
	}

	return map[Xmas][2]int64{
		KeyX:    {1, 4001},
		KeyM:    {1, 4001},
		KeyA:    {1, 4001},
		KeyS:    {1, 4001},
		KeyNone: ending,
	}
}

func numberOfCombinations(ranges []map[Xmas][2]int64) (result int64) {
	for _, line := range ranges {
		result += (line[KeyX][1] - line[KeyX][0]) *
			(line[KeyM][1] - line[KeyM][0]) *
			(line[KeyA][1] - line[KeyA][0]) *
			(line[KeyS][1] - line[KeyS][0]) *
			(line[KeyNone][1] - line[KeyNone][0])
	}

	return result
}

func combinations(root *Node) (result []map[Xmas][2]int64) {
	if root.Outcome != "" {
		return []map[Xmas][2]int64{root.Outcome.startPart()}
	}

	listLt := combinations(root.LessThanTrue)
	listGte := combinations(root.LessThanFalse)

	for _, partialLt := range listLt {
		next, exists := partialLt[root.Key]

		if exists {
			partialLt[root.Key] = [2]int64{max(1, next[0]), min(root.LessThan, next[1])}

			continue
		}

		partialLt[root.Key] = [2]int64{1, root.LessThan}
	}

	for _, partialGte := range listGte {
		next, exists := partialGte[root.Key]

		if exists {
			partialGte[root.Key] = [2]int64{max(root.LessThan, next[0]), min(4001, next[1])}

			continue
		}

		partialGte[root.Key] = [2]int64{root.LessThan, 4001}
	}

	result = append(result, listLt...)
	result = append(result, listGte...)

	return result
}

func Part2() (combsCount int64) {
	system := parse()

	tree := system.createTree(Begin, 0)

	result := combinations(tree)

	return numberOfCombinations(result)
}
