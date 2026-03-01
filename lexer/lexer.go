package lexer

import (
	"github.com/DanieleRuffo93/go-interpreter/token"
	"unicode"
	"unicode/utf8"
)

// TODO: Could implement row column traking and specify where an error has been found
type Lexer struct {
	input        string
	position     int  // current position in input
	nextPosition int  // next position in input, used to peek next char
	ch           rune // Used Rune to support Unicode
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // calling this to set Lexer properly the first time
	return l
}

// Unicode support added
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

	l.skipWhiteSpaces()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readWhile(isLetter) // get the identifier
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok // return early here, cursor and next char has been read by readIdentifier()
		} else if isDigit(l.ch) {
			// TODO: Just integers for now. Could add float, hex, ecc handling. Waiting to implement evauluetor to see what needs to change
			tok.Type = token.INT
			tok.Literal = l.readWhile(isDigit) // get the number
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}
	l.readChar() // Move cursors and read next char
	return tok
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readWhile(condition func(rune) bool) string {
	position := l.position
	for condition(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func (l *Lexer) skipWhiteSpaces() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func (l *Lexer) peekChar() rune {
	if l.nextPosition >= len(l.input) {
		return 0
	}
	rune, _ := utf8.DecodeRuneInString(l.input[l.nextPosition:])
	return rune
}
