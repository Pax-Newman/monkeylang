package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.Parse()
	if program == nil {
		t.Fatalf("Parse() returned nil program")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got `%d`", len(program.Statements))
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

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("Expected Statement==`*ast.LetStatement` Got `%T`", stmt)
		return false
	}

	if letStmt.TokenLiteral() != "let" {
		t.Errorf("Expected Statement.TokenLiteral==`let` Got `%q`", stmt.TokenLiteral())
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("Expected Statement.Name.Value==`%s` Got `%s`", name, letStmt.Name.Value)
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("Expected Statement.Name==`%s` Got `%s`", name, letStmt.Name)
	}

	return true
}
