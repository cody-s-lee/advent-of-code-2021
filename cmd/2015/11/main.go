package main

import (
	"fmt"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

// Part one is the travelling salesman problem. Ugh.
func main() {
	body := get_input.GetInput("https://adventofcode.com/2015/day/11/input")
	body = strings.TrimSpace(body)
	passwordOne := []byte(body)
	passwordOne = increment(passwordOne)
	for !isValid(passwordOne) {
		passwordOne = increment(passwordOne)
	}

	passwordTwo := []byte(body)
	passwordTwo = increment(passwordTwo)
	for !isValid(passwordTwo) {
		passwordTwo = increment(passwordTwo)
	}
	passwordTwo = increment(passwordTwo)
	for !isValid(passwordTwo) {
		passwordTwo = increment(passwordTwo)
	}

	fmt.Printf("Part one result: %s\n", string(passwordOne))
	fmt.Printf("Part two result: %s\n", string(passwordTwo))
}

func increment(password []byte) []byte {
	i := len(password) - 1
	done := false
	for !done {
		password[i] = password[i] + 1
		if password[i] == 'z'+1 {
			password[i] = 'a'
			i--
		} else {
			done = true
		}
	}

	return password
}

func isValid(password []byte) bool {
	hasStraight := false
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1]-1 && password[i+1] == password[i+2]-1 {
			hasStraight = true
			break
		}
	}
	if !hasStraight {
		return false
	}

	for i := 0; i < len(password); i++ {
		if password[i] == 'i' || password[i] == 'o' || password[i] == 'l' {
			return false
		}
	}

	pairsFound := 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairsFound++
			i++
		}
	}
	return pairsFound >= 2
}
