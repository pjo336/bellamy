package token

import (
	"bytes"
	"log"
)

// TokenType defines the type of Token
// It uses a string to make debugging easier
// For better performance, consider a byte[]
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func FromChar(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func FromMultiChar(tokenType TokenType, chs []byte) Token {
	var buf bytes.Buffer
	_, err := buf.Write(chs)
	if err != nil {
		log.Fatalf("token.FromMultiChar: Could not write multi char string to Token Literal")
	}
	return Token{Type: tokenType, Literal: buf.String()}
}
