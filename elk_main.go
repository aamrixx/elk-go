package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		Die("Invalid arguments", 0)
	}

	file, err := os.Open(args[1])
	if err != nil {
		Die("Failed to open the file", 0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lexerPos = 0
		lexerLine = scanner.Text()
		tokens := Lex()
		Parse(tokens)
		Interpret(tokens)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
