package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	body := get_input.Body("https://adventofcode.com/2021/day/17/input")

	_, body = get_input.SplitPair(body, "x=")
	exes, wyes := get_input.SplitPair(body, ", y=")
	leftString, rightString := get_input.SplitPair(exes, "..")
	bottomString, topString := get_input.SplitPair(wyes, "..")

	left, err := strconv.Atoi(leftString)
	if err != nil {
		log.Fatal(err)
	}
	right, err := strconv.Atoi(rightString)
	if err != nil {
		log.Fatal(err)
	}
	bottom, err := strconv.Atoi(bottomString)
	if err != nil {
		log.Fatal(err)
	}
	top, err := strconv.Atoi(topString)
	if err != nil {
		log.Fatal(err)
	}

	// given k+1 steps:
	// vx + (vx-1) + (vx-2) + ... + max(0, (vx-k)) >=l && <= r
	// if k > vx -> 1 + 2 + ... + vx = vx * (vx-1) / 2
	// else
	// vy + (vy+1) + (vy+2) + ... + (vy + k) <= u && >= d
	// = (k+1)*vy + (k*(k-1))/2 = k vy + vy + k^2/2 - k/2 = 1/2 k^2 + (1/2 - vy) k + vy

	// possible results are -vy -vy-1 -vy-2 ... -k vy - (triangular number k)
	bestMaxHeight := 0
	solutions := 0
	for vy := -bottom; vy >= bottom; vy-- {
		for vx := 0; vx < 200; vx++ {
			if maxHeight, ok := check(vx, vy, left, right, bottom, top); ok {
				if maxHeight > bestMaxHeight {
					bestMaxHeight = maxHeight
				}
				solutions++
			}
		}
	}

	fmt.Printf("Best targetting max height %d\n", bestMaxHeight)
	fmt.Printf("Found %d solutions\n", solutions)
}

func check(vx, vy, left, right, bottom, top int) (int, bool) {
	x, y := 0, 0
	maxHeight := 0

	for x <= right && y >= bottom {
		x, y = x+vx, y+vy
		if y > maxHeight {
			maxHeight = y
		}
		vx--
		if vx < 0 {
			vx = 0
		}
		vy--

		if x >= left && x <= right && y <= top && y >= bottom {
			return maxHeight, true
		}
	}

	return 0, false
}
