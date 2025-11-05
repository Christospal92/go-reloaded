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

		// if it's not directive, just we put them in the out
		if t.Type != Directive {
			out = append(out, t)
			continue
		}

		// here we know that there is directive → we have to check if there is (hex) or (bin)
		_, base, ok := parseNumberDirective(t.Value)
		if !ok {
			// they are not directive, keep it as it is
			out = append(out, t)
			continue
		}

		// find the previous Word/Number token to out
		for j := len(out) - 1; j >= 0; j-- {
			if out[j].Type == Word || out[j].Type == Number {
				// we try to do parse in the text
				dec, err := strconv.ParseInt(out[j].Value, base, 64)
				if err == nil {
					out[j].Value = strconv.FormatInt(dec, 10)
					out[j].Type = Number
				}
				break
			}
		}

		// we don't try the directive to out → we cut it down
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
