package parser

import (
	"bellamy/ast"
	"bellamy/lexer"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 1234;
	`
	numStatements := 3
	program := setupTest(t, input)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, numStatements, len(program.Statements), "program.Statements does not contain %d statements. got %d statements", numStatements, len(program.Statements))

	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", true},
		{"let foobar = abc", "foobar", "abc"},
	}

	for _, tt := range tests {
		program := setupTest(t, tt.input)

		assert.Equal(t, 1, len(program.Statements))
		stmt := program.Statements[0]
		testLetStatement(t, stmt, tt.expectedIdentifier)
		val := stmt.(*ast.LetStatement).Value
		testLiteralExpression(t, val, tt.expectedValue)
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

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"return 5;", 5},
		{"return true;", true},
		{"return foobar;", "foobar"},
	}

	for _, tt := range tests {
		program := setupTest(t, tt.input)

		assert.Equal(t, 1, len(program.Statements))

		stmt := program.Statements[0]
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		assert.True(t, ok)
		assert.Equal(t, "return", returnStmt.TokenLiteral())
		testLiteralExpression(t, returnStmt.ReturnValue, tt.expectedValue)
	}
}

func TestIdentifierExpressions(t *testing.T) {
	input := "foobar"
	numStatements := 1
	program := setupTest(t, input)

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
	program := setupTest(t, input)

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
	program := setupTest(t, input)
	assert.Equal(t, numStatements, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "program.Statements[0] is not expression statement, got %T", program.Statements[0])
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	assert.True(t, ok, "exp not IntegerLiteral, got %T", stmt.Expression)
	assert.Equal(t, int64(5), literal.Value)
	assert.Equal(t, "5", literal.TokenLiteral())
}

func TestStringLiteralExpressions(t *testing.T) {
	input := `"this is a string!"`
	numStatements := 1
	program := setupTest(t, input)

	assert.Equal(t, numStatements, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "program.Statements[0] is not expression statement, got %T", program.Statements[0])

	literal, ok := stmt.Expression.(*ast.StringLiteral)
	assert.True(t, ok, "exp not StringLiteral, got %T", stmt.Expression)
	assert.Equal(t, "this is a string!", literal.Value)
	assert.Equal(t, "this is a string!", literal.TokenLiteral())
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
		program := setupTest(t, tt.input)

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
		program := setupTest(t, tt.input)
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
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
		{
			"a * [1, 2, 3, 4][b * c] * d",
			"((a * ([1, 2, 3, 4][(b * c)])) * d)",
		},
		{
			"add(a * b[2], b[1], 2 * [1, 2][1])",
			"add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))",
		},
	}

	for _, tt := range tests {
		program := setupTest(t, tt.input)
		actual := program.String()
		assert.Equal(t, tt.expected, actual)
	}
}

func TestIfExpression(t *testing.T) {
	input := "if (x < y) { x };"
	numStatements := 1
	program := setupTest(t, input)

	assert.Equal(t, numStatements, len(program.Statements))
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	exp, ok := stmt.Expression.(*ast.IfExpression)
	assert.True(t, ok)
	testInfixExpression(t, exp.Condition, "x", "<", "y")
	assert.Equal(t, numStatements, len(exp.Consequence.Statements))
	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	testIdentifier(t, consequence.Expression, "x")
	assert.Nil(t, exp.Alternative)
}

func TestIfElseExpression(t *testing.T) {
	input := `if (x < y) { x } else { y }`
	numStatements := 1
	program := setupTest(t, input)

	assert.Equal(t, numStatements, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	exp, ok := stmt.Expression.(*ast.IfExpression)
	assert.True(t, ok)
	testInfixExpression(t, exp.Condition, "x", "<", "y")
	assert.Equal(t, numStatements, len(exp.Consequence.Statements))

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	testIdentifier(t, consequence.Expression, "x")
	assert.Equal(t, numStatements, len(exp.Alternative.Statements))

	alternative, ok := exp.Alternative.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	testIdentifier(t, alternative.Expression, "y")
}

func TestFunctionLiteralParsing(t *testing.T) {
	input := `fn(x, y) { x + y; }`
	numStatements := 1
	program := setupTest(t, input)

	assert.Equal(t, numStatements, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	assert.True(t, ok)
	assert.Equal(t, 2, len(function.Parameters))

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	assert.Equal(t, 1, len(function.Body.Statements))

	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "fn() {};", expectedParams: []string{}},
		{input: "fn(x) {};", expectedParams: []string{"x"}},
		{input: "fn(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		program := setupTest(t, tt.input)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		function := stmt.Expression.(*ast.FunctionLiteral)

		assert.Equal(t, len(tt.expectedParams), len(function.Parameters))

		for i, ident := range tt.expectedParams {
			testLiteralExpression(t, function.Parameters[i], ident)
		}
	}
}

func TestCallExpressionParsing(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5);"
	program := setupTest(t, input)

	assert.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	exp, ok := stmt.Expression.(*ast.CallExpression)
	assert.True(t, ok)
	testIdentifier(t, exp.Function, "add")

	assert.Equal(t, 3, len(exp.Arguments))
	testLiteralExpression(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}

func TestArrayLiteral(t *testing.T) {
	input := "[1,2 * 3, 3 + 3]"
	program := setupTest(t, input)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	array, ok := stmt.Expression.(*ast.ArrayLiteral)
	assert.True(t, ok)
	assert.Equal(t, 3, len(array.Elements))
	testIntegerLiteral(t, array.Elements[0], 1)
	testInfixExpression(t, array.Elements[1], 2, "*", 3)
	testInfixExpression(t, array.Elements[2], 3, "+", 3)
}

func TestIndexExpressions(t *testing.T) {
	input := "myArray[1 + 1]"
	program := setupTest(t, input)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	indexExp, ok := stmt.Expression.(*ast.IndexExpression)
	assert.True(t, ok)
	testIdentifier(t, indexExp.Left, "myArray")
	testInfixExpression(t, indexExp.Index, 1, "+", 1)
}

// Helper methods

func setupTest(t *testing.T, input string) *ast.Program {
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p, false)
	return program
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) {
	switch v := expected.(type) {
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
