package parser

import (
	"fmt"

	"github.com/sapslaj/pl-challenges/go/monkey/ast"
	"github.com/sapslaj/pl-challenges/go/monkey/lexer"
	"github.com/sapslaj/pl-challenges/go/monkey/token"
)

type ParseError error

func NewParseError(format string, a ...any) ParseError {
	err := fmt.Errorf(format, a...)
	return ParseError(err)
}

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []ParseError
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}
	p.NextToken()
	p.NextToken()
	return p
}

func (p *Parser) newParseError(format string, a ...any) {
	p.errors = append(p.errors, NewParseError(format, a...))
}

func (p *Parser) Errors() []ParseError {
	return p.errors
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) CurTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) PeekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) ExpectPeek(t token.TokenType) bool {
	if p.PeekTokenIs(t) {
		p.NextToken()
		return true
	}
	p.newParseError(
		"expected next token to be %s, got %s instead", t, p.peekToken.Type,
	)
	return false
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{
		Token: p.curToken,
	}

	if !p.ExpectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.ExpectPeek(token.ASSIGN) {
		return nil
	}

	for !p.CurTokenIs(token.SEMICOLON) {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{
		Token: p.curToken,
	}

	p.NextToken()

	for !p.CurTokenIs(token.SEMICOLON) {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) ParseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.ParseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.NextToken()
	}

	return program
}