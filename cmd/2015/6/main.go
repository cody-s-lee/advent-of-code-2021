package main

import (
	"fmt"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
	"github.com/cody-s-lee/advent-of-code-2021/internal/point"
)

type action int

const (
	on action = iota
	off
	toggle
)

func main() {
	lines := get_input.Lines("https://adventofcode.com/2015/day/6/input")

	var oldGrid [1000][1000]bool
	var newGrid [1000][1000]int

	for _, line := range lines {
		line = strings.TrimSpace(line)

		action, start, end := parse(line)

		for i := start.X; i <= end.X; i++ {
			for j := start.Y; j <= end.Y; j++ {
				switch action {
				case on:
					oldGrid[i][j] = true
					newGrid[i][j] = newGrid[i][j] + 1
				case off:
					oldGrid[i][j] = false
					newGrid[i][j] = newGrid[i][j] - 1
					if newGrid[i][j] < 0 {
						newGrid[i][j] = 0
					}
				case toggle:
					oldGrid[i][j] = !oldGrid[i][j]
					newGrid[i][j] = newGrid[i][j] + 2
				}
			}
		}
	}

	oldLightsLit := 0
	newLightsBrightness := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if oldGrid[i][j] {
				oldLightsLit++
			}
			newLightsBrightness += newGrid[i][j]
		}
	}

	fmt.Printf("Old grid has %d lights lit\n", oldLightsLit)
	fmt.Printf("New grid has total %d brightness\n", newLightsBrightness)
}

func parse(line string) (action, point.Point, point.Point) {

	var action action

	if strings.HasPrefix(line, "turn on ") {
		action = on
		line = strings.TrimPrefix(line, "turn on ")
	} else if strings.HasPrefix(line, "turn off ") {
		action = off
		line = strings.TrimPrefix(line, "turn off ")
	} else if strings.HasPrefix(line, "toggle ") {
		action = toggle
		line = strings.TrimPrefix(line, "toggle ")
	}

	s, t := get_input.SplitPair(line, " through ")

	return action, point.Parse(s), point.Parse(t)
}
