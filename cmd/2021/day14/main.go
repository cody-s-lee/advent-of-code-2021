package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	lines := get_input.Lines("https://adventofcode.com/2021/day/14/input")

	var polymer string
	rules := make(map[string]*Rule)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		} else if strings.Contains(line, "->") {
			pair, insertion := get_input.SplitPair(line, "->")
			l, r := get_input.SplitPair(pair, "")
			rules[pair] = &Rule{
				Left:  l + insertion,
				Right: insertion + r,
			}
		} else {
			polymer = line
		}
	}

	pairs := make(map[string]int)
	for i := 0; i < len(polymer)-1; i++ {
		pair := string(polymer[i : i+2])
		pairs[pair]++
	}

	partOneResult := 0
	partTwoResult := 0

	{
		for i := 0; i < 10; i++ {
			fmt.Printf("Applying update #%d\n", i+1)
			pairs = update(pairs, rules)
		}

		elementCounts := make(map[string]int)
		elementCounts[polymer[0:1]]++
		elementCounts[polymer[len(polymer)-1:]]++

		for pair, count := range pairs {
			elementCounts[pair[0:1]] += count
			elementCounts[pair[1:2]] += count
		}

		mostCommonCount := 0
		leastCommonCount := math.MaxInt
		for _, count := range elementCounts {
			if count > mostCommonCount {
				mostCommonCount = count
			}
			if count < leastCommonCount {
				leastCommonCount = count
			}
		}

		partOneResult = (mostCommonCount - leastCommonCount) / 2
	}

	fmt.Printf("Part one result: %d\n", partOneResult)

	{
		for i := 10; i < 40; i++ {
			fmt.Printf("Applying update #%d\n", i+1)
			pairs = update(pairs, rules)
		}

		elementCounts := make(map[string]int)
		elementCounts[polymer[0:1]]++
		elementCounts[polymer[len(polymer)-1:]]++

		for pair, count := range pairs {
			elementCounts[pair[0:1]] += count
			elementCounts[pair[1:2]] += count
		}

		mostCommonCount := 0
		leastCommonCount := math.MaxInt
		for _, count := range elementCounts {
			if count > mostCommonCount {
				mostCommonCount = count
			}
			if count < leastCommonCount {
				leastCommonCount = count
			}
		}

		partTwoResult = (mostCommonCount - leastCommonCount) / 2
	}

	fmt.Printf("Part one result: %d\n", partOneResult)
	fmt.Printf("Part two result: %d\n", partTwoResult)

}

func update(pairs map[string]int, rules map[string]*Rule) map[string]int {
	newPairs := make(map[string]int)

	for pair, count := range pairs {
		newPairs[rules[pair].Left] += count
		newPairs[rules[pair].Right] += count
	}

	return newPairs
}

type Rule struct {
	Left  string
	Right string
}
