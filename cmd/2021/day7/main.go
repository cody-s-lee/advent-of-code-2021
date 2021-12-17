package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	body := get_input.Body("https://adventofcode.com/2021/day/7/input")
	body = strings.Trim(body, "[]\n ")

	fmt.Printf("%v\n", body)

	values := make([]int, 0)
	for _, s := range strings.Split(body, ",") {
		s = strings.Trim(s, " ")
		if s == "" {
			continue
		}
		x, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Could not convert \"%s\"\n", s)
		}
		values = append(values, x)
	}

	fmt.Printf("Part one result: %d\n", partOne(values))

	fmt.Printf("Part two result: %d\n", partTwo(values))
}

func partOne(values []int) int {
	min, max := values[0], values[0]

	for _, x := range values {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}

	fmt.Printf("Smallest postition found %d, biggest %d\n", min, max)

	bestPosition, bestCost := 0, math.MaxInt
	for n := min; n <= max; n++ {
		totalCost := 0
		for _, x := range values {
			cost := x - n

			if cost < 0 {
				totalCost -= cost
			} else {
				totalCost += cost
			}
		}
		fmt.Printf("Cost for position %d = %d\n", n, totalCost)
		if totalCost < bestCost {
			bestPosition = n
			bestCost = totalCost
		}
	}

	fmt.Printf("Best cost %d found at position %d\n", bestCost, bestPosition)

	return bestCost
}

func partTwo(values []int) int {
	min, max := values[0], values[0]

	for _, x := range values {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}

	fmt.Printf("Smallest postition found %d, biggest %d\n", min, max)

	bestPosition, bestCost := 0, math.MaxInt
	for n := min; n <= max; n++ {
		totalCost := 0
		for _, x := range values {
			distance := x - n

			if distance < 0 {
				distance = -distance
			}

			totalCost += (distance * (distance + 1)) / 2
		}
		fmt.Printf("Cost for position %d = %d\n", n, totalCost)
		if totalCost < bestCost {
			bestPosition = n
			bestCost = totalCost
		}
	}

	fmt.Printf("Best cost %d found at position %d\n", bestCost, bestPosition)

	return bestCost
}
