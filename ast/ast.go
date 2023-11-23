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

type Attribute struct {
	Name  string
	Value string
}

type Object struct {
	Attributes []Attribute
}

type Identifier struct {
	Token token.Token
	Value Object
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type SetStatement struct {
	Node
	Statement
	Name  string
	Token token.Token
	Value Expression
}

func (ss *SetStatement) statementNode() {}

func (ss *SetStatement) TokenLiteral() string {
	return ss.Token.Literal
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

type QueryStatement struct {
	Node
	Statement
	Object Object
}
