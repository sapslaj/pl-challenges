package ast

import (
	"strings"

	"github.com/sapslaj/pl-challenges/go/monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var b strings.Builder
	for _, s := range p.Statements {
		b.WriteString(s.String())
	}
	return b.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) ExpressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) StatementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var b strings.Builder
	b.WriteString(ls.TokenLiteral())
	b.WriteString(" ")
	b.WriteString(ls.Name.String())
	b.WriteString(" = ")
	if ls.Value != nil {
		b.WriteString(ls.Value.String())
	}
	b.WriteString(";")
	return b.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) StatementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var b strings.Builder
	b.WriteString(rs.TokenLiteral())
	b.WriteString(" ")
	if rs.ReturnValue != nil {
		b.WriteString(rs.ReturnValue.String())
	}
	b.WriteString(";")
	return b.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) StatementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) ExpressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var b strings.Builder
	b.WriteString("(")
	b.WriteString(pe.Operator)
	b.WriteString(pe.Right.String())
	b.WriteString(")")
	return b.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) ExpressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var b strings.Builder
	b.WriteString("(")
	b.WriteString(ie.Left.String())
	b.WriteString(" ")
	b.WriteString(ie.Operator)
	b.WriteString(" ")
	b.WriteString(ie.Right.String())
	b.WriteString(")")
	return b.String()
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) ExpressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) ExpressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }
