package main

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestExampleSimple(t *testing.T) {
	t.Log("Foo")
	bytes, err := ioutil.ReadFile("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	body := string(bytes)

	t.Log(body)

	lines := strings.Split(body, "\n")

	dotField, foldInstructions := Parse(lines)
	t.Logf("\n%s\n", Draw(dotField))

	dotField = Fold(dotField, foldInstructions[0:1])
	t.Logf("\n%s\n", Draw(dotField))

	dotField = Fold(dotField, foldInstructions[1:])
	t.Logf("\n%s\n", Draw(dotField))
}
