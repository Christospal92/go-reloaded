package transform

import (
	"strings"
	"unicode"
)

// ApplyArticles turns "a" into "an" when the next word starts with a vowel or 'h'.
// Keeps the capitalization: "A" -> "AN", "a" -> "an".
func ApplyArticles(tokens []Token) []Token {
	out := make([]Token, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		// Only for Word tokens which are "a" ή "A"
		if t.Type == Word && strings.EqualFold(t.Value, "a") {
			nextWord := findNextWord(tokens, i+1) // ignore spaces/punct/directives
			if nextWord != "" {
				first := []rune(nextWord)[0]
				if isVowelOrH(first) {
					// ✅ Keeps capital
					if t.Value == "A" {
						out = append(out, Token{Value: "AN", Type: Word})
					} else {
						out = append(out, Token{Value: "an", Type: Word})
					}
					continue
				}
			}
		}

		// default: keep the token as it be
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
			// take out quotes from the word (for example "'apple" -> "apple")
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
