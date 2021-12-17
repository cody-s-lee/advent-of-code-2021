package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

// Part one is the travelling salesman problem. Ugh.
func main() {
	body := get_input.Body("https://adventofcode.com/2015/day/10/input")
	body = strings.TrimSpace(body)
	sequence := make([]int, 0)

	for _, letter := range strings.Split(body, "") {
		n, err := strconv.Atoi(letter)
		if err != nil {
			log.Fatalf("Could not convert %s\n", letter)
		}
		sequence = append(sequence, n)
	}

	partOneResult := run(sequence, 40)
	partTwoResult := run(sequence, 50)

	fmt.Printf("Part one result: %d\n", partOneResult)
	fmt.Printf("Part two result: %d\n", partTwoResult)
}

func run(sequence []int, steps int) int {
	for step := 0; step < steps; step++ {
		newSequence := make([]int, 0)

		count := 0
		n := sequence[0]
		for i := 0; i < len(sequence); i++ {
			if sequence[i] == n {
				count++
				continue
			}
			newSequence = append(newSequence, count)
			newSequence = append(newSequence, n)
			count = 1
			n = sequence[i]
		}
		newSequence = append(newSequence, count)
		newSequence = append(newSequence, n)

		sequence = newSequence
		fmt.Printf("Sequence after steps #%d: %v\n", step, sequence)
	}

	fmt.Printf("Sequence after steps steps: %v\n", sequence)
	return len(sequence)
}
