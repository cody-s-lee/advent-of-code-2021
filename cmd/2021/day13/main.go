package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
	"github.com/cody-s-lee/advent-of-code-2021/internal/point"
)

func main() {
	lines := get_input.Lines("https://adventofcode.com/2021/day/13/input")

	dotField, foldInstructions := Parse(lines)

	dotField = Fold(dotField, foldInstructions[0:1])
	partOneResult := Count(dotField)

	dotField = Fold(dotField, foldInstructions[1:])
	partTwoResult := Draw(dotField)

	fmt.Printf("Part one found %d dots\n", partOneResult)
	fmt.Printf("Part two result:\n%s\n", partTwoResult)
}

func Parse(lines []string) ([][]bool, []FoldInstruction) {

	dots := make([]point.Point, 0)
	foldInstructions := make([]FoldInstruction, 0)

	maxX, maxY := 0, 0

	for _, line := range lines {
		if strings.Contains(line, "fold along") {
			fi := ParseFoldInstruction(line)
			foldInstructions = append(foldInstructions, fi)

		} else if strings.Contains(line, ",") {
			dot := point.Parse(line)
			dots = append(dots, dot)
			if dot.X > maxX {
				maxX = dot.X
			}
			if dot.Y > maxY {
				maxY = dot.Y
			}
		}
	}

	maxX++
	maxY++

	dotField := make([][]bool, maxX)
	for i := 0; i < maxX; i++ {
		dotField[i] = make([]bool, maxY)
	}

	for _, dot := range dots {
		dotField[dot.X][dot.Y] = true
	}

	fmt.Printf("Found %d dots\n", len(dotField))

	return dotField, foldInstructions
}

func Fold(dotField [][]bool, foldInstructions []FoldInstruction) [][]bool {
	fmt.Printf("Folds: %v\n", foldInstructions)

	for _, fi := range foldInstructions {
		fmt.Printf("Fold: %v\n", fi)
		switch fi.Orientation {
		case x:
			for y := fi.LineNumber; y < len(dotField[0]); y++ {
				for x := 0; x < len(dotField); x++ {
					if dotField[x][y] {
						newY := 2*fi.LineNumber - y
						fmt.Printf("Moving dot around %d from (%d, %d) to (%d, %d)\n", fi.LineNumber, x, y, x, newY)
						dotField[x][y] = false
						dotField[x][newY] = true
					}
				}
			}
		case y:
			for x := fi.LineNumber; x < len(dotField); x++ {
				for y := 0; y < len(dotField[0]); y++ {
					if dotField[x][y] {
						newX := 2*fi.LineNumber - x
						fmt.Printf("Moving dot around %d from (%d, %d) to (%d, %d)\n", fi.LineNumber, x, y, newX, y)
						dotField[x][y] = false
						dotField[newX][y] = true
					}
				}
			}
		}
	}

	return dotField
}

func Count(dotField [][]bool) int {
	count := 0
	for x := 0; x < len(dotField); x++ {
		for y := 0; y < len(dotField[0]); y++ {
			if dotField[x][y] {
				count++
			}
		}
	}

	return count
}

type FoldInstruction struct {
	Orientation Orientation
	LineNumber  int
}

func ParseFoldInstruction(line string) FoldInstruction {
	_, rest := get_input.SplitPair(line, "along")
	orientationString, lineNumberString := get_input.SplitPair(rest, "=")
	lineNumber, err := strconv.Atoi(lineNumberString)
	if err != nil {
		panic(err)
	}

	return FoldInstruction{
		Orientation: Orientation(orientationString),
		LineNumber:  lineNumber,
	}
}

type Orientation string

const (
	x Orientation = "y"
	y Orientation = "x"
)

func Draw(dotField [][]bool) string {
	maxX := 0
	maxY := 0
	for y := 0; y < len(dotField[0]); y++ {
		for x := 0; x < len(dotField); x++ {
			if dotField[x][y] {
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	maxX++
	maxY++

	result := ""
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if dotField[x][y] {
				result += "#"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}
