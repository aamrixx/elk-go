package main

import (
	"fmt"
	"unicode"
)

var lineCount uint64 = 1
var dataStoreMap = make(map[string]DataStore)

func ParseIdentifer(token Token) {
	for i := 0; i < len(token.literal); i++ {
		c := rune(token.literal[i])
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) && c != '_' {
			DieExtended("Invalid identifier literal -> Invalid character ", []Token{token}, lineCount)
		}
	}
}

func ParseNumber(token Token) {
	i := 0
	if token.literal[0] == '-' {
		i = 1
	}

	var dotCount uint64
	for _ = i; i < len(token.literal); i++ {
		c := token.literal[i]

		if c == '.' {
			dotCount++
		}

		if !(c >= '0' && c <= '9') && c != '.' {
			DieExtended("Invalid number literal -> ", []Token{token}, lineCount)
		}
	}

	if token.literal[len(token.literal)-1] == '.' {
		DieExtended("Invalid number literal -> Numbers can not end in decimal points", []Token{token}, lineCount)
	}

	if dotCount > 1 {
		DieExtended("Invalid number literal -> Too many decimal points ", []Token{token}, lineCount)
	}
}

func ParseVariable(tokens []Token) {
	if len(tokens) != 3 {
		DieExtended("Invalid variable statement -> ", tokens, lineCount)
	}

	ParseIdentifer(tokens[1])

	if !IsDataStoreEmpty(dataStoreMap[tokens[1].literal]) {
		DieExtended("Redefined variable -> ", tokens, lineCount)
	}

	if tokens[2].kind == "TT_IDN" {
		ParseIdentifer(tokens[2])

		variable := dataStoreMap[tokens[2].literal]
		if IsDataStoreEmpty(variable) {
			DieExtended("Undefined variable -> ", []Token{tokens[2]}, lineCount)
		}

		tokens = tokens[:len(tokens)-1]
		tokens = append(tokens, Token{variable.kind, variable.data})
	} else if tokens[2].kind != "TT_STRING" {
		ParseNumber(tokens[2])
	}

	dataStoreMap[tokens[1].literal] = DataStore{false, tokens[2].kind, tokens[2].literal}
}

func ParseConstant(tokens []Token) {
	if len(tokens) != 3 {
		DieExtended("Invalid constant statement -> ", tokens, lineCount)
	}

	ParseIdentifer(tokens[1])

	if !IsDataStoreEmpty(dataStoreMap[tokens[1].literal]) {
		DieExtended("Redefined variable -> ", tokens, lineCount)
	}

	if tokens[2].kind == "TT_IDN" {
		ParseIdentifer(tokens[2])

		variable := dataStoreMap[tokens[2].literal]
		if IsDataStoreEmpty(variable) {
			DieExtended("Undefined variable -> ", []Token{tokens[2]}, lineCount)
		}

		tokens = tokens[:len(tokens)-1]
		tokens = append(tokens, Token{variable.kind, variable.data})
	} else if tokens[2].kind != "TT_STRING" {
		ParseNumber(tokens[2])
	}

	dataStoreMap[tokens[1].literal] = DataStore{true, tokens[2].kind, tokens[2].literal}
}

func ParseRedefine(tokens []Token) {
	if len(tokens) != 2 {
		DieExtended("Invalid redefine statement -> ", tokens, lineCount)
	}

	ParseIdentifer(tokens[0])

	if IsDataStoreEmpty(dataStoreMap[tokens[0].literal]) {
		DieExtended("Undefined variable -> ", tokens, lineCount)
	}

	variable := dataStoreMap[tokens[0].literal]
	if IsDataStoreEmpty(variable) {
		DieExtended("Undefined variable -> ", []Token{tokens[0]}, lineCount)
	}

	if variable.constant {
		DieExtended("Variable is constant -> Can not change constants ", tokens, lineCount)
	}

	switch tokens[i].kind {
	case "TT_IDN":
	case "TT_ADD", "TT_SUB", "TT_MUL", "TT_DIV":
		ParseMaths(tokens)
	case "TT_NUM":
	}

	if tokens[1].kind == "TT_IDN" {
		ParseIdentifer(tokens[1])

		variable := dataStoreMap[tokens[1].literal]
		if IsDataStoreEmpty(variable) {
			DieExtended("Undefined variable -> ", []Token{tokens[0]}, lineCount)
		}

		tokens = tokens[:len(tokens)-1]
		tokens = append(tokens, Token{variable.kind, variable.data})
	} else if tokens[1].kind == "" {

	} else if tokens[1].kind == "TT_NUM" {
		ParseNumber(tokens[1])
	}

	dataStoreMap[tokens[0].literal] = DataStore{false, tokens[1].kind, tokens[1].literal}
}

func ParsePrint(tokens []Token) {
	for i := 0; i < len(tokens); i++ {
		if tokens[i].kind == "TT_STRING" {
			if i+1 != len(tokens) && tokens[i+1].kind != "TT_COMMA" {
				DieExtended("Expected comma -> ", tokens, lineCount)
			}
		} else if tokens[i].kind == "TT_IDN" {
			ParseIdentifer(tokens[i])

			variable := dataStoreMap[tokens[i].literal]
			if IsDataStoreEmpty(variable) {
				DieExtended("Undefined variable -> ", []Token{tokens[i]}, lineCount)
			}

			if i+1 != len(tokens) && tokens[i+1].kind != "TT_COMMA" {
				DieExtended("Expected comma -> ", tokens, lineCount)
			}

			tokens[i] = Token{variable.kind, variable.data}
		} else if tokens[i].kind == "TT_NUM" {
			ParseNumber(tokens[i])
			if i+1 != len(tokens) && tokens[i+1].kind != "TT_COMMA" {
				DieExtended("Expected comma -> ", tokens, lineCount)
			}
		} else if tokens[i].kind == "TT_COMMA" {
			if i+1 != len(tokens) && tokens[i+1].kind == "TT_COMMA" {
				DieExtended("Expected string, number or identifier -> ", tokens, lineCount)
			}
		}
	}
}

func ParseMaths(tokens []Token) {
	for i := 0; i < len(tokens); i++ {
		if tokens[i].kind == "TT_NUM" {
			ParseNumber(tokens[i])
		} else if tokens[i].kind == "TT_IDN" {
			ParseIdentifer(tokens[i])

			variable := dataStoreMap[tokens[i].literal]
			if IsDataStoreEmpty(variable) {
				DieExtended("Undefined variable -> ", []Token{tokens[i]}, lineCount)
			}

			tokens[i] = Token{variable.kind, variable.data}
		}
	}
}

func Parse(tokens []Token) {
	if len(tokens) < 1 {
		return
	}

	lineCount++

	switch tokens[0].kind {
	case "TT_ADD", "TT_SUB", "TT_MUL", "TT_DIV":
		ParseMaths(tokens)
	case "TT_VARIABLE":
		ParseVariable(tokens)
	case "TT_CONSTANT":
		ParseConstant(tokens)
	case "TT_PRINT":
		ParsePrint(tokens)
	case "TT_IDN":
		ParseRedefine(tokens)
	default:
		DieExtended("Unknown literal -> ", []Token{tokens[0]}, lineCount)
	}

	for i := 0; i < len(tokens); i++ {
		fmt.Println(tokens[i])
	}

	fmt.Println(dataStoreMap)
}
