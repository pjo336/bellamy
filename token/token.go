package token

// TokenType defines the type of Token
// It uses a string to make debugging easier
// For better performance, consider a byte[]
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
