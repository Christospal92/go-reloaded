package transform

import "strings"

// Transform is the main entrypoint for the text processing pipeline.
func Transform(input string) string {
	// 1) break input into tokens
	tokens := Tokenize(input)

	// 2) apply stages in order
	tokens = ApplyNumbers(tokens)
	tokens = ApplyCasing(tokens)
	tokens = ApplyArticles(tokens)
	tokens = ApplyFormat(tokens)

	// 3) remove any leftover directives (safety)
	clean := make([]Token, 0, len(tokens))
	for _, tk := range tokens {
		if tk.Type == Directive {
			continue
		}
		clean = append(clean, tk)
	}

	// 4) turn tokens back into string
	return Detokenize(clean)
}

// Detokenize converts tokens back into a string.
// We assume that Format has already fixed spaces around punctuation,
// so here we just join respecting Space tokens.
func Detokenize(tokens []Token) string {
	var sb strings.Builder

	for i, tk := range tokens {
		switch tk.Type {
		case Space:
			// avoid double spaces
			if sb.Len() > 0 {
				sb.WriteString(" ")
			}
		default:
			sb.WriteString(tk.Value)

			// optional: if next is Word and current is Word, add space
			if i+1 < len(tokens) && tokens[i+1].Type == Word && tk.Type == Word {
				sb.WriteString(" ")
			}
		}
	}

	return strings.TrimSpace(sb.String())
}
