package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
	"github.com/cody-s-lee/advent-of-code-2021/internal/point"
)

func main() {
	lines := get_input.Lines("https://adventofcode.com/2021/day/11/input")
	fmt.Printf("%v\n", lines)

	fmt.Printf("Part one result: %d\n", partOne(lines))

	fmt.Printf("Part two result: %d\n", partTwo(lines))

}

func partOne(lines []string) int {
	total := 0
	var octopusses [10][10]int

	for j, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		for i, l := range line {
			x, err := strconv.Atoi(string(l))
			if err != nil {
				log.Fatalf("Could not convert %s\n", string(l))
			}
			octopusses[i][j] = x
		}
	}

	for n := 0; n < 100; n++ { // Simulate 100 steps
		fmt.Printf("Starting step #%d\n", n)
		flashers := make([]point.Point, 0)
		flashed := make(map[point.Point]struct{}, 0)
		for i := 0; i < 10; i++ { // Stimulate every octopus
			for j := 0; j < 10; j++ {
				octopusses[i][j] = octopusses[i][j] + 1
				if octopusses[i][j] > 9 {
					flashers = append(flashers, point.New(i, j))
					total++
				}
			}
		}

		for len(flashers) > 0 { // flash all octopusses
			// grab a flasher off the queue
			flasher := flashers[0]
			flashers = flashers[1:]
			flashed[flasher] = exists
			fmt.Printf("Considering flasher %v\n", flasher)

			// find adjacent octopusses
			for i := flasher.X - 1; i <= flasher.X+1; i++ {
				for j := flasher.Y - 1; j <= flasher.Y+1; j++ {
					if i >= 0 && i <= 9 && j >= 0 && j <= 9 {
						p := point.New(i, j)
						fmt.Printf("Found adjacent octopus %v\n", p)
						// stimulate them by the flash
						octopusses[i][j] = octopusses[i][j] + 1
						if octopusses[i][j] > 9 {
							_, ok := flashed[p]                // hasn't flashed already
							if !ok && !contains(flashers, p) { // isn't about to flash
								// add any additional flashers to the flash queue
								flashers = append(flashers, p)
								total++
								fmt.Printf("Adjacent octopus %v flashing\n", p)
							}
						}
					}
				}
			}
		}

		for i := 0; i < 10; i++ { // Reset all octopusses
			for j := 0; j < 10; j++ {
				if octopusses[i][j] > 9 {
					octopusses[i][j] = 0
				}
			}
		}
		fmt.Printf("Finishing step #%d\n", n)
	}

	return total
}

func contains(queue []point.Point, point point.Point) bool {
	for _, p := range queue {
		if p == point {
			return true
		}
	}
	return false
}

func partTwo(lines []string) int {
	var octopusses [10][10]int

	for j, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		for i, l := range line {
			x, err := strconv.Atoi(string(l))
			if err != nil {
				log.Fatalf("Could not convert %s\n", string(l))
			}
			octopusses[i][j] = x
		}
	}

	for n := 0; n < 10000; n++ { // Simulate 100 steps
		fmt.Printf("Starting step #%d\n", n)
		flashers := make([]point.Point, 0)
		flashed := make(map[point.Point]struct{}, 0)
		for i := 0; i < 10; i++ { // Stimulate every octopus
			for j := 0; j < 10; j++ {
				octopusses[i][j] = octopusses[i][j] + 1
				if octopusses[i][j] > 9 {
					flashers = append(flashers, point.New(i, j))
				}
			}
		}

		for len(flashers) > 0 { // flash all octopusses
			// grab a flasher off the queue
			flasher := flashers[0]
			flashers = flashers[1:]
			flashed[flasher] = exists
			fmt.Printf("Considering flasher %v\n", flasher)

			// find adjacent octopusses
			for i := flasher.X - 1; i <= flasher.X+1; i++ {
				for j := flasher.Y - 1; j <= flasher.Y+1; j++ {
					if i >= 0 && i <= 9 && j >= 0 && j <= 9 {
						p := point.New(i, j)
						fmt.Printf("Found adjacent octopus %v\n", p)
						// stimulate them by the flash
						octopusses[i][j] = octopusses[i][j] + 1
						if octopusses[i][j] > 9 {
							_, ok := flashed[p]                // hasn't flashed already
							if !ok && !contains(flashers, p) { // isn't about to flash
								// add any additional flashers to the flash queue
								flashers = append(flashers, p)
								fmt.Printf("Adjacent octopus %v flashing\n", p)
							}
						}
					}
				}
			}
		}

		fmt.Printf("%d octopuses flashed on step #%d\n", len(flashed), n+1)
		if len(flashed) == 100 { // Party!
			return n + 1
		}

		for i := 0; i < 10; i++ { // Reset all octopusses
			for j := 0; j < 10; j++ {
				if octopusses[i][j] > 9 {
					octopusses[i][j] = 0
				}
			}
		}
		fmt.Printf("Finishing step #%d\n", n)
	}

	return 0
}

var exists = struct{}{}
