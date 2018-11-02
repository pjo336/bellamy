package parser

import (
	"Bellamy/ast"
	"Bellamy/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 1234;
	`
	numStatements := 3
	l:= lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p, false)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != numStatements {
		t.Fatalf("program.Statements does not contain %d statements. got %d statements", numStatements, len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got %q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("letStmt.Name.Value not '%s'. got %s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got %s", name, letStmt.Name)
		return false
	}

	return true
}

func TestParserErrors(t *testing.T) {
	input := `
	let x 5;
	`
	l:= lexer.New(input)
	p := New(l)

	p.ParseProgram()
	// we expect errors here
	es := checkParserErrors(t, p, true)
	if es[0] != "expected next token to be =, got INT" {
		t.Fatalf("Expected specific error message but got %q", es[0])
	}
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

	if len(program.Statements) != numStatements {
		t.Fatalf("program.statements does not contain %d statements, got %d", numStatements, len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement is not ReturnStatement, got %T", stmt)
			continue // why continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt token literal not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
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