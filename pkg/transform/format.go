package transform

// ApplyFormat normalizes spaces around punctuation and quotes.
func ApplyFormat(tokens []Token) []Token {
	// 1) First, we correct the punction
	toks := attachPunctuation(tokens)

	// 2) After, we correct the quotes
	toks = fixQuotes(toks)

	// 3) Compressing spaces
	toks = compressSpaces(toks)
	return toks
}

// attachPunctuation makes punctuation stick to the previous token
// and ensures space after it when needed.
func attachPunctuation(tokens []Token) []Token {
	out := make([]Token, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		// 1) Whatever is not punctuation stays as it is
		if t.Type != Punctuation {
			out = append(out, t)
			continue
		}

		// 2) Before we put punctuation, clean all the previous spaces
		//    Except from this special situation ": " before the opening quote (')
		if t.Value == "'" {
			// keep space if before it is ":" + space
			if !(len(out) >= 2 &&
				out[len(out)-1].Type == Space &&
				out[len(out)-2].Type == Punctuation && out[len(out)-2].Value == ":") {
				for len(out) > 0 && out[len(out)-1].Type == Space {
					out = out[:len(out)-1]
				}
			}
		} else {
			for len(out) > 0 && out[len(out)-1].Type == Space {
				out = out[:len(out)-1]
			}
		}

		// 3) Add the following punctuation
		out = append(out, t)

		// 4) Special rule for the ":" → Always exactly one space after
		if t.Value == ":" {
			// if the next token is Space → consume it
			if i+1 < len(tokens) && tokens[i+1].Type == Space {
				i++
			}
			// put exactly one space after the ':'
			out = append(out, Token{Value: " ", Type: Space})
			// move to the next cicle (dont use another spacing logic)
			continue
		}

		// 5) if there is anohter token, check how you will use spacing/features
		if i+1 < len(tokens) {
			next := tokens[i+1]

			// (A) "?." or "!." → remove the full stop
			if (t.Value == "?" || t.Value == "!") && next.Type == Punctuation && next.Value == "." {
				i++ // skip "."
				continue
			}

			// (B) General rule spacing depending on the next token
			if next.Type == Space {
				// If the space is followed by punctuation AND it is NOT an opening quote
				// cut the space (for example "BAMM !!")
				if i+2 < len(tokens) && tokens[i+2].Type == Punctuation && tokens[i+2].Value != "'" {
					i++ // cut the space
					continue
				}
				// otherwise keep one space
				i++
				out = append(out, Token{Value: " ", Type: Space})

			} else if next.Type == Punctuation {
				// Exception: if an opening quote (') immediately follows, add a space before the quote.
				if next.Value == "'" {
					out = append(out, Token{Value: " ", Type: Space})
				}
				// otherwise: we don't put space ( "!!", "...", "!?")
			} else {
				// next will be a word → put space
				out = append(out, Token{Value: " ", Type: Space})
			}
		}
	}

	return out
}

// fixQuotes trims spaces right after opening quote and right before closing quote.
func fixQuotes(tokens []Token) []Token {
	out := make([]Token, 0, len(tokens))

	i := 0
	for i < len(tokens) {
		t := tokens[i]

		// όχι quote → απλώς πέρασέ το
		if t.Type != Punctuation || t.Value != "'" {
			out = append(out, t)
			i++
			continue
		}

		// opening quote
		out = append(out, t)
		i++

		// If there is a space immediately after the opening → consume the space.
		if i < len(tokens) && tokens[i].Type == Space {
			i++
		}

		// πρόσθεσε ό,τι υπάρχει μέχρι να βρούμε closing quote
		startInner := len(out)
		for i < len(tokens) {
			// if we fing closing quote
			if tokens[i].Type == Punctuation && tokens[i].Value == "'" {
				// if before the closure there is space → erase it
				if len(out) > startInner && out[len(out)-1].Type == Space {
					out = out[:len(out)-1]
				}
				// put
				out = append(out, tokens[i])
				i++
				break
			}
			out = append(out, tokens[i])
			i++
		}
	}

	return out
}

// compressSpaces merges consecutive Space tokens into a single one
// and trims leading/trailing spaces.
func compressSpaces(tokens []Token) []Token {
	out := make([]Token, 0, len(tokens))
	prevSpace := false

	for _, tk := range tokens {
		if tk.Type == Space {
			// skip leading space
			if len(out) == 0 {
				continue
			}
			// collapse multiple spaces
			if prevSpace {
				continue
			}
			out = append(out, tk)
			prevSpace = true
		} else {
			out = append(out, tk)
			prevSpace = false
		}
	}
	// trim trailing space
	if len(out) > 0 && out[len(out)-1].Type == Space {
		out = out[:len(out)-1]
	}
	return out
}
