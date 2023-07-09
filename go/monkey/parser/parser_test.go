package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sapslaj/pl-challenges/go/monkey/ast"
	"github.com/sapslaj/pl-challenges/go/monkey/lexer"
	"github.com/sapslaj/pl-challenges/go/monkey/parser"
)

func TestValidLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	assert.Len(t, p.Errors(), 0, "parser has %d errors: %v", len(p.Errors()), p.Errors())
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

func TestInvalidLetStatements(t *testing.T) {
	input := `
let x 5;
let = 10;
let 838383;
`

	l := lexer.New(input)
	p := parser.New(l)

	p.ParseProgram()
	errs := p.Errors()
	assert.Len(t, errs, 3)
	for _, err := range errs {
		assert.ErrorContains(t, err, "expected next token to be")
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	assert.Len(t, p.Errors(), 0, "parser has %d errors: %v", len(p.Errors()), p.Errors())
	assert.NotNil(t, program)
	assert.Len(t, program.Statements, 3)

	for _, stmt := range program.Statements {
		returnStmt := stmt.(*ast.ReturnStatement)
		assert.Equal(t, "return", returnStmt.TokenLiteral())
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	assert.Len(t, p.Errors(), 0, "parser has %d errors: %v", len(p.Errors()), p.Errors())
	assert.NotNil(t, program)
	assert.Len(t, program.Statements, 1)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	ident := stmt.Expression.(*ast.Identifier)
	assert.Equal(t, "foobar", ident.TokenLiteral())
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	assert.Len(t, p.Errors(), 0, "parser has %d errors: %v", len(p.Errors()), p.Errors())
	assert.NotNil(t, program)
	assert.Len(t, program.Statements, 1)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	literal := stmt.Expression.(*ast.IntegerLiteral)
	assert.Equal(t, "5", literal.TokenLiteral())
	assert.Equal(t, int64(5), literal.Value)
}
