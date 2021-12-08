package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func parse(s string) (uint16, bool) {
	x, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, false
	}
	return uint16(x), true
}

func main() {
	partOneResult := partOne()
	partTwoResult := partTwo()
	fmt.Printf("After part one a has signal value of %d\n", partOneResult)
	fmt.Printf("After part two a has signal value of %d\n", partTwoResult)
}

func partOne() uint16 {
	lines := get_input.Lines("https://adventofcode.com/2015/day/7/input")

	signal := make(map[string]uint16)
	queue := lines[:]

	for len(queue) > 0 {
		line := queue[0]
		queue = queue[1:]

		fmt.Printf("Considering line \"%s\", queue size %d\n", line, len(queue))

		left, dst := get_input.SplitPair(line, "->")

		// NOT
		if strings.HasPrefix(left, "NOT") {
			_, src := get_input.SplitPair(left, " ")

			x, ok := signal[src]
			if !ok {
				fmt.Printf("Waiting until location %s has signal\n", dst)
				queue = append(queue, line)
				continue
			}

			signal[dst] = ^x
			continue
		}

		// AND
		if strings.Contains(left, "AND") {
			srcA, srcB := get_input.SplitPair(left, "AND")
			a, ok := parse(srcA)
			if !ok { // Is a a literal?
				a, ok = signal[srcA]
				if !ok {
					fmt.Printf("Waiting until location %s has signal\n", srcA)
					queue = append(queue, line)
					continue
				}
			}
			b, ok := parse(srcB)
			if !ok { // Is b a literal?
				b, ok = signal[srcB]
				if !ok {
					fmt.Printf("Waiting until location %s has signal\n", srcB)
					queue = append(queue, line)
					continue
				}
			}

			signal[dst] = a & b
			continue
		}

		// OR
		if strings.Contains(left, "OR") {
			srcA, srcB := get_input.SplitPair(left, "OR")
			a, ok := parse(srcA)
			if !ok { // Is a a literal?
				a, ok = signal[srcA]
				if !ok {
					fmt.Printf("Waiting until location %s has signal\n", srcA)
					queue = append(queue, line)
					continue
				}
			}
			b, ok := parse(srcB)
			if !ok { // Is b a literal?
				b, ok = signal[srcB]
				if !ok {
					fmt.Printf("Waiting until location %s has signal\n", srcB)
					queue = append(queue, line)
					continue
				}
			}

			signal[dst] = a | b
			continue
		}

		// LSHIFT
		if strings.Contains(left, "LSHIFT") {
			src, s := get_input.SplitPair(left, "LSHIFT")
			x, ok := signal[src]
			if !ok {
				fmt.Printf("Waiting until location %s has signal\n", src)
				queue = append(queue, line)
				continue
			}

			n, _ := parse(s)
			signal[dst] = x << n
			continue
		}

		// RSHIFT
		if strings.Contains(left, "RSHIFT") {
			src, s := get_input.SplitPair(left, "RSHIFT")
			x, ok := signal[src]
			if !ok {
				fmt.Printf("Waiting until location %s has signal\n", src)
				queue = append(queue, line)
				continue
			}

			n, _ := parse(s)
			signal[dst] = x >> n
			continue
		}

		// Literal assignment
		x, ok := parse(left)
		if !ok {
			x, ok := signal[left]
			if !ok {
				fmt.Printf("Waiting until location %s has signal\n", left)
				queue = append(queue, line)
				continue
			}
			signal[dst] = x
			continue
		}
		fmt.Printf("Assigning literal %d to location %s\n", x, dst)
		signal[dst] = x
		continue

	}
	return signal["a"]
}

func partTwo() uint16 {
	lines := get_input.Lines("https://adventofcode.com/2015/day/7/input")

	signal := make(map[string]uint16)
	signal["b"] = 16076
	queue := lines[:]

	for len(queue) > 0 {
		line := queue[0]
		queue = queue[1:]

		fmt.Printf("Considering line \"%s\", queue size %d\n", line, len(queue))

		left, dst := get_input.SplitPair(line, "->")

		// skip re-assignment of wire b
		if dst == "b" {
			fmt.Printf("Skipping wire b assignment\n")
			continue
		}

		// NOT
		if strings.HasPrefix(left, "NOT") {
			_, src := get_input.SplitPair(left, " ")

			x, ok := signal[src]
			if !ok {
				fmt.Printf("Waiting until location %s has signal\n", dst)
				queue = append(queue, line)
				continue
			}

			signal[dst] = ^x
			continue
		}

		// AND
		if strings.Contains(left, "AND") {
			srcA, srcB := get_input.SplitPair(left, "AND")
			a, ok := parse(srcA)
			if !ok { // Is a a literal?
				a, ok = signal[srcA]
				if !ok {
					fmt.Printf("Waiting until location %s has signal\n", srcA)
					queue = append(queue, line)
					continue
				}
			}
			b, ok := parse(srcB)
			if !ok { // Is b a literal?
				b, ok = signal[srcB]
				if !ok {
					fmt.Printf("Waiting until location %s has signal\n", srcB)
					queue = append(queue, line)
					continue
				}
			}

			signal[dst] = a & b
			continue
		}

		// OR
		if strings.Contains(left, "OR") {
			srcA, srcB := get_input.SplitPair(left, "OR")
			a, ok := parse(srcA)
			if !ok { // Is a a literal?
				a, ok = signal[srcA]
				if !ok {
					fmt.Printf("Waiting until location %s has signal\n", srcA)
					queue = append(queue, line)
					continue
				}
			}
			b, ok := parse(srcB)
			if !ok { // Is b a literal?
				b, ok = signal[srcB]
				if !ok {
					fmt.Printf("Waiting until location %s has signal\n", srcB)
					queue = append(queue, line)
					continue
				}
			}

			signal[dst] = a | b
			continue
		}

		// LSHIFT
		if strings.Contains(left, "LSHIFT") {
			src, s := get_input.SplitPair(left, "LSHIFT")
			x, ok := signal[src]
			if !ok {
				fmt.Printf("Waiting until location %s has signal\n", src)
				queue = append(queue, line)
				continue
			}

			n, _ := parse(s)
			signal[dst] = x << n
			continue
		}

		// RSHIFT
		if strings.Contains(left, "RSHIFT") {
			src, s := get_input.SplitPair(left, "RSHIFT")
			x, ok := signal[src]
			if !ok {
				fmt.Printf("Waiting until location %s has signal\n", src)
				queue = append(queue, line)
				continue
			}

			n, _ := parse(s)
			signal[dst] = x >> n
			continue
		}

		// Literal assignment
		x, ok := parse(left)
		if !ok {
			x, ok := signal[left]
			if !ok {
				fmt.Printf("Waiting until location %s has signal\n", left)
				queue = append(queue, line)
				continue
			}
			signal[dst] = x
			continue
		}
		fmt.Printf("Assigning literal %d to location %s\n", x, dst)
		signal[dst] = x
		continue

	}

	return signal["a"]
}
