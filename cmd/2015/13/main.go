package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

// Part one is the travelling salesman problem. Ugh.
func main() {
	lines := get_input.Lines("https://adventofcode.com/2015/day/13/input")

	happiness := make(map[string]map[string]int)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		who, rest := get_input.SplitPair(line, " would ")
		action, rest := get_input.SplitPair(rest, " happiness units by sitting next to ")
		them := strings.Trim(rest, " .")

		_, ok := happiness[who]
		if !ok {
			happiness[who] = make(map[string]int)
		}

		slug, amount := get_input.SplitPair(action, " ")
		value, _ := strconv.Atoi(amount)
		if slug == "lose" {
			value = -value
		}

		happiness[who][them] = value
	}

	people := make([]string, 0)
	for who := range happiness {
		people = append(people, who)
	}
	base := []string{}
	remainder := people[:]
	arrangements := pick(base, remainder)
	fmt.Printf("Found %d arrangements\n", len(arrangements))

	bestCycleArrangement := arrangements[0]
	bestNonCycleArrangement := arrangements[0]

	for _, arrangement := range arrangements {
		if CycleScore(happiness, arrangement) > CycleScore(happiness, bestCycleArrangement) {
			bestCycleArrangement = arrangement
		}
		if NonCycleScore(happiness, arrangement) > NonCycleScore(happiness, bestNonCycleArrangement) {
			bestNonCycleArrangement = arrangement
		}
	}

	partOneResult := CycleScore(happiness, bestCycleArrangement)
	partTwoResult := NonCycleScore(happiness, bestNonCycleArrangement)

	fmt.Printf("Part one result: %d\n", partOneResult)
	fmt.Printf("Part two result: %d\n", partTwoResult)
}

func pick(base []string, people []string) [][]string {
	arrangements := make([][]string, 0)

	if len(people) == 0 {
		arrangements = append(arrangements, base)
		return arrangements
	}

	for i, who := range people {
		newBase := make([]string, len(base)+1)
		copy(newBase, base)

		newPeople := make([]string, 0)
		for j, person := range people {
			if i == j {
				continue
			}
			newPeople = append(newPeople, person)
		}

		newBase[len(newBase)-1] = who
		arrangements = append(arrangements, pick(newBase, newPeople)...)
	}

	return arrangements
}

func CycleScore(happiness map[string]map[string]int, arrangement []string) int {
	if len(arrangement) == 0 {
		return 0
	}
	score := 0
	for i := 0; i < len(arrangement); i++ {
		j := i + 1
		if j == len(arrangement) {
			j = 0
		}

		who, them := arrangement[i], arrangement[j]
		whoToThem := happiness[who][them]
		themToWho := happiness[them][who]

		score += whoToThem
		score += themToWho
	}

	return score
}

func NonCycleScore(happiness map[string]map[string]int, arrangement []string) int {
	if len(arrangement) == 0 {
		return 0
	}
	score := 0
	for i := 0; i < len(arrangement)-1; i++ {
		j := i + 1

		who, them := arrangement[i], arrangement[j]
		whoToThem := happiness[who][them]
		themToWho := happiness[them][who]

		score += whoToThem
		score += themToWho
	}

	return score
}

func ArrangementShortString(arrangement []string) string {
	if len(arrangement) == 0 {
		return ""
	}
	runes := make([]rune, len(arrangement))
	for i, who := range arrangement {
		runes[i] = []rune(who)[0]
	}
	return string(runes)
}
func ArrangementString(happiness map[string]map[string]int, arrangement []string) string {
	return fmt.Sprintf("%s(%d)", ArrangementShortString(arrangement), CycleScore(happiness, arrangement))
}
