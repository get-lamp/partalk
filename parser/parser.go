package parser

import (
	"partalk/ast"
	"partalk/lexer"
	"partalk/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.lexer.NextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.IDENT:
		return p.parseIdentifier()
	default:
		panic("Unexpected token")
	}
}

func (p *Parser) parseIdentifier() ast.Statement {

	switch p.peekToken.Type {
	case token.EOF:
		return &ast.SetStatement{Token: p.currToken, Name: p.currToken.Literal, Value: nil}
	case token.LBRACE:
		return &ast.Identifier{p.currToken, p.parseEntity()}
	default:
		panic("Unexpected token while parsing identifier")
	}
}

func (p *Parser) parseEntity() ast.Entity {

	entity := ast.Entity{}

	if p.peekToken.Type != token.IDENT {
		panic("Expected identifier. Got something else")
	}

	p.lexer.NextToken()

	for p.currToken.Type != token.RBRACE {

		if p.currToken.Type == token.EOF {
			panic("Unexpected EOF")
		}

		property := ast.Property{Name: p.currToken.Literal, Value: p.peekToken.Literal}
		entity.Properties = append(entity.Properties, property)

		p.lexer.NextToken()
	}

	return entity
}
