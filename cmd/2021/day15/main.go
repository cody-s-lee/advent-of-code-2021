package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
	"github.com/cody-s-lee/advent-of-code-2021/internal/point"
)

func main() {
	lines := get_input.Lines("https://adventofcode.com/2021/day/15/input")

	partOneResult := partOne(lines)
	partTwoResult := partTwo(lines)

	fmt.Printf("Part one result: %d\n", partOneResult)
	fmt.Printf("Part two result: %d\n", partTwoResult)
}

func partOne(lines []string) int {
	width, height := len(lines[0]), len(lines)

	risk := make(map[*point.Point]int)     // risk at that individual point
	distance := make(map[*point.Point]int) // distance from source to point
	points := make([][]*point.Point, height)
	queue := make([]*point.Point, 0)

	maxX, maxY := width-1, height-1

	for x := 0; x < width; x++ {
		points[x] = make([]*point.Point, height)
	}

	for y, line := range lines {
		for x, r := range line {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal(err)
			}

			point := &point.Point{
				X: x,
				Y: y,
			}
			points[x][y] = point
			risk[point] = n
			distance[point] = math.MaxInt
			queue = append(queue, point)
		}
	}

	source := points[0][0]
	target := points[maxX][maxY]

	previous := make(map[*point.Point]*point.Point) // prior point in path from source to point
	distance[source] = 0

	for len(queue) > 0 {
		u, newQ := closest(queue, distance)
		queue = newQ

		if u == target {
			break
		}

		for _, v := range neighbors(points, u) {
			newDist := distance[u] + risk[v]
			if newDist < distance[v] {
				distance[v] = newDist
				previous[v] = u
			}
		}

	}

	totalRisk := 0
	p := target
	for p != source {
		totalRisk += risk[p]
		p = previous[p]
	}

	return totalRisk
}
func partTwo(lines []string) int {
	tileWidth, tileHeight := len(lines[0]), len(lines)
	width, height := tileWidth*5, tileHeight*5

	risk := make(map[*point.Point]int)     // risk at that individual point
	distance := make(map[*point.Point]int) // distance from source to point
	points := make([][]*point.Point, height)
	queue := make([]*point.Point, 0)

	maxX, maxY := width-1, height-1

	for x := 0; x < width; x++ {
		points[x] = make([]*point.Point, height)
	}

	for y, line := range lines {
		for x, r := range line {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal(err)
			}

			for dX := 0; dX < 5; dX++ {
				for dY := 0; dY < 5; dY++ {
					eX := x + dX*tileWidth
					eY := y + dY*tileWidth
					eRisk := (n + dX + dY)
					if eRisk > 9 { // max eRisk is 9 + 4 + 4 = 17 -> 8
						eRisk -= 9 // 10 wraps to 1
					}
					if eRisk == 0 || eRisk > 9 {
						log.Fatalf("eRisk (%d, %d) out of range: %d\n", eX, eY, eRisk)
					}

					point := &point.Point{
						X: eX,
						Y: eY,
					}
					points[eX][eY] = point
					risk[point] = eRisk
					distance[point] = math.MaxInt
					queue = append(queue, point)
				}
			}
		}
	}

	source := points[0][0]
	target := points[maxX][maxY]

	previous := make(map[*point.Point]*point.Point) // prior point in path from source to point
	distance[source] = 0

	fmt.Printf("Moving on to search algo\n")

	i := 0
	for len(queue) > 0 {
		i++
		if i%100 == 0 {
			fmt.Printf("Iteration %d, queue length %d\n", i, len(queue))
		}
		u, newQ := closest(queue, distance)
		queue = newQ

		if u == target {
			break
		}

		for _, v := range neighbors(points, u) {
			newDist := distance[u] + risk[v]
			if newDist < distance[v] {
				distance[v] = newDist
				previous[v] = u
			}
		}
	}

	fmt.Printf("Done searching\n")

	totalRisk := 0
	p := target
	for p != source {
		totalRisk += risk[p]
		p = previous[p]
	}

	return totalRisk
}

func neighbors(points [][]*point.Point, p *point.Point) []*point.Point {
	neighbors := make([]*point.Point, 0)
	{
		// up
		x, y := p.X, p.Y-1
		if y >= 0 {
			neighbors = append(neighbors, points[x][y])
		}
	}
	{
		// right
		x, y := p.X+1, p.Y
		if x < len(points) {
			neighbors = append(neighbors, points[x][y])
		}
	}
	{
		// down
		x, y := p.X, p.Y+1
		if y < len(points[0]) {
			neighbors = append(neighbors, points[x][y])
		}
	}
	{
		// left
		x, y := p.X-1, p.Y
		if x >= 0 {
			neighbors = append(neighbors, points[x][y])
		}
	}
	return neighbors
}

func closest(queue []*point.Point, distance map[*point.Point]int) (*point.Point, []*point.Point) {
	var closest *point.Point
	remainder := make([]*point.Point, 0)

	for _, p := range queue {
		if closest == nil || distance[p] < distance[closest] {
			closest = p
		}
	}
	for _, p := range queue {
		if p != closest {
			remainder = append(remainder, p)
		}
	}

	return closest, remainder
}
