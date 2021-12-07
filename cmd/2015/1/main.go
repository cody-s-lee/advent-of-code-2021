package main

import (
	"fmt"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	body := get_input.GetInput("https://adventofcode.com/2015/day/1/input")
	body = strings.Trim(body, "[]\n")
	fmt.Printf("%v\n", body)

	floor := 0
	basementPosition := 0
	for i, l := range body {
		switch l {
		case '(':
			floor++
		case ')':
			floor--
		}
		if basementPosition == 0 && floor < 0 {
			basementPosition = i + 1
		}
	}

	fmt.Printf("Reached floor %d\n", floor)
	fmt.Printf("Reached basement at position %d\n", basementPosition)
}
