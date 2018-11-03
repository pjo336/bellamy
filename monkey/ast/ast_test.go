package ast

import (
	"Bellamy/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVal"},
					Value: "anotherVal",
				},
			},
		},
	}

	expected := "let myVar = anotherVal;"
	if program.String() != expected {
		t.Errorf("program.String() incorrect, expected %q, got %q", expected, program.String())
	}
}
