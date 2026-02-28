package lexer

import (
	"github.com/DanieleRuffo93/go-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `+=(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, test := range tests {
		tok := l.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("test[%d] - Wrong Type\nExpected: %q\nGot:%q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("test[%d] - Wrong Literal\nExpected: %q\nGot:%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}
