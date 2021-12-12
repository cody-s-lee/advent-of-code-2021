package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

// Part one is the travelling salesman problem. Ugh.
func main() {
	lines := get_input.Lines("https://adventofcode.com/2015/day/9/input")
	locations := make(map[string]bool)
	legs := make([]Leg, 0)

	for _, line := range lines {
		leg := parse(line)
		legs = append(legs, leg)
		locations[leg.Start] = true
		locations[leg.End] = true
	}

	roots := make([]*Node, 0)
	for location := range locations {
		fmt.Printf("Found location %s\n", location)
		roots = append(roots, &Node{
			Location: location,
			Children: make([]*Node, 0),
		})
	}

	toExplore := make([]*Node, 0)
	toExplore = append(toExplore, roots...)

	i := 0
	for len(toExplore) > 0 {
		node := toExplore[0] // grab a node off explore list
		toExplore = toExplore[1:]

		// extend that node
		for location := range locations {
			if !node.hasAncestor(location) {
				child := &Node{
					Parent:   node,
					Location: location,
					Distance: distance(legs, location, node.Location),
					Children: make([]*Node, 0),
				}
				node.Children = append(node.Children, child)
				toExplore = append(toExplore, child)
				i++
				if i%1000 == 0 {
					fmt.Printf("Loaded %d nodes\n", i)
				}
			}
		}
		fmt.Printf("Exploring %d nodes\n", len(toExplore))
	}

	fmt.Printf("Roots: %s\n", roots)

	partOneResult := 0
	partTwoResult := 0

	toExplore = make([]*Node, 0)
	toExplore = append(toExplore, roots...)
	terminals := make([]*Node, 0)
	for len(toExplore) > 0 {
		node := toExplore[0]
		toExplore = toExplore[1:]

		if len(node.Children) == 0 {
			terminals = append(terminals, node)
		} else {
			toExplore = append(toExplore, node.Children...)
		}
	}

	shortestPathLength := math.MaxInt
	longestPathLength := 0
	for _, terminal := range terminals {
		fmt.Printf("%s\n", terminal.Path())
		if terminal.length() == len(locations) && terminal.totalDistance() < shortestPathLength {
			shortestPathLength = terminal.totalDistance()
		}
		if terminal.length() == len(locations) && terminal.totalDistance() > longestPathLength {
			longestPathLength = terminal.totalDistance()
		}
	}

	partOneResult = shortestPathLength
	partTwoResult = longestPathLength

	fmt.Printf("Part one result: %d\n", partOneResult)
	fmt.Printf("Part two result: %d\n", partTwoResult)
}

func distance(legs []Leg, paris string, london string) int {
	for _, leg := range legs {
		if (paris == leg.Start && london == leg.End) || (paris == leg.End && london == leg.Start) {
			return leg.Distance
		}
	}

	return -1
}

func parse(line string) Leg {
	locations, distanceString := get_input.SplitPair(line, " = ")
	distance, err := strconv.Atoi(distanceString)
	if err != nil {
		log.Fatalf("Could not parse distance from line %s\n", line)
	}
	start, end := get_input.SplitPair(locations, " to ")
	return Leg{
		Start:    start,
		End:      end,
		Distance: distance,
	}
}

type Leg struct {
	Start    string
	End      string
	Distance int
}

type Node struct {
	Location string
	Distance int
	Parent   *Node
	Children []*Node
}

func (node Node) totalDistance() int {
	total := node.Distance
	if node.Parent != nil {
		total += node.Parent.totalDistance()
	}

	return total
}

func (node Node) hasAncestor(location string) bool {
	if node.Location == location {
		return true
	}
	if node.Parent == nil {
		return false
	}
	return node.Parent.hasAncestor(location)
}

func (node Node) length() int {
	if node.Parent == nil {
		return 1
	}
	return 1 + node.Parent.length()
}

func (node Node) Path() string {
	if node.Parent == nil {
		return node.Location
	}
	return fmt.Sprintf("%s - %s", node.Parent.Path(), node.Location)
}

func (node Node) String() string {
	return fmt.Sprintf("{%s, %d, %v}", node.Location, node.Distance, node.Children)
}
