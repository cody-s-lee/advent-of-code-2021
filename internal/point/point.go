package point

import (
	"log"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

type Point struct {
	X int
	Y int
}

func New(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func Parse(s string) Point {
	m, n := get_input.SplitPair(strings.TrimSpace(s), ",")

	x, err := strconv.Atoi(m)
	if err != nil {
		log.Fatalf("Failed to convert \"%s\" from \"%s\"\n", m, s)
	}

	y, err := strconv.Atoi(n)
	if err != nil {
		log.Fatalf("Failed to convert \"%s\" from \"%s\"\n", m, s)
	}

	return New(x, y)
}
