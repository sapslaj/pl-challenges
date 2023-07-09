package lexer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sapslaj/pl-challenges/go/monkey/lexer"
	"github.com/sapslaj/pl-challenges/go/monkey/token"
)

func TestNextToken(t *testing.T) {
	t.Parallel()

	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tok.Type, tt.expectedType)
		assert.Equal(t, tok.Literal, tt.expectedLiteral)
	}
}
