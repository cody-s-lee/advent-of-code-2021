package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	body := get_input.GetInput("https://adventofcode.com/2015/day/4/input")
	body = strings.Trim(body, " \n")

	answerShort := -1
	answerLong := -1
	for i := 0; true; i++ {
		if i%100000 == 0 {
			fmt.Printf("Testing secret of %d\n", i)
		}
		input := body + strconv.FormatInt(int64(i), 10)
		hashBytes := md5.Sum([]byte(input))
		hash := hex.EncodeToString(hashBytes[:])

		if i%100000 == 0 {
			fmt.Printf("Secret %d has hash of %s\n", i, hash)
		}

		if i%100000 == 0 {
			fmt.Printf("Prefix of %s is %s\n", hash, string(hash[0:5]))
		}
		if string(hash[0:5]) == "00000" && answerShort == -1 {
			answerShort = i
		}
		if string(hash[0:6]) == "000000" {
			answerLong = i
			break
		}
	}

	fmt.Printf("The short answer is %d\n", answerShort)
	fmt.Printf("The long answer is %d\n", answerLong)
}
