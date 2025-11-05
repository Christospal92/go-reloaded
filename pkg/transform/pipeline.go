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

	// 3) remove any leftover directives (safety) — before the format
	clean := make([]Token, 0, len(tokens))
	for _, tk := range tokens {
		if tk.Type == Directive {
			continue
		}
		clean = append(clean, tk)
	}

	// 4) format (punctuation, quotes, spaces)
	clean = ApplyFormat(clean)

	// 5) turn tokens back into string
	return Detokenize(clean)
}

// Detokenize converts tokens back into a string.
// Write exactly one space for every Space token
// and just connects the rest as they are.
func Detokenize(tokens []Token) string {
	var sb strings.Builder

	for _, tk := range tokens {
		if tk.Type == Space {
			// write only one space
			sb.WriteByte(' ')
			continue
		}
		sb.WriteString(tk.Value)
	}

	return strings.TrimSpace(sb.String())
}
