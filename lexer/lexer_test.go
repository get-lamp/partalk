package lexer

import (
	"partalk/token"
	"testing"
)

func TestReadChar(t *testing.T) {
	line := "A line of text"
	lex := New(line)

	for c := 0; c < len(line); c++ {
		if lex.char != line[c] {
			t.Errorf("Expecting %s, read %s", string(line[c]), string(lex.char))
		}
		lex.readChar()
	}

	if lex.char != 0 {
		t.Fatalf("Unexpected final state. Lexer.char == %s. Expected 0.", string(lex.char))
	}
}

func TestNextToken(t *testing.T) {
	line := "A line of text"
	lex := New(line)

	tok := lex.NextToken()

	if tok.Type != token.IDENT || tok.Literal != "A" {
		t.Fatal()
	}
}
