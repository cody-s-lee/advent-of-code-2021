package main

import (
	"fmt"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
	"github.com/cody-s-lee/advent-of-code-2021/internal/point"
)

func main() {
	lines := get_input.Lines("https://adventofcode.com/2021/day/5/input")

	board := make(map[point.Point]int)
	maxX, maxY := 0, 0

	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}

		from_str, to_str := get_input.SplitPair(l, " -> ")
		from, to := point.Parse(from_str), point.Parse(to_str)

		if from.X > maxX {
			maxX = from.X
		}
		if from.Y > maxY {
			maxY = from.Y
		}
		if to.X > maxX {
			maxX = to.X
		}
		if to.Y > maxY {
			maxY = to.Y
		}

		fmt.Printf("Placing from %v to %v on board\n", from, to)

		if from.Y == to.Y { // horizontal
			fmt.Printf("Processing horizontal line\n")
			if from.X > to.X { // if leftward flip it around
				from, to = to, from
			}
			for i := from.X; i <= to.X; i++ {
				p := point.New(i, from.Y)
				board[p] = board[p] + 1
			}
		} else if from.X == to.X { // vertical
			fmt.Printf("Processing vertical line\n")
			if from.Y > to.Y { // if upward flip it around
				from, to = to, from
			}
			for j := from.Y; j <= to.Y; j++ {
				p := point.New(from.X, j)
				board[p] = board[p] + 1
			}
		} else { // diagonal
			fmt.Printf("Processing diagonal line\n")

			dx, dy := 1, 1
			if from.X > to.X {
				dx = -1
			}
			if from.Y > to.Y {
				dy = -1
			}
			for i, j := from.X, from.Y; i != to.X && j != to.Y; i, j = i+dx, j+dy {
				p := point.New(i, j)
				board[p] = board[p] + 1
			}
			{
				p := point.New(to.X, to.Y)
				board[p] = board[p] + 1
			}

		}
	}

	n := 0
	for _, v := range board {
		if v > 1 {
			n++
		}
	}

	fmt.Printf("Found %d overlaps\n", n)

}
