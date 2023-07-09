package parser

import (
	"fmt"
	"strconv"

	"github.com/sapslaj/pl-challenges/go/monkey/ast"
	"github.com/sapslaj/pl-challenges/go/monkey/lexer"
	"github.com/sapslaj/pl-challenges/go/monkey/token"
)

type ParseError error

func NewParseError(format string, a ...any) ParseError {
	err := fmt.Errorf(format, a...)
	return ParseError(err)
}

type PrefixParseFn func() ast.Expression
type InfixParseFn func(ast.Expression) ast.Expression

type Precedence int

const (
	_ Precedence = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var Precedences = map[token.TokenType]Precedence{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

type Parser struct {
	l              *lexer.Lexer
	curToken       token.Token
	peekToken      token.Token
	errors         []ParseError
	prefixParseFns map[token.TokenType]PrefixParseFn
	infixParseFns  map[token.TokenType]InfixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:              l,
		prefixParseFns: make(map[token.TokenType]PrefixParseFn),
		infixParseFns:  make(map[token.TokenType]InfixParseFn),
	}

	p.RegisterPrefix(token.IDENT, p.parseIdentifier)
	p.RegisterPrefix(token.INT, p.parseIntegerLiteral)
	p.RegisterPrefix(token.TRUE, p.parseBoolean)
	p.RegisterPrefix(token.FALSE, p.parseBoolean)
	p.RegisterPrefix(token.BANG, p.parsePrefixExpression)
	p.RegisterPrefix(token.MINUS, p.parsePrefixExpression)

	p.RegisterInfix(token.PLUS, p.parseInfixExpression)
	p.RegisterInfix(token.MINUS, p.parseInfixExpression)
	p.RegisterInfix(token.SLASH, p.parseInfixExpression)
	p.RegisterInfix(token.ASTERISK, p.parseInfixExpression)
	p.RegisterInfix(token.EQ, p.parseInfixExpression)
	p.RegisterInfix(token.NOT_EQ, p.parseInfixExpression)
	p.RegisterInfix(token.LT, p.parseInfixExpression)
	p.RegisterInfix(token.GT, p.parseInfixExpression)

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

func (p *Parser) RegisterPrefix(tokenType token.TokenType, fn PrefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) RegisterInfix(tokenType token.TokenType, fn InfixParseFn) {
	p.infixParseFns[tokenType] = fn
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

func (p *Parser) PeekPrecedence() Precedence {
	precedence, ok := Precedences[p.peekToken.Type]
	if ok {
		return precedence
	}
	return LOWEST
}

func (p *Parser) CurPrecendence() Precedence {
	precedence, ok := Precedences[p.curToken.Type]
	if ok {
		return precedence
	}
	return LOWEST
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

func (p *Parser) parseExpression(precedence Precedence) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.newParseError(
			"no prefix parse function for %s found", p.curToken.Type,
		)
		return nil
	}
	leftExp := prefix()
	for !p.PeekTokenIs(token.SEMICOLON) && precedence < p.PeekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.NextToken()

		leftExp = infix(leftExp)
	}
	return leftExp
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

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token: p.curToken,
	}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.PeekTokenIs(token.SEMICOLON) {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.NextToken()

	exp.Right = p.parseExpression(PREFIX)

	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.CurPrecendence()
	p.NextToken()
	exp.Right = p.parseExpression(precedence)
	return exp
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{
		Token: p.curToken,
	}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		p.newParseError("could not parse %q as integer: %v", p.curToken.Literal, err)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: p.curToken,
		Value: p.CurTokenIs(token.TRUE),
	}
}

func (p *Parser) ParseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
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
