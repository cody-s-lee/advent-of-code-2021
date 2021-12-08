package main

import (
	"fmt"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	lines := get_input.Lines("https://adventofcode.com/2015/day/8/input")

	totalCodeChars := 0
	totalMemoryChars := 0
	totalEncChars := 0

	for _, line := range lines {
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fmt.Println("Looking at line:")
		fmt.Println(line)

		totalCodeChars += len(line) // count all characters for code

		fmt.Printf("Found %d code characters\n", len(line))

		memLine := line[1 : len(line)-1]                   // get rid of starting and ending quote
		memLine = strings.ReplaceAll(memLine, "\\\\", "/") // discount escaped backslashes
		memLine = strings.ReplaceAll(memLine, "\\\"", "/") // discount escaped quotes

		fmt.Println("Interim line:")
		fmt.Println(memLine)

		memoryChars := len(memLine) - strings.Count(memLine, "\\x")*3 // discount ascii escapes by 3
		fmt.Printf("Found %d memory characters\n", memoryChars)
		totalMemoryChars += memoryChars

		encLine := line[:]
		encLine = strings.ReplaceAll(encLine, "\\", "\\\\") // encode all backslashes
		encLine = strings.ReplaceAll(encLine, "\"", "\\\"") // encode all quotes
		encLine = "\"" + encLine + "\""
		fmt.Println("New encoded line:")
		fmt.Println(encLine)
		totalEncChars += len(encLine)

	}

	fmt.Printf("Part one: Diff size %d characters\n", totalCodeChars-totalMemoryChars)
	fmt.Printf("Part two: Diff size %d characters\n", totalEncChars-totalCodeChars)
}
