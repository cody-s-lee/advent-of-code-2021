package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

/*

0: _        1:          2: _        3: _        4:
..| |           |          _|          _|         |_|
..|_| -> 6      | -> 2    |_  -> 5     _| -> 5      | -> 4

5: _        6: _        7: _        8: _        9: _
..|_          |_            |         |_|         |_|
.. _| -> 5    |_| -> 6      | -> 3    |_| -> 7     _| -> 6

s    n
2 -> 1
3 -> 7
4 -> 4
5 -> 2, 3, 5
6 -> 0, 6, 9
7 -> 8

*/

func main() {
	lines := get_input.Lines("https://adventofcode.com/2021/day/8/input")

	fmt.Printf("%v\n", lines)

	fmt.Printf("Part one result: %d\n", partOne(lines))

	fmt.Printf("Part two result: %d\n", partTwo(lines))
}

func partOne(lines []string) int {
	digitCount := 0

	for _, line := range lines {
		_, outputs := get_input.SplitPair(line, "|")
		for _, digit := range strings.Split(outputs, " ") {
			switch len(digit) {
			case 2:
				digitCount++
			case 3:
				digitCount++
			case 4:
				digitCount++
			case 7:
				digitCount++
			}
		}
	}

	return digitCount
}

func overlap(x, y string) int {
	overlap := 0
	for _, letter := range y {
		if strings.ContainsRune(x, letter) {
			overlap++
		}
	}
	return overlap
}

func containsSubset(outer, inner string) bool {
	for _, letter := range inner {
		if !strings.ContainsRune(outer, letter) {
			return false
		}
	}
	return true
}

func partTwo(lines []string) int {
	total := 0

	for _, line := range lines {
		inputs, outputs := get_input.SplitPair(line, "|")

		var digits digits

		/*
		   s    n
		   2 -> 1
		   3 -> 7
		   4 -> 4
		   5 -> 2, 3, 5
		   6 -> 0, 6, 9
		   7 -> 8
		*/
		var fives, sixes []string
		// Start with the easy ones, and set up the hard ones
		for _, digit := range strings.Split(strings.TrimSpace(inputs), " ") {
			digit = sorted(digit)
			switch len(digit) {
			case 2:
				digits[1] = digit
			case 3:
				digits[7] = digit
			case 4:
				digits[4] = digit
			case 7:
				digits[8] = digit
			case 5:
				fives = append(fives, digit)
			case 6:
				sixes = append(sixes, digit)
			}
		}
		for _, digit := range fives {
			if containsSubset(digit, digits[1]) {
				// 5 segments, occupying the 1s segment makes this a 3
				digits[3] = digit
				continue
			}
			if overlap(digit, digits[4]) == 2 {
				digits[2] = digit
				continue
			}
			if overlap(digit, digits[4]) == 3 {
				digits[5] = digit
				continue
			}
		}
		for _, digit := range sixes {
			if !containsSubset(digit, digits[5]) {
				digits[0] = digit
				continue
			}
			if containsSubset(digit, digits[5]) && !containsSubset(digit, digits[1]) {
				digits[6] = digit
			}
			if containsSubset(digit, digits[5]) && containsSubset(digit, digits[1]) {
				digits[9] = digit
			}

		}

		lineValue := 0

		for _, output := range strings.Split(strings.TrimSpace(outputs), " ") {
			output = sorted(output)
			lineValue = lineValue*10 + digits.find(output)
		}

		total += lineValue
	}

	return total
}

type digits [10]string

func (digits digits) find(output string) int {
	for i, s := range digits {
		if s == output {
			return i
		}
	}
	return -1
}

func sorted(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}
