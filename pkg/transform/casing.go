package transform

import (
	"strconv"
	"strings"
	"unicode"
)

// ApplyCasing applies casing transformations like (up), (low), (cap)
// and their numbered variants (up, 2), etc.
func ApplyCasing(tokens []Token) []Token {
	out := make([]Token, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		if t.Type != Directive {
			out = append(out, t)
			continue
		}

		// Try to parse directive
		mode, count, ok := parseCasingDirective(t.Value)
		if !ok {
			// Not a casing directive, keep as-is
			out = append(out, t)
			continue
		}

		// Apply transformation backwards to previous "count" Word tokens
		applied := 0
		for j := len(out) - 1; j >= 0 && applied < count; j-- {
			if out[j].Type == Word {
				out[j].Value = applyCaseMode(out[j].Value, mode)
				applied++
			}
		}
		// Skip adding the directive itself
	}

	return out
}

// parseCasingDirective reads a directive like "(up)" or "(cap, 3)"
// and returns (mode, count, ok)
func parseCasingDirective(s string) (string, int, bool) {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "(") || !strings.HasSuffix(s, ")") {
		return "", 0, false
	}
	body := strings.TrimSpace(s[1 : len(s)-1])
	parts := strings.Split(body, ",")
	mode := strings.ToLower(strings.TrimSpace(parts[0]))

	if mode != "up" && mode != "low" && mode != "cap" {
		return "", 0, false
	}

	count := 1
	if len(parts) == 2 {
		val := strings.TrimSpace(parts[1])
		n, err := strconv.Atoi(val)
		if err == nil && n > 0 {
			count = n
		}
	}

	return mode, count, true
}

// applyCaseMode applies one of the modes to a string.
func applyCaseMode(s, mode string) string {
	switch mode {
	case "up":
		return strings.ToUpper(s)
	case "low":
		return strings.ToLower(s)
	case "cap":
		return capitalizeWord(s)
	default:
		return s
	}
}

// capitalizeWord capitalizes only the first letter of a word.
func capitalizeWord(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(strings.ToLower(s))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
