package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	lines := get_input.Lines("https://adventofcode.com/2015/day/2/input")

	totalWrappingPaper := 0
	totalRibbon := 0

	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		r, s, t := get_input.SplitTriplet(l, "x")
		x, err := strconv.Atoi(r)
		if err != nil {
			log.Fatalf("Could not convert value \"%s\"\n", r)
		}
		y, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Could not convert value \"%s\"\n", s)
		}
		z, err := strconv.Atoi(t)
		if err != nil {
			log.Fatalf("Could not convert value \"%s\"\n", t)
		}

		a, b, c := x*y, x*z, y*z
		smallest := a
		if b < smallest {
			smallest = b
		}
		if c < smallest {
			smallest = c
		}

		totalWrappingPaper += 2*a + 2*b + 2*c + smallest // all faces plus extra equal to smallest face
		totalRibbon += x * y * z                         // ribbon bow requires volume
		totalRibbon += 2 * (x + y + z)

		biggestSide := x
		if y > biggestSide {
			biggestSide = y
		}
		if z > biggestSide {
			biggestSide = z
		}
		totalRibbon -= 2 * biggestSide
	}

	fmt.Printf("Total wrapping paper area needed is %d sqft\n", totalWrappingPaper)
	fmt.Printf("Total ribbon needed is %d ft\n", totalRibbon)
}
