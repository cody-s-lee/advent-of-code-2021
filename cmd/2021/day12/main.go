package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

const (
	cn_start string = "start"
	cn_end   string = "end"
)

func main() {
	lines := get_input.Lines("https://adventofcode.com/2021/day/12/input")

	start := newCave(cn_start)
	end := newCave(cn_end)
	allCaves := map[string]*cave{
		cn_start: start,
		cn_end:   end,
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		left, right := get_input.SplitPair(line, "-")

		if _, ok := allCaves[left]; !ok {
			allCaves[left] = newCave(left)
		}
		if _, ok := allCaves[right]; !ok {
			allCaves[right] = newCave(right)
		}

		leftCave := allCaves[left]
		rightCave := allCaves[right]

		addNeighbor(leftCave, rightCave)
	}

	fmt.Printf("Found %d caves: %v\n", len(allCaves), allCaves)

	completeShortPaths := make([]Path, 0)
	completeLongPaths := make([]Path, 0)
	explorablePaths := make([]Path, 0)
	explorablePaths = append(explorablePaths, Path{start}) // seed the exploration

	i := 0
	for len(explorablePaths) > 0 {
		path := explorablePaths[0]
		explorablePaths = explorablePaths[1:]
		if i%10000 == 0 {
			fmt.Printf("Exploring path length %d %v\n", len(path), path)
		}
		i++
		last := path[len(path)-1]

		for c := range last.Neighbors {
			if c == start {
				continue
			}
			newPath := make(Path, len(path))
			copy(newPath, path)
			newPath = append(newPath, c)

			if c == end {
				if isValidShortPath(newPath) {
					completeShortPaths = append(completeShortPaths, newPath)
				}
				completeLongPaths = append(completeLongPaths, newPath)
			} else if isValidLongPath(newPath) {
				explorablePaths = append(explorablePaths, newPath)
			}
		}
	}

	fmt.Printf("Found %d complete short paths\n", len(completeShortPaths))
	fmt.Printf("Found %d complete long paths\n", len(completeLongPaths))
}

type Path []*cave

func PathsContain(paths []Path, path Path) bool {
	for _, p := range paths {
		if len(p) != len(path) {
			continue
		}
		pathDiffers := false
		for i := range p {
			if p[i] != path[i] {
				pathDiffers = true
				break
			}
		}
		if pathDiffers {
			continue
		}
		return true
	}
	return false
}

func isValidShortPath(path Path) bool {
	smallCaves := make(map[*cave]int)

	for _, cave := range path {
		if cave.Size == small {
			smallCaves[cave]++
			if smallCaves[cave] > 1 {
				return false
			}
		}
	}
	return true
}

func isValidLongPath(path Path) bool {
	smallCaves := make(map[*cave]int)
	revisited := false

	for _, cave := range path {
		if cave.Size == small {
			smallCaves[cave]++
			if smallCaves[cave] > 1 {
				if revisited {
					return false
				} else {
					revisited = true
				}
			}
		}
	}
	return true
}

func addNeighbor(cave, neighbor *cave) {
	cave.Neighbors[neighbor] = true
	neighbor.Neighbors[cave] = true
}

func newCave(name string) *cave {
	size := small
	r, _ := utf8.DecodeRuneInString(name)
	if unicode.IsUpper(r) {
		size = big
	}

	return &cave{
		Name:      name,
		Size:      size,
		Neighbors: make(map[*cave]bool),
	}
}

func (cave *cave) String() string {
	return fmt.Sprintf("Cave{%s, %s, [%d neighbors]}", cave.Name, cave.Size, len(cave.Neighbors))
}

func (size size) String() string {
	return caveSizes[size]
}

var caveSizes = map[size]string{
	small: "Small",
	big:   "Big",
}

type cave struct {
	Name      string
	Size      size
	Neighbors map[*cave]bool
}

type size int

const (
	small size = iota
	big
)
