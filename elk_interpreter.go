package main

import "fmt"

func InterpretPrint(tokens []Token) {
	tokens = tokens[1:]
	for _, token := range tokens {
		if token.kind != "TT_COMMA" {
			fmt.Printf("%s", token.literal)
		}
	}
	fmt.Printf("\n")
}

func InterpretMaths(tokens []Token) {
	/*switch tokens[0].kind {
	case "TT_ADD":
		tokens = tokens[1:]
		for _, token := range tokens {

		}
	case "TT_SUB":
	case "TT_MUL":
	case "TT_DIV":

	}*/
}

func Interpret(tokens []Token) {
	if len(tokens) < 1 {
		return
	}

	switch tokens[0].kind {
	case "TT_ADD", "TT_SUB", "TT_MUL", "TT_DIV":
		InterpretMaths(tokens)
	case "TT_PRINT":
		InterpretPrint(tokens)
	}
}
