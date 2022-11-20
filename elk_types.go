package main

type Token struct {
	kind    string
	literal string
}

type DataStore struct {
	constant bool
	kind     string
	data     string
}

func IsTokenEmpty(token Token) bool {
	if token == (Token{"", ""}) {
		return true
	}
	return false
}

func IsDataStoreEmpty(ds DataStore) bool {
	if ds.data == "" && ds.kind == "" {
		return true
	}
	return false
}

func TokenToString(tokens []Token) string {
	if len(tokens) < 1 {
		return ""
	}

	var result string
	for _, token := range tokens {
		result += token.literal + " "
	}

	return result[:len(result)-1]
}
