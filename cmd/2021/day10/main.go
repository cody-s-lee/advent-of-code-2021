package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	lines := get_input.Lines("https://adventofcode.com/2021/day/10/input")
	fmt.Printf("%v\n", lines)

	fmt.Printf("Part one result: %d\n", partOne(lines))

	fmt.Printf("Part two result: %d\n", partTwo(lines))

}

func getOpenBraces() []rune {
	return []rune{'(', '{', '[', '<'}
}

func getClosedBraces() []rune {
	return []rune{')', '}', ']', '>'}
}

func isOpenBrace(letter rune) bool {
	for _, l := range getOpenBraces() {
		if l == letter {
			return true
		}
	}
	return false
}

func isClosedBrace(letter rune) bool {
	for _, l := range getClosedBraces() {
		if l == letter {
			return true
		}
	}
	return false
}

func getBraceClosureMap() map[rune]rune {
	return map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
}
func getBracePointMap() map[rune]int {
	return map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
}

func getBracePoints(letter rune) int {
	return getBracePointMap()[letter]
}

func matches(open, close rune) bool {
	return getClosedBrace(open) == close
}

func getClosedBrace(open rune) rune {
	return getBraceClosureMap()[open]
}

func partOne(lines []string) int {
	total := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		openBraces := make([]rune, 0)
		for _, letter := range line {
			if isOpenBrace(letter) {
				openBraces = append(openBraces, letter)
			} else if isClosedBrace(letter) {
				openBrace := openBraces[len(openBraces)-1]     // get the rightmost open brace
				openBraces = openBraces[0 : len(openBraces)-1] // shorten the slice of braces

				if !matches(openBrace, letter) { // found our illegal brace!
					total += getBracePoints(letter)
					break
				}
			} else {
				log.Fatalf("Could not determine what letter '%v' represents\n", letter)
			}

		}
	}

	return total
}

func partTwo(lines []string) int {
	scores := make([]int, 0)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		isCorrupt := false
		openBraces := make([]rune, 0)
		for _, letter := range line {
			if isOpenBrace(letter) {
				openBraces = append(openBraces, letter)
			} else if isClosedBrace(letter) {
				openBrace := openBraces[len(openBraces)-1]     // get the rightmost open brace
				openBraces = openBraces[0 : len(openBraces)-1] // shorten the slice of braces

				if !matches(openBrace, letter) { // found a corrupted line, discard the whole line
					isCorrupt = true
					break
				}
			} else {
				log.Fatalf("Could not determine what letter '%v' represents\n", letter)
			}
		}

		if !isCorrupt && len(openBraces) > 0 { //incomplete line!
			closedBraces := reverse(openBraces)
			lineValue := 0
			for _, letter := range closedBraces {
				lineValue = lineValue*5 + completionValue(letter)
			}
			scores = append(scores, lineValue)
		}
	}

	sort.Ints(scores)
	fmt.Printf("Found %d scores %v\n, picking score #%d\n", len(scores), scores, len(scores)/2)
	return scores[len(scores)/2]
}

func reverse(letters []rune) []rune {
	srettel := make([]rune, len(letters))
	for i, l := range letters {
		srettel[len(letters)-i-1] = l
	}
	return srettel
}

func completionValue(letter rune) int {
	values := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	return values[letter]
}
