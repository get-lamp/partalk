package ast

import "partalk/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
}

type Expression interface {
	Node
	StatementNode()
}

type Property struct {
	Name  string
	Value string
}

type Entity struct {
	Properties []Property
}

type Identifier struct {
	Token token.Token
	Value Entity
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type SetStatement struct {
	Node
	Statement
	Token token.Token
	Name  string
	Value Expression
}

func (ss *SetStatement) statementNode() {}

func (ss *SetStatement) TokenLiteral() string {
	return ss.Token.Literal
}

type Element struct {
	identifier string
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
