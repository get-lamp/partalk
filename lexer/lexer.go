package lexer

import "partalk/token"

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skip(n int) {
	l.position += n
	l.readPosition = l.position + 1
	l.readChar()
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token
	l.skipWhitespace()

	switch l.char {
	case 0:
		return token.Token{Type: token.EOF, Literal: ""}
	case '=':
		t = l.parseEqualSymbol()
	case '!':
		t = l.parseBangSymbol()
	case '>':
		t = l.parseGreaterSymbol()
	case '<':
		t = l.parseLesserSymbol()
	case '+':
		t = l.parsePlusSymbol()
	case '-':
		t = l.parseMinusSymbol()
	case '(':
		t = token.Token{Type: token.LPAREN, Literal: string(l.char)}
	case ')':
		t = token.Token{Type: token.RPAREN, Literal: string(l.char)}
	default:
		if isLetter(l.char) {
			return token.Token{Type: token.IDENT, Literal: l.readIdentifier()}
		} else if isDigit(l.char) {
			return token.Token{Type: token.INT, Literal: l.readNumber()}
		} else {
			l.skip(1)
			return token.Token{Type: token.ILLEGAL, Literal: string(l.char)}
		}
	}

	l.readChar()
	return t
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) parseEqualSymbol() token.Token {
	nextChar := string(l.peekChar())

	switch nextChar {
	case token.ASSIGN:
		l.skip(1)
		return token.Token{Type: token.EQ, Literal: token.ASSIGN + string(nextChar)}
	default:
		return token.Token{Type: token.ASSIGN, Literal: string(l.char)}
	}
}

func (l *Lexer) parseBangSymbol() token.Token {
	nextChar := string(l.peekChar())

	switch nextChar {
	case token.EQ:
		l.skip(1)
		return token.Token{Type: token.NOT_EQ, Literal: token.BANG + string(nextChar)}
	default:
		return token.Token{Type: token.BANG, Literal: string(l.char)}
	}
}

func (l *Lexer) parseGreaterSymbol() token.Token {
	nextChar := string(l.peekChar())

	switch nextChar {
	case token.GT:
		l.skip(1)
		return token.Token{Type: token.SHIFT_RIGHT, Literal: token.GT + string(nextChar)}
	default:
		return token.Token{Type: token.GT, Literal: string(l.char)}
	}
}

func (l *Lexer) parseLesserSymbol() token.Token {
	nextChar := string(l.peekChar())

	switch nextChar {
	case token.LT:
		l.skip(1)
		return token.Token{Type: token.SHIFT_LEFT, Literal: token.LT + string(nextChar)}
	default:
		return token.Token{Type: token.LT, Literal: string(l.char)}
	}
}

func (l *Lexer) parsePlusSymbol() token.Token {
	nextChar := string(l.peekChar())

	switch nextChar {
	case token.PLUS:
		l.skip(1)
		return token.Token{Type: token.INC, Literal: token.PLUS + string(nextChar)}
	default:
		return token.Token{Type: token.PLUS, Literal: string(l.char)}
	}
}

func (l *Lexer) parseMinusSymbol() token.Token {
	nextChar := string(l.peekChar())

	switch nextChar {
	case token.MINUS:
		l.skip(1)
		return token.Token{Type: token.DEC, Literal: token.MINUS + string(nextChar)}
	default:
		return token.Token{Type: token.MINUS, Literal: string(l.char)}
	}
}
