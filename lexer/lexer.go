package lexer

import (
	"github.com/DanieleRuffo93/go-interpreter/token"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int // current position in input
	nextPosition int // next position in input, used to peek next char
	ch           rune
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Unicode support added -- this is different from the book version
func (l *Lexer) readChar() {
	size := 1
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		r, u_size := utf8.DecodeRuneInString(l.input[l.nextPosition:])
		l.ch = r
		size = u_size
	}

	l.position = l.nextPosition
	l.nextPosition += size
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
