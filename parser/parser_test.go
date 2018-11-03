package parser

import (
	"bellamy/ast"
	"bellamy/lexer"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 1234;
	`
	numStatements := 3
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p, false)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, numStatements, len(program.Statements), "program.Statements does not contain %d statements. got %d statements", numStatements, len(program.Statements))

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tt.expectedIdentifier)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) {
	assert.Equal(t, "let", s.TokenLiteral(), "s.TokenLiteral not 'let'. got %q", s.TokenLiteral())
	letStmt, ok := s.(*ast.LetStatement)
	assert.True(t, ok, "letStmt.Name.Value not '%s'. got %s", name, letStmt.Name.Value)
	assert.Equal(t, name, letStmt.Name.TokenLiteral(), "s.Name not '%s'. got %s", name, letStmt.Name)
}

func TestParserErrors(t *testing.T) {
	input := `
	let x 5;
	`
	l := lexer.New(input)
	p := New(l)

	p.ParseProgram()
	// we expect errors here
	es := checkParserErrors(t, p, true)
	assert.Equal(t, "expected next token to be =, got INT", es[0])
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 12345;
	return add(x, y);
	`
	numStatements := 3
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p, false)

	assert.Equal(t, numStatements, len(program.Statements))

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		assert.True(t, ok, "statement is not ReturnStatement, got %T", stmt)
		assert.Equal(t, "return", returnStmt.TokenLiteral())
	}
}

func TestIdentifierExpressions(t *testing.T) {
	input := "foobar"
	numStatements := 1
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p, false)

	assert.Equal(t, numStatements, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "program.Statements[0] is not ast.ExpressionStatement, got %T", program.Statements[0])

	ident, ok := stmt.Expression.(*ast.Identifier)
	assert.True(t, ok, "exp not *ast.Identifier, got %T", stmt.Expression)
	assert.Equal(t, "foobar", ident.Value)
	assert.Equal(t, "foobar", ident.TokenLiteral())
}

func TestBoolean(t *testing.T) {
	input := "true;"
	numStatements := 1
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p, false)

	assert.Equal(t, numStatements, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "program.Statements[0] is not ast.ExpressionStatement, got %T", program.Statements[0])

	b, ok := stmt.Expression.(*ast.Boolean)
	assert.True(t, ok, "exp not *ast.Boolean, got %T", stmt.Expression)
	assert.Equal(t, true, b.Value)
	assert.Equal(t, "true", b.TokenLiteral())
}

func TestIntegerLiteralExpressions(t *testing.T) {
	input := "5;"
	numStatements := 1
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p, false)
	assert.Equal(t, numStatements, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "program.Statements[0] is not expression statement, got %T", program.Statements[0])
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	assert.True(t, ok, "exp not IntegerLiteral, got %T", stmt.Expression)
	assert.Equal(t, int64(5), literal.Value)
	assert.Equal(t, "5", literal.TokenLiteral())
}

func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue interface{}
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
		{"!true;", "!", true},
		{"!false;", "!", false},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p, false)

		assert.Equal(t, 1, len(program.Statements))

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(t, ok, "program.Statements[0] is not expression statement, got %T", program.Statements[0])

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		assert.True(t, ok, "stmt not PrefixExpression, got %T", stmt.Expression)
		assert.Equal(t, tt.operator, exp.Operator)
		testLiteralExpression(t, exp.Right, tt.integerValue)
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p, false)

		assert.Equal(t, 1, len(program.Statements))

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(t, ok, "program.Statements[0] is not expression statement, got %T", program.Statements[0])

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		assert.True(t, ok, "stmt not InfixExpression, got %T", stmt.Expression)
		testInfixExpression(t, exp, tt.leftValue, tt.operator, tt.rightValue)
		testLiteralExpression(t, exp.Left, tt.leftValue)

		assert.Equal(t, tt.operator, exp.Operator)
		testLiteralExpression(t, exp.Right, tt.rightValue)
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"(5 + 5) * 2 * (5 + 5)",
			"(((5 + 5) * 2) * (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p, false)

		actual := program.String()
		assert.Equal(t, tt.expected, actual)
	}
}

// Helper methods

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) {
	switch v:= expected.(type) {
	case int:
		testIntegerLiteral(t, exp, int64(v))
	case int64:
		testIntegerLiteral(t, exp, v)
	case string:
		testIdentifier(t, exp, v)
	case bool:
		testBooleanLiteral(t, exp, v)
	default:
		t.Errorf("Type of exp not handled, got %T", v)
	}
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) {
	b, ok := exp.(*ast.Boolean)
	assert.True(t, ok, "Expected boolean, got %t", exp)
	assert.Equal(t, value, b.Value)
	assert.Equal(t, fmt.Sprintf("%t", value), b.TokenLiteral())
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, op string, right interface{}) {
	opExp, ok := exp.(*ast.InfixExpression)
	assert.True(t, ok, "exp is not ast.InfixExpression, got %T", exp)
	testLiteralExpression(t, opExp.Left, left)
	assert.Equal(t, op, opExp.Operator)
	testLiteralExpression(t, opExp.Right, right)
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) {
	i, ok := il.(*ast.IntegerLiteral)
	assert.True(t, ok, "il not *ast.IntegerLiteral, got %T", il)
	assert.Equal(t, value, i.Value)
	assert.Equal(t, fmt.Sprintf("%d", value), i.TokenLiteral())
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) {
	ident, ok := exp.(*ast.Identifier)
	assert.True(t, ok, "exp not *ast.Identifier, got %T", exp)
	assert.Equal(t, value, ident.Value)
	assert.Equal(t, value, ident.TokenLiteral())
}

func checkParserErrors(t *testing.T, p *Parser, expected bool) []string {
	errors := p.Errors()
	if expected && len(errors) > 0 {
		return errors
	}
	if !expected {
		checkErrors(t, errors)
		return nil
	}
	return nil
}

func checkErrors(t *testing.T, errors []string) {
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser had %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
