package ast

import (
	"bellamy/token"
	"bytes"
	"strings"
)

type InvokeExpression struct {
	Subject        Expression
	Token          token.Token
	CallExpression *CallExpression
}

func (ie *InvokeExpression) expressionNode() {}

func (ie *InvokeExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InvokeExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ie.CallExpression.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ie.Subject.String())
	out.WriteString(ie.CallExpression.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}
