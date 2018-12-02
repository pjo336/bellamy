package parser

import (
	"bellamy/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvokeLiteral(t *testing.T) {
	// Should be <expression><period><functionLiteral><call>
	input := `"string".length();`
	program := SetupParserTest(t, input)

	assert.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	exp, ok := stmt.Expression.(*ast.InvokeExpression)
	assert.True(t, ok)

	// Make sure subject is a string literal
	literal, ok := exp.Subject.(*ast.StringLiteral)
	assert.True(t, ok)
	assert.Equal(t, "string", literal.Value)

	testIdentifier(t, exp.CallExpression.Function, "length")
	assert.Equal(t, 0, len(exp.CallExpression.Arguments))
}
