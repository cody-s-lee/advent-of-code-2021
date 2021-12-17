package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

// Part one is the travelling salesman problem. Ugh.
func main() {
	body := get_input.Body("https://adventofcode.com/2015/day/12/input")
	body = strings.TrimSpace(body)
	numbers := make([]int, 0)

	n := 0
	m := 1
	for _, r := range body {
		if r == '-' {
			m = -1
		} else if r >= '0' && r <= '9' {
			x, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatalf("Could not convert %s\n", string(r))
			}
			n = 10*n + x
		} else {
			n = n * m
			if n != 0 {
				numbers = append(numbers, n)
			}
			n = 0
			m = 1
		}
	}

	partOneResult := 0
	for _, n := range numbers {
		partOneResult += n
	}

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(body), &jsonMap)
	if err != nil {
		log.Fatalf("Could not unmarshal %s\n", err)
	}

	toExplore := make([]map[string]interface{}, 0)
	toExplore = append(toExplore, jsonMap)

	partTwoResult := 0

	for len(toExplore) > 0 {
		node := toExplore[0]
		toExplore = toExplore[1:]

		nodeValue := 0
		nodeExplore := make([]map[string]interface{}, 0)
		isRed := false
		for _, v := range node {
			if v == "red" {
				isRed = true
				break
			}
			switch w := v.(type) {
			case float64:
				nodeValue += int(w)
			case []interface{}:
				array := w
				for len(array) > 0 {
					q := array[0]
					array = array[1:]
					switch r := q.(type) {
					case float64:
						nodeValue += int(r)
					case []interface{}:
						array = append(array, r...)
					case map[string]interface{}:
						nodeExplore = append(nodeExplore, r)
					}
				}
			case map[string]interface{}:
				nodeExplore = append(nodeExplore, w)
			}
		}
		if !isRed {
			toExplore = append(toExplore, nodeExplore...)
			partTwoResult += nodeValue
		}
	}

	fmt.Printf("Part one result: %d\n", partOneResult)
	fmt.Printf("Part two result: %d\n", partTwoResult)
}
