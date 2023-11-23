package parser

import (
	"partalk/ast"
	"partalk/lexer"
	"partalk/token"
	"slices"
)

type Parser struct {
	lexer     *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.IDENT:
		return p.parseIdentifier()
	case token.QUERY:
		return p.parseQuery()
	default:
		return nil
	}
}

func (p *Parser) parseQueryStatement() ast.QueryStatement {
	return ast.QueryStatement{}
}

func (p *Parser) parseQuery() ast.Statement {
	if !p.currTokenIs(token.QUERY) {
		panic("Expected ? symbol.")
	}

	var query ast.QueryStatement

	switch p.peekToken.Type {
	case token.EOF:
		query = ast.QueryStatement{}
	case token.IDENT:
		query = p.parseQueryStatement()
	default:
		panic("Unexpected symbol while parsing query statement")
	}

	return query
}

func (p *Parser) parseIdentifier() ast.Statement {

	identifier := p.currToken

	switch p.peekToken.Type {
	case token.EOF:
		return &ast.SetStatement{Token: p.currToken, Name: p.currToken.Literal, Value: nil}
	case token.LBRACE:
		p.nextToken()
		return &ast.Identifier{identifier, p.parseObject()}
	default:
		panic("Unexpected token while parsing identifier")
	}
}

func (p *Parser) parseQuote(delimiters []token.TokenType) string {
	var quote = ""

	for !slices.Contains(delimiters, p.currToken.Type) {
		if p.currToken.Type == token.EOF {
			panic("Unexpected end of file")
		}
		quote += p.currToken.Literal
		p.nextToken()
	}
	return quote
}

func (p *Parser) parseProperty() ast.Attribute {

	if !p.currTokenIs(token.IDENT) {
		panic("Expected identifier while parsing property")
	}

	name := p.currToken.Literal

	if !p.expectPeek(token.COLON) {

	} else {
		p.nextToken()
	}

	switch p.currToken.Type {
	case token.QUOTE:
		p.nextToken()
		return ast.Attribute{Name: name, Value: p.parseQuote([]token.TokenType{token.QUOTE})}
	case token.LBRACE:
		panic("Not implemented")
		//return ast.Attribute{Name: name, Value: p.parseObject()}
	default:
		return ast.Attribute{Name: name, Value: p.parseQuote([]token.TokenType{token.COMMA, token.RBRACE})}
	}
}

func (p *Parser) parseObject() ast.Object {

	entity := ast.Object{}

	if !p.currTokenIs(token.LBRACE) {
		panic("Expected left brace. Got something else")
	}

	p.nextToken()

	for !slices.Contains([]token.TokenType{token.RBRACE, token.EOF}, p.currToken.Type) {

		switch p.currToken.Type {
		case token.EOF:
			panic("Unexpected EOF")
		case token.IDENT:
			entity.Attributes = append(entity.Attributes, p.parseProperty())
		default:
			panic("Cannot handle input")
		}

		p.nextToken()
	}

	return entity
}
