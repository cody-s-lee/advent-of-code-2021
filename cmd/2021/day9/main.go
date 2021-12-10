package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
	"github.com/cody-s-lee/advent-of-code-2021/internal/point"
)

func main() {
	lines := get_input.Lines("https://adventofcode.com/2021/day/9/input")
	fmt.Printf("%v\n", lines)

	fmt.Printf("Part one result: %d\n", partOne(lines))

	fmt.Printf("Part two result: %d\n", partTwo(lines))

}

func partOne(lines []string) int {
	risk := 0
	heights := make(heightMap)

	for y, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		for x, letter := range line {
			h, err := strconv.Atoi(string(letter))
			if err != nil {
				log.Fatalf("Could not convert %v\n", letter)
			}

			heights[point.New(x, y)] = h
		}
	}

	for p, h := range heights {
		if heights.isLow(p) {
			risk += 1 + h
		}
	}

	return risk
}

type heightMap map[point.Point]int

func (hm heightMap) isLow(p point.Point) bool {
	h, ok := hm[p]
	if !ok {
		return false // points not on the map can't be low points
	}

	adj := hm.adjacent(p)

	for _, x := range adj {
		if x <= h {
			return false
		}
	}

	return true // found no lower adjacent points
}

func partTwo(lines []string) int {
	heights := make(heightMap)
	basins := make(map[point.Point]int)

	for y, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		for x, letter := range line {
			h, err := strconv.Atoi(string(letter))
			if err != nil {
				log.Fatalf("Could not convert %v\n", letter)
			}

			heights[point.New(x, y)] = h
		}
	}

	for p := range heights {
		if heights.isLow(p) {
			basins[p] = heights.basinSize(p)
		}
	}

	basinSizes := make([]int, 0)
	for _, size := range basins {
		basinSizes = append(basinSizes, size)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	basinMultiple := 1
	basinMultiple *= basinSizes[0]
	basinMultiple *= basinSizes[1]
	basinMultiple *= basinSizes[2]

	return basinMultiple
}

func (hm heightMap) basinSize(p point.Point) int {
	fmt.Printf("Looking at basin at low point %v\n", p)

	basin := make(heightMap)
	basin[p] = hm[p]

	queue := make([]point.Point, 0)
	queue = append(queue, p)

	for {
		q := queue[0]
		fmt.Printf("Looking at point %v in basin at %v\n", q, p)

		queue = queue[1:]

		adj := hm.adjacent(q)
		fmt.Printf("Found adjacent points %v to consider\n", adj)

		for r, h := range adj {
			fmt.Printf("Considering adjacent point %v\n", r)
			_, ok := basin[r]
			if ok {
				fmt.Println("Already in basin")
				continue
			}
			if h == 9 {
				fmt.Println("It was a wall")
				continue
			}
			fmt.Printf("Adding to basin with value %d and to queue for consideration\n", h)
			basin[r] = h
			queue = append(queue, r)
		}

		if len(queue) == 0 {
			break
		}
	}

	fmt.Printf("Basin at %v is %v with size %d\n", p, basin, len(basin))

	return len(basin)
}

func (hm heightMap) adjacent(p point.Point) heightMap {
	points := make(heightMap)
	{
		_, ok := hm[p]
		if !ok {
			return points
		}
	}
	{
		q := point.New(p.X, p.Y-1)
		h, ok := hm[q]
		if ok {
			points[q] = h
		}
	}
	{
		q := point.New(p.X, p.Y+1)
		h, ok := hm[q]
		if ok {
			points[q] = h
		}
	}
	{
		q := point.New(p.X+1, p.Y)
		h, ok := hm[q]
		if ok {
			points[q] = h
		}
	}
	{
		q := point.New(p.X-1, p.Y)
		h, ok := hm[q]
		if ok {
			points[q] = h
		}
	}
	return points
}
