package transform

// ApplyFormat normalizes spaces around punctuation and quotes.
func ApplyFormat(tokens []Token) []Token {
	// 1) πρώτα διορθώνουμε τα σημεία στίξης
	toks := attachPunctuation(tokens)

	// 2) μετά διορθώνουμε τα quotes
	toks = fixQuotes(toks)

	// 3) συμπίεση κενών (ΝΕΟ)
	toks = compressSpaces(toks)
	return toks
}

// attachPunctuation makes punctuation stick to the previous token
// and ensures space after it when needed.
func attachPunctuation(tokens []Token) []Token {
	out := make([]Token, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		// 1) Ό,τι δεν είναι punctuation περνάει όπως είναι
		if t.Type != Punctuation {
			out = append(out, t)
			continue
		}

		// 2) ΠΡΙΝ βάλουμε punctuation, καθάρισε ΟΛΑ τα προηγούμενα spaces
		//    ΕΚΤΟΣ από την ειδική περίπτωση ": " πριν από opening quote (')
		if t.Value == "'" {
			// διατήρησε το space αν ακριβώς πριν είναι ":" + space
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

		// 3) Πρόσθεσε το τρέχον punctuation
		out = append(out, t)

		// 4) Ειδικός κανόνας για το ":" → ΠΑΝΤΑ ακριβώς ένα space μετά
		if t.Value == ":" {
			// αν το επόμενο token είναι Space → κατανάλωσέ το
			if i+1 < len(tokens) && tokens[i+1].Type == Space {
				i++
			}
			// βάλε ΑΚΡΙΒΩΣ ένα space μετά από ':'
			out = append(out, Token{Value: " ", Type: Space})
			// προχώρα στον επόμενο κύκλο (μην εφαρμόσεις άλλο spacing logic)
			continue
		}

		// 5) Αν υπάρχει επόμενο token, δες πώς θα χειριστείς spacing/ιδιαιτερότητες
		if i+1 < len(tokens) {
			next := tokens[i+1]

			// (A) "?." ή "!." → πέτα την τελεία
			if (t.Value == "?" || t.Value == "!") && next.Type == Punctuation && next.Value == "." {
				i++ // skip "."
				continue
			}

			// (B) Γενικός κανόνας spacing ανάλογα με το επόμενο token
			if next.Type == Space {
				// Αν μετά το space ακολουθεί punctuation ΚΑΙ ΔΕΝ είναι opening quote,
				// κόψε το space (π.χ. "BAMM !!")
				if i+2 < len(tokens) && tokens[i+2].Type == Punctuation && tokens[i+2].Value != "'" {
					i++ // φάε το space
					continue
				}
				// αλλιώς κράτα ΕΝΑ space
				i++
				out = append(out, Token{Value: " ", Type: Space})

			} else if next.Type == Punctuation {
				// ΕΞΑΙΡΕΣΗ: αν αμέσως μετά ακολουθεί opening quote ('), βάλε space πριν από το quote
				if next.Value == "'" {
					out = append(out, Token{Value: " ", Type: Space})
				}
				// διαφορετικά: δεν βάζουμε space (για "!!", "...", "!?")
			} else {
				// επόμενο είναι λέξη → βάλε space
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
