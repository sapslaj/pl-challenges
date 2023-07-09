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
	assert.Len(t, errs, 4)
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

func TestValidIntegerLiteralExpression(t *testing.T) {
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

func TestBooleanExpressions(t *testing.T) {
	tests := []struct {
		input string
		lit   string
		value bool
	}{
		{"true;", "true", true},
		{"false;", "false", false},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)

		program := p.ParseProgram()
		assert.Len(t, p.Errors(), 0, "parser has %d errors: %v", len(p.Errors()), p.Errors())
		assert.NotNil(t, program)
		assert.Len(t, program.Statements, 1)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		boolean := stmt.Expression.(*ast.Boolean)
		assert.Equal(t, tt.lit, boolean.Token.Literal)
		assert.Equal(t, tt.value, boolean.Value)
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	tests := []struct {
		input    string
		operator string
		value    any
	}{
		{"!5", "!", int64(5)},
		{"-15;", "-", int64(15)},
		{"!true;", "!", true},
		{"!false;", "!", false},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)

		program := p.ParseProgram()
		assert.Len(t, p.Errors(), 0, "parser has %d errors: %v", len(p.Errors()), p.Errors())
		assert.NotNil(t, program)
		assert.Len(t, program.Statements, 1)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		exp := stmt.Expression.(*ast.PrefixExpression)
		assert.Equal(t, tt.operator, exp.Operator)
		switch tt.value.(type) {
		case int64:
			right := exp.Right.(*ast.IntegerLiteral)
			assert.Equal(t, tt.value, right.Value)
		case bool:
			right := exp.Right.(*ast.Boolean)
			assert.Equal(t, tt.value, right.Value)
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	tests := []struct {
		input      string
		leftValue  any
		operator   string
		rightValue any
	}{
		{"5 + 5;", int64(5), "+", int64(5)},
		{"5 - 5;", int64(5), "-", int64(5)},
		{"5 * 5;", int64(5), "*", int64(5)},
		{"5 / 5;", int64(5), "/", int64(5)},
		{"5 > 5;", int64(5), ">", int64(5)},
		{"5 < 5;", int64(5), "<", int64(5)},
		{"5 == 5;", int64(5), "==", int64(5)},
		{"5 != 5;", int64(5), "!=", int64(5)},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)

		program := p.ParseProgram()
		assert.Len(t, p.Errors(), 0, "parser has %d errors: %v", len(p.Errors()), p.Errors())
		assert.NotNil(t, program)
		assert.Len(t, program.Statements, 1)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		exp := stmt.Expression.(*ast.InfixExpression)
		assert.Equal(t, tt.operator, exp.Operator)

		switch tt.leftValue.(type) {
		case int:
			left := exp.Left.(*ast.IntegerLiteral)
			assert.Equal(t, tt.leftValue, left.Value)
			right := exp.Right.(*ast.IntegerLiteral)
			assert.Equal(t, tt.rightValue, right.Value)
		case bool:
			left := exp.Left.(*ast.Boolean)
			assert.Equal(t, tt.leftValue, left.Value)
			right := exp.Right.(*ast.Boolean)
			assert.Equal(t, tt.rightValue, right.Value)
		}
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
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)

		program := p.ParseProgram()
		assert.Len(t, p.Errors(), 0, "parser has %d errors: %v", len(p.Errors()), p.Errors())
		assert.NotNil(t, program)
		assert.Equal(t, tt.expected, program.String())
	}
}
