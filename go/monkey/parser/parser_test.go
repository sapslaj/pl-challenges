package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sapslaj/pl-challenges/go/monkey/ast"
	"github.com/sapslaj/pl-challenges/go/monkey/lexer"
	"github.com/sapslaj/pl-challenges/go/monkey/parser"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	assert.NotNil(t, program)
	assert.Len(t, program.Statements, 3)

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		assert.Equal(t, "let", stmt.TokenLiteral())
		letStmt := stmt.(*ast.LetStatement)
		assert.Equal(t, tt.expectedIdentifier, letStmt.Name.Value)
		assert.Equal(t, tt.expectedIdentifier, letStmt.Name.TokenLiteral())
	}
}
