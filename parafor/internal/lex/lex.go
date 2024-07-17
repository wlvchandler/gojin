package lex

import (
	"localhost/wlvchandler/parafor/internal/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	c_char       byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.c_char = 0
	} else {
		l.c_char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var c_token token.Token

	l.skipWs()

	switch l.c_char {
	case ':':
		c_token = newToken(token.COLON, l.c_char)
	case 0:
		c_token.Literal = ""
		c_token.Type = token.EOF
	default:
		if isLetter(l.c_char) {
			c_token.Literal = l.readIdentifier()
			c_token.Type = token.IdentifierLookup(c_token.Literal)
			return c_token
		} else if isDigit(l.c_char) {
			c_token.Type = token.INT
			c_token.Literal = l.readNumber()
			return c_token
		} else {
			c_token = newToken(token.ILLEGAL, l.c_char)
		}
	}

	l.readChar()
	return c_token
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.c_char) || isDigit(l.c_char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.c_char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func newToken(ttype token.TokenType, c_char byte) token.Token {
	return token.Token{Type: ttype, Literal: string(c_char)}
}

func (l *Lexer) skipWs() {
	for {
		switch l.c_char {
		case ' ', '\t', '\n', '\r':
			l.readChar()
		default:
			return
		}
	}
}

func isLetter(c_char byte) bool {
	return 'a' <= c_char && c_char <= 'z' || 'A' <= c_char && c_char <= 'Z' || c_char == '_'
}

func isDigit(c_char byte) bool {
	return '0' <= c_char && c_char <= '9'
}
