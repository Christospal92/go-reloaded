package transform

import (
	"strconv"
	"strings"
)

// ApplyNumbers walks through the token list and applies (hex) and (bin)
// to the previous word token. The directive token itself is removed.
func ApplyNumbers(tokens []Token) []Token {
	out := make([]Token, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		// αν ΔΕΝ είναι directive, απλώς το περνάμε στο out
		if t.Type != Directive {
			out = append(out, t)
			continue
		}

		// εδώ ξέρουμε ότι είναι directive → πρέπει να δούμε αν είναι (hex) ή (bin)
		_, base, ok := parseNumberDirective(t.Value)
		if !ok {
			// δεν είναι αριθμητικό directive, κράτα το ως έχει
			out = append(out, t)
			continue
		}

		// βρες το προηγούμενο Word/Number token στο out
		for j := len(out) - 1; j >= 0; j-- {
			if out[j].Type == Word || out[j].Type == Number {
				// προσπαθούμε να κάνουμε parse το κείμενο
				dec, err := strconv.ParseInt(out[j].Value, base, 64)
				if err == nil {
					out[j].Value = strconv.FormatInt(dec, 10)
					out[j].Type = Number
				}
				break
			}
		}

		// ΔΕΝ προσθέτουμε το directive στο out → το “τρώμε”
	}

	return out
}

// parseNumberDirective checks if a directive is (hex) or (bin)
func parseNumberDirective(s string) (name string, base int, ok bool) {
	s = strings.TrimSpace(s)
	if len(s) < 3 {
		return "", 0, false
	}
	if s[0] != '(' || s[len(s)-1] != ')' {
		return "", 0, false
	}
	body := strings.TrimSpace(s[1 : len(s)-1])
	body = strings.ToLower(body)

	switch body {
	case "hex":
		return "hex", 16, true
	case "bin":
		return "bin", 2, true
	default:
		return "", 0, false
	}
}
