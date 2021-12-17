package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	bs := get_input.Body("https://adventofcode.com/2021/day/2/input")

	distance, depth := 0, 0

	for _, action := range strings.Split(strings.TrimSuffix(bs, "\n"), "\n") {
		fmt.Printf("Moving %s\n", action)

		splits := strings.SplitN(action, " ", 2)
		direction, s := splits[0], splits[1]
		x, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Could not convert \"%s\" to int\n", s)
		}

		switch direction {
		case "forward":
			distance += x
		case "up":
			depth -= x
		case "down":
			depth += x
		}
	}

	fmt.Printf("distance * depth = %d\n", distance*depth)
}
