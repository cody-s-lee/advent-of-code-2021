package main

import (
	"fmt"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	lines := get_input.Lines("https://adventofcode.com/2015/day/5/input")

	oldNiceCount := 0
	newNiceCount := 0

	for _, line := range lines {
		{ // old rules
			// vowels
			vowelCount := 0
			vowelCount += strings.Count(line, "a")
			vowelCount += strings.Count(line, "e")
			vowelCount += strings.Count(line, "i")
			vowelCount += strings.Count(line, "o")
			vowelCount += strings.Count(line, "u")
			hasAtLeastThreeVowels := vowelCount >= 3

			// doubles
			hasDoubles := false
			for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
				if strings.Contains(line, string(letter)+string(letter)) {
					hasDoubles = true
					break
				}
			}

			// forbidden couplets
			hasForbiddenCouplet := false
			for _, forbiddenCouplet := range []string{"ab", "cd", "pq", "xy"} {
				if strings.Contains(line, forbiddenCouplet) {
					hasForbiddenCouplet = true
					break
				}
			}

			if hasAtLeastThreeVowels && hasDoubles && !hasForbiddenCouplet {
				oldNiceCount++
			}
		}
		{ // new rules
			// doubled non-overlapping couplet
			hasDoubledNonOverlappingCouplet := false
			for i := 0; i < len(line)-3 && !hasDoubledNonOverlappingCouplet; i++ {
				for j := i + 2; j < len(line)-1; j++ {
					if line[i] == line[j] && line[i+1] == line[j+1] {
						hasDoubledNonOverlappingCouplet = true
						break
					}
				}
			}

			// split double
			hasSplitDouble := false
			for i := 0; i < len(line)-2; i++ {
				if line[i] == line[i+2] {
					hasSplitDouble = true
					break
				}
			}

			if hasDoubledNonOverlappingCouplet && hasSplitDouble {
				newNiceCount++
			}
		}
	}

	fmt.Printf("Found %d nice lines with old rules\n", oldNiceCount)
	fmt.Printf("Found %d nice lines with new rules\n", newNiceCount)
}
