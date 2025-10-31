package transform

import "strings"

// TokenType represents the category of a token (word, punctuation, directive, etc.)
type TokenType int

const (
	Word TokenType = iota
	Number
	Directive
	Punctuation
	Space
	Unknown
)

// Token represents a unit of text — a word, symbol, or directive.
type Token struct {
	Value string    // actual text, e.g. "hello" or "(up,2)"
	Type  TokenType // category of token
}

// NewToken is a simple constructor
func NewToken(value string, t TokenType) Token {
	return Token{
		Value: value,
		Type:  t,
	}
}

// IsDirective checks if a token is a transformation directive like (up), (hex), etc.
func (t Token) IsDirective() bool {
	return strings.HasPrefix(t.Value, "(") && strings.HasSuffix(t.Value, ")")
}

// IsWord checks if token is a normal text word (not punctuation or directive)
func (t Token) IsWord() bool {
	return t.Type == Word
}

// Clone returns a copy (useful in pipeline stages)
func (t Token) Clone() Token {
	return Token{
		Value: t.Value,
		Type:  t.Type,
	}
}
