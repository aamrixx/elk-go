package main

import (
	"fmt"
	"os"
)

func Die(reason string, lineCount uint64) {
	fmt.Println("\x1b[0;31mERROR :\x1b[0m", reason)
	if lineCount != 0 {
		fmt.Println("On line : ", lineCount)
	}
	os.Exit(1)
}

func DieExtended(reason string, tokens []Token, lineCount uint64) {
	if len(tokens) < 1 {
		fmt.Print("\x1b[0;31mERROR : \x1b[0m", reason, "'", tokens[0].literal, "'")
	} else {
		fmt.Print("\x1b[0;31mERROR : \x1b[0m", reason, "'", TokenToString(tokens), "'")
	}

	if lineCount != 0 {
		fmt.Printf(" -> Line %d\n", lineCount)
	}

	os.Exit(1)
}

func DebugPrint(text string) {
	fmt.Println("\x1b[0;32mDebug : \x1b[0m", text)
}
