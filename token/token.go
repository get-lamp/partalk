package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ          = "=="
	NOT_EQ      = "!="
	SHIFT_RIGHT = ">>"
	SHIFT_LEFT  = "<<"
	INC         = "++"
	DEC         = "--"
	RIGHT       = "->"
	LEFT        = "<-"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	DOT       = "."

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	COLON  = ":"
	QUOTE  = "\""

	// Keywords
	FUNCTION = "FUNCTION"
	QUERY    = "?"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func (t *Token) Foo() bool {
	return true
}
