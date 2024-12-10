package parser

import (
	"partalk/lexer"
	"testing"
)

func TestParseDeclareStatement(t *testing.T) {
	l := lexer.New("something")
	p := New(l)
	program := p.ParseProgram()

	if program == nil {

	}
}

func TestParseSchema(t *testing.T) {
	l := lexer.New("something {foo, bar, baz, qux}")
	p := New(l)
	program := p.ParseProgram()

	if program == nil {

	}
}

func TestParseObject(t *testing.T) {
	l := lexer.New("something {foo: bar, baz: qux}")
	p := New(l)
	program := p.ParseProgram()

	if program == nil {

	}
}

func TestParseNestedObject(t *testing.T) {
	l := lexer.New("something {bar: {baz: qux}}")
	p := New(l)
	program := p.ParseProgram()

	if program == nil {

	}
}

func TestParseDotNotation(t *testing.T) {
	l := lexer.New("something.blop")
	p := New(l)
	program := p.ParseProgram()

	if program == nil {

	}
}
