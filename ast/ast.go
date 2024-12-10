package ast

import "partalk/token"

type Attribute struct {
	Name  string
	Value string
}

type Object struct {
	Attributes []Attribute
}

func (o *Object) TokenLiteral() string {
	return ""
}

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
}

type Expression interface {
	Node
}

type Identifier struct {
	Token  token.Token
	Parent *Identifier
	Child  *Identifier
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type DeclareStatement struct {
	Node
	Statement
	Name  string
	Token token.Token
}

func (ss *DeclareStatement) statementNode() {}

func (ss *DeclareStatement) TokenLiteral() string {
	return ss.Token.Literal
}

type AssignStatement struct {
	Left  Identifier
	Right Object
}

func (as *AssignStatement) statementNode() {}

func (as *AssignStatement) TokenLiteral() string {
	return as.Left.Token.Literal
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
