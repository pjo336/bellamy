package evaluator

import (
	"bellamy/lexer"
	"bellamy/object"
	"bellamy/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct{
		input string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct{
		input string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testBooleanObject(t *testing.T, o object.Object, expected bool) {
	result, ok := o.(*object.Boolean)
	assert.True(t, ok)
	assert.Equal(t, expected, result.Value)
}

func testIntegerObject(t *testing.T, o object.Object, expected int64) {
	result, ok := o.(*object.Integer)
	assert.True(t, ok)
	assert.Equal(t, expected, result.Value)
}
