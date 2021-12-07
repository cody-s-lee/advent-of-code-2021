package main

import (
	"fmt"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
	"github.com/cody-s-lee/advent-of-code-2021/internal/point"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fmt.Printf("Part 1\n")

	body := get_input.GetInput("https://adventofcode.com/2015/day/3/input")
	body = strings.TrimSpace(body)

	houses := make(map[point.Point]int)
	x, y := 0, 0
	p := point.New(x, y)
	houses[p] = houses[p] + 1

	for _, d := range body {
		switch d {
		case '^':
			y--
		case 'v':
			y++
		case '<':
			x--
		case '>':
			x++
		}
		p = point.New(x, y)
		houses[p] = houses[p] + 1
	}

	visitedHouseCount := 0
	for _, v := range houses {
		if v > 0 {
			visitedHouseCount++
		}
	}

	fmt.Printf("Visited %d houses\n", visitedHouseCount)
}

func partTwo() {
	fmt.Printf("Part 2\n")

	body := get_input.GetInput("https://adventofcode.com/2015/day/3/input")
	body = strings.TrimSpace(body)

	houses := make(map[point.Point]int)
	santa := true
	x, y := 0, 0
	r, s := 0, 0
	p := point.New(x, y)
	houses[p] = houses[p] + 1

	for _, d := range body {
		switch d {
		case '^':
			if santa {
				y--
			} else {
				s--
			}
		case 'v':
			if santa {
				y++
			} else {
				s++
			}
		case '<':
			if santa {
				x--
			} else {
				r--
			}
		case '>':
			if santa {
				x++
			} else {
				r++
			}
		}
		if santa {
			p = point.New(x, y)
		} else {
			p = point.New(r, s)
		}
		houses[p] = houses[p] + 1
		santa = !santa
	}

	visitedHouseCount := 0
	for _, v := range houses {
		if v > 0 {
			visitedHouseCount++
		}
	}

	fmt.Printf("Visited %d houses\n", visitedHouseCount)
}
