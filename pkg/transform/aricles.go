package transform

import (
	"strings"
	"unicode"
)

// ApplyArticles turns "a" into "an" when the next word starts with a vowel or 'h'.
func ApplyArticles(tokens []Token) []Token {
	out := make([]Token, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		// Εξετάζουμε μόνο Word tokens που είναι "a" (ή "A")
		if t.Type == Word && strings.EqualFold(t.Value, "a") {
			// βρες το επόμενο word token (αγνόησε spaces & punctuation & directives)
			nextWord := findNextWord(tokens, i+1)
			if nextWord != "" {
				first := []rune(nextWord)[0]
				if isVowelOrH(first) {
					// διατήρηση κεφαλαίου αν ήταν "A"
					if t.Value == "A" {
						out = append(out, Token{Value: "An", Type: Word})
					} else {
						out = append(out, Token{Value: "an", Type: Word})
					}
					continue
				}
			}
		}

		// default: κράτα το token ως έχει
		out = append(out, t)
	}

	return out
}

// findNextWord returns the Value of the next Word token after index "start".
// It skips spaces, punctuation and directives.
func findNextWord(tokens []Token, start int) string {
	for j := start; j < len(tokens); j++ {
		switch tokens[j].Type {
		case Word:
			// μπορεί να έχει μπροστά quote, π.χ. "'apple"
			return trimLeadingQuotes(tokens[j].Value)
		case Space, Punctuation, Directive:
			continue
		default:
			continue
		}
	}
	return ""
}

func trimLeadingQuotes(s string) string {
	for len(s) > 0 {
		r := rune(s[0])
		if r == '\'' || r == '"' {
			s = s[1:]
			continue
		}
		break
	}
	return s
}

func isVowelOrH(r rune) bool {
	r = unicode.ToLower(r)
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h'
}
