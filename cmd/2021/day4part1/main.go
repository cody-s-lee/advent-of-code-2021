package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func newBoard() *[5][5]int {
	var board [5][5]int
	return &board
}

func main() {
	lines := get_input.Lines("https://adventofcode.com/2021/day/4/input")

	calls := make([]int, 0)
	for _, s := range strings.Split(lines[0], ",") {
		x, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			log.Fatalf("Could not convert \"%s\" when getting calls\n", s)
		}
		calls = append(calls, x)
	}

	fmt.Printf("Calls: %v\n", calls)

	boards := make([]*[5][5]int, 0) // boards are always 5x5

	var board *[5][5]int // boards are always 5x5
	i := 0
	for n, s := range lines[1:] {
		s = strings.TrimSpace(s)
		fmt.Printf("Looking at line %d, \"%s\"\n", n, s)
		if s == "" { // prep a new board on newline
			board = newBoard()
			boards = append(boards, board)
			i = 0 // reset line
		} else { // append all values to the current board and line
			s = strings.ReplaceAll(s, "  ", " ")
			for j, t := range strings.Split(s, " ") {
				t = strings.TrimSpace(t)
				fmt.Printf("Looking at item %d, \"%s\"\n", j, t)
				x, err := strconv.Atoi(strings.TrimSpace(t))
				if err != nil {
					log.Fatalf("Could not convert \"%s\" when reading in board line %d, position %d\n", t, i, j)
				}
				board[i][j] = x
				fmt.Printf("Current board %v\n", board)
			}
			i++ // increment line
		}
	}

	fmt.Printf("Boards: %v\n", boards)

	leastTurns := len(calls)
	bestScore := 0

	for _, board := range boards {
		turns, score := score(*board, calls)
		if turns == 0 {
			// no win
			continue
		} else if turns < leastTurns {
			// winner
			leastTurns = turns
			bestScore = score
		} else if turns == leastTurns {
			// verify via score
			if score > bestScore {
				bestScore = score
			}
		}
	}

	fmt.Printf("Got score %d in %d turns\n", bestScore, leastTurns)
}

func contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func subset(a []int, b []int) bool {
	for _, x := range b {
		if !contains(a, x) {
			return false
		}
	}
	return true
}

func score(board [5][5]int, calls []int) (int, int) {
	turns, score := 0, 0

	for n := 4; n < len(calls); n++ {
		usedCalls := calls[0:n]
		winner := false

		// check for winner
		// rows
		for i := 0; i < 5; i++ {
			if subset(usedCalls, board[i][:]) {
				winner = true
				break
			}
		}

		// columns
		for j := 0; j < 5; j++ {
			var col [5]int
			for i := 0; i < 5; i++ {
				col[i] = board[i][j]
			}
			if subset(usedCalls, col[:]) {
				winner = true
				break
			}
		}

		// calculate and return score
		if winner {
			turns = n + 1
			score = 0
			for r := 0; r < 5; r++ {
				for s := 0; s < 5; s++ {
					if !contains(usedCalls, board[r][s]) {
						score += board[r][s]
					}
				}
			}
			score = score * usedCalls[len(usedCalls)-1]
			break
		}
	}

	return turns, score
}
