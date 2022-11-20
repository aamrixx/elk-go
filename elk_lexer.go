package main

import (
	"bytes"
)

var lexerLine string = ""
var lexerPos uint64 = 0

func LexSymbol(sym byte) Token {
	switch sym {
	case ' ':
		return Token{"", ""}
	case '#':
		return Token{"", ""}
	case '+':
		return Token{"TT_ADD", "+"}
	case '-':
		return Token{"TT_SUB", "-"}
	case '*':
		return Token{"TT_MUL", "*"}
	case '/':
		return Token{"TT_DIV", "/"}
	case ',':
		return Token{"TT_COMMA", ","}
	case '"':
		lexerPos++
		buffer := bytes.NewBufferString("")
		for lexerLine[lexerPos] != '"' {
			buffer.WriteByte(lexerLine[lexerPos])
			lexerPos++
		}
		return Token{"TT_STRING", buffer.String()}
	}
	return Token{"", ""}
}

func LexKeyword(kw string) Token {
	switch kw {
	case "":
		return Token{"", ""}
	case "variable":
		return Token{"TT_VARIABLE", "variable"}
	case "constant":
		return Token{"TT_CONSTANT", "constant"}
	case "print":
		return Token{"TT_PRINT", "print"}
	default:
		for _, c := range kw {
			if !(c >= '0' && c <= '9') && c != '.' && c != '-' {
				return Token{"TT_IDN", kw}
			}
		}
		return Token{"TT_NUM", kw}
	}
}

func Lex() []Token {
	var tokens []Token

	for int(lexerPos) < len(lexerLine) {
		token := LexSymbol(lexerLine[lexerPos])
		if !IsTokenEmpty(token) {
			tokens = append(tokens, token)
		} else {
			buffer := bytes.NewBufferString("")
			if lexerPos > 0 && lexerLine[lexerPos-1] == '-' {
				tokens = tokens[:len(tokens)-1]
				buffer.WriteByte('-')
			}

			for int(lexerPos) < len(lexerLine) && lexerLine[lexerPos] != ' ' {
				buffer.WriteByte(lexerLine[lexerPos])
				lexerPos++
			}

			token = LexKeyword(buffer.String())
			if !IsTokenEmpty(token) {
				tokens = append(tokens, token)
			}
		}
		lexerPos++
	}

	return tokens
}
