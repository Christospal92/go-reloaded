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

	// 3) remove any leftover directives (safety) — πριν το format
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
// Γράφει ακριβώς ένα space για κάθε Space token
// και απλώς ενώνει τα υπόλοιπα όπως είναι.
func Detokenize(tokens []Token) string {
	var sb strings.Builder

	for _, tk := range tokens {
		if tk.Type == Space {
			// γράψε ένα και μόνο space
			sb.WriteByte(' ')
			continue
		}
		sb.WriteString(tk.Value)
	}

	return strings.TrimSpace(sb.String())
}
