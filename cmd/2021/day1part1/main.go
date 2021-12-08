package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	bs := get_input.GetInput("https://adventofcode.com/2021/day/1/input")

	var dl []int
	for _, s := range strings.Split(bs, "\n") {
		s = strings.Trim(s, " ")
		if s != "" {
			x, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("Could not convert #{s}\n")
			}
			dl = append(dl, x)
		}
	}

	increases := -1
	p := -1
	for _, x := range dl {
		increase := x > p
		slug := "decreased"
		if increase {
			increases++
			slug = "increased"
		}
		p = x
		fmt.Printf("%d (%s)\n", x, slug)
	}
	fmt.Printf("%d total increases\n", increases)
}
