package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	body := get_input.Body("https://adventofcode.com/2021/day/6/input")
	body = strings.Trim(body, "[]")

	fmt.Printf("%v\n", body)

	population := make(map[int]int)

	for _, s := range strings.Split(body, ",") {
		s = strings.TrimSpace(s)
		timer, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Could not convert \"%s\"\n", s)
		}
		population[timer] = population[timer] + 1
	}

	day := 0
	fmt.Printf("Initial population: %v\n", population)
	for day = 0; day < 256; day++ {
		prev := make(map[int]int)
		for k, v := range population {
			prev[k] = v
		}
		population[0] = prev[1]
		population[1] = prev[2]
		population[2] = prev[3]
		population[3] = prev[4]
		population[4] = prev[5]
		population[5] = prev[6]
		population[6] = prev[0] + prev[7]
		population[7] = prev[8]
		population[8] = prev[0]
		fmt.Printf("Population after day %d: %v\n", day+1, population)
	}

	total := 0
	for _, v := range population {
		total += v
	}

	fmt.Printf("Population: %v\nDay: %d\tTotal: %d\n", population, day, total)
}
