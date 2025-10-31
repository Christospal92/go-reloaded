package transform

import (
	"unicode"
)

func isPunctuation(r rune) bool {
	switch r {
	case '.', ',', '!', '?', ':', ';', '\'':
		return true
	default:
		return false
	}
}

// Tokenize breaks the input string into a slice of Tokens.
// Rules:
// - sequences of spaces -> one Space token
// - directives like (up), (up, 2), (hex), (bin) -> Directive token
// - single punctuation chars -> Punctuation token
// - everything else -> Word token
func Tokenize(input string) []Token {
	var tokens []Token

	runes := []rune(input)
	n := len(runes)
	i := 0

	for i < n {
		r := runes[i]

		// 1) spaces
		if unicode.IsSpace(r) {
			// collapse consecutive spaces/newlines/tabs into one Space token
			j := i + 1
			for j < n && unicode.IsSpace(runes[j]) {
				j++
			}
			tokens = append(tokens, NewToken(" ", Space))
			i = j
			continue
		}

		// 2) directive: starts with '(' and ends with ')'
		if r == '(' {
			j := i + 1
			for j < n && runes[j] != ')' {
				j++
			}
			if j < n && runes[j] == ')' {
				text := string(runes[i : j+1])
				// we accept anything in parentheses as directive for now
				tokens = append(tokens, NewToken(text, Directive))
				i = j + 1
				continue
			}
			// if there is '(' but no closing, fallthrough as word
		}

		// 3) punctuation
		if isPunctuation(r) {
			tokens = append(tokens, NewToken(string(r), Punctuation))
			i++
			continue
		}

		// 4) word: read until space, punctuation or '('
		j := i + 1
		for j < n && !unicode.IsSpace(runes[j]) && !isPunctuation(runes[j]) && runes[j] != '(' {
			j++
		}
		word := string(runes[i:j])
		tokens = append(tokens, NewToken(word, Word))
		i = j
	}

	return tokens
}
