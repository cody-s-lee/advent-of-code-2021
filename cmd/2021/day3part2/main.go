package main

import (
	"fmt"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func filter(bits [][]bool, pos int, isOx bool) [][]bool {
	fmt.Printf("Filtering for position %d, isOx: %t\n", pos, isOx)

	f := make([][]bool, 0)

	nT, nF := 0, 0

	for i := 0; i < len(bits); i++ {
		fmt.Printf("Considering bits %v, looking at position %d, value is %t\n", bits[i], pos, bits[i][pos])
		switch bits[i][pos] {
		case true:
			nT++
		case false:
			nF++
		}
	}

	fmt.Printf("Found %d true and %d false\n", nT, nF)

	if nT == nF {
		nT++
	}

	for i := 0; i < len(bits); i++ {
		switch isOx {
		case true: // keep most common
			if nT > nF && bits[i][pos] {
				f = append(f, bits[i])
			} else if nF > nT && !bits[i][pos] {
				f = append(f, bits[i])
			}
		case false: // keep least common
			if nT > nF && !bits[i][pos] {
				f = append(f, bits[i])
			} else if nF > nT && bits[i][pos] {
				f = append(f, bits[i])
			}
		}

	}

	return f
}

func main() {
	bs := get_input.Body("https://adventofcode.com/2021/day/3/input")

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

	oxs, co2s := bits, bits

	for j := 0; j < x; j++ {
		if len(oxs) > 1 {
			oxs = filter(oxs, j, true)
		}
		if len(co2s) > 1 {
			co2s = filter(co2s, j, false)
		}
	}

	ox, co2 := oxs[0], co2s[0]

	fmt.Printf("ox = %v, co2 = %v", ox, co2)

	oxb, co2b := "", ""
	oxx, co2x := 0, 0

	for j := 0; j < x; j++ {
		p := x - j - 1

		fmt.Printf("ox[%d] = %t, co2[%d] = %t, p = %d, x = %d\n", j, ox[j], j, co2[j], p, 1<<p)

		if ox[j] {
			oxx += 1 << p
			oxb = oxb + "1"
		} else {
			oxb = oxb + "0"
		}

		if co2[j] {
			co2x += 1 << p
			co2b = co2b + "1"
		} else {
			co2b = co2b + "0"
		}
	}
	oxb, co2b = oxb+"b", co2b+"b"

	fmt.Printf("ox = %s = %d\n", oxb, oxx)
	fmt.Printf("co2 = %s = %d\n", co2b, co2x)

	fmt.Printf("ox * co2 = %d\n", oxx*co2x)
}
