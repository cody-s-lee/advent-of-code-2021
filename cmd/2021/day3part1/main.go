package main

import (
	"fmt"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	bs := get_input.GetInput("https://adventofcode.com/2021/day/3/input")

	fmt.Print(bs)

	list := strings.Split(strings.TrimSuffix(bs, "\n"), "\n")

	x, y := len(list[0]), len(list)

	// 2-dimensional array of bits
	bits := make([][]bool, y)

	for i := 0; i < y; i++ {
		s := list[i]
		bits[i] = make([]bool, x)

		chars := []rune(s)

		for j := 0; j < x; j++ {
			switch chars[j] {
			case '0':
				bits[i][j] = false
			case '1':
				bits[i][j] = true
			}
		}

		fmt.Printf("Decomposed line %s into %v\n", s, bits[i])
	}

	gammaBits := make([]bool, x)

	for j := 0; j < x; j++ {
		fmt.Printf("Setting bit #%d of gamma\n", j)
		n := 0

		for i := 0; i < y; i++ {
			if bits[i][j] {
				n++
			}
		}
		gammaBits[j] = n > y/2

		fmt.Printf("Found %d of %d to be true, that means we set gammaBits[%d] to %t\n", n, y, j, gammaBits[j])
	}

	gamma, epsilon := 0, 0
	gb, eb := "", ""

	for j := 0; j < x; j++ {
		p := x - j - 1

		fmt.Printf("gammaBits[%d] = %t, p = %d, x = %d\n", j, gammaBits[j], p, 1<<p)

		if gammaBits[j] {
			gamma += 1 << p
			gb = gb + "1"
			eb = eb + "0"
		} else {
			epsilon += 1 << p
			gb = gb + "0"
			eb = eb + "1"
		}
	}
	gb, eb = gb+"b", eb+"b"

	fmt.Printf("gamma = %s = %d\n", gb, gamma)
	fmt.Printf("epsilon = %s = %d\n", eb, epsilon)

	fmt.Printf("gamma, epsilon = %d, %d\n", gamma, epsilon)

	fmt.Printf("gamma * epsilon = %d\n", gamma*epsilon)
}
