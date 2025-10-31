package transform

// ApplyFormat normalizes spaces around punctuation and quotes.
func ApplyFormat(tokens []Token) []Token {
	// 1) πρώτα διορθώνουμε τα σημεία στίξης
	toks := attachPunctuation(tokens)

	// 2) μετά διορθώνουμε τα quotes
	toks = fixQuotes(toks)

	return toks
}

// ----------------------------------------------------
// helper 1: attachPunctuation
// ----------------------------------------------------
func attachPunctuation(tokens []Token) []Token {
	out := make([]Token, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		// αν δεν είναι punctuation, απλά το περνάμε
		if t.Type != Punctuation {
			out = append(out, t)
			continue
		}

		// είναι punctuation: πρέπει να κολλήσει στο προηγούμενο
		if len(out) > 0 && out[len(out)-1].Type == Space {
			out = out[:len(out)-1]
		}
		out = append(out, t)

		// δες τι έρχεται μετά
		if i+1 < len(tokens) {
			next := tokens[i+1]

			// αν είναι space + μετά punctuation → μην βάλεις space
			if next.Type == Space {
				if i+2 < len(tokens) && tokens[i+2].Type == Punctuation {
					i++ // τρώμε το space
					continue
				}
				// κανονική περίπτωση: space μετά το punctuation
				i++
				out = append(out, Token{Value: " ", Type: Space})
			} else if next.Type == Punctuation {
				// ακολουθεί άλλο punctuation → δεν βάζουμε space
			} else {
				// ακολουθεί word → βάλε space
				out = append(out, Token{Value: " ", Type: Space})
			}
		}
	}

	return out
}

// ----------------------------------------------------
// helper 2: fixQuotes
// ----------------------------------------------------
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

		// αν έχει space αμέσως μετά το άνοιγμα → τρώμε το space
		if i < len(tokens) && tokens[i].Type == Space {
			i++
		}

		// πρόσθεσε ό,τι υπάρχει μέχρι να βρούμε closing quote
		startInner := len(out)
		for i < len(tokens) {
			// αν βρήκαμε closing quote
			if tokens[i].Type == Punctuation && tokens[i].Value == "'" {
				// αν πριν το κλείσιμο έχει space → σβήστο
				if len(out) > startInner && out[len(out)-1].Type == Space {
					out = out[:len(out)-1]
				}
				// βάλτο
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
