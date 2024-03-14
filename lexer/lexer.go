package lexer

import (
	"monkey/token"
	"strings"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points at curr)
	readPosition int  // current reading position (points after curr)
	cur          byte // what we're currently looking at
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()

	return lexer
}

func (l *Lexer) peek() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// Advance the reader by one character
func (l *Lexer) readChar() {
	// If we're at the end of the road, set curr to a null byte for EOF
	if l.readPosition >= len(l.input) {
		l.cur = 0
	} else {
		l.cur = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for strings.IndexRune(" \t\n\r", rune(l.cur)) >= 0 {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.cur) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.cur) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) Next() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.cur {
	case ';':
		tok = token.New(token.SEMICOLON, l.cur)
	case '(':
		tok = token.New(token.LPAREN, l.cur)
	case ')':
		tok = token.New(token.RPAREN, l.cur)
	case '{':
		tok = token.New(token.LBRACE, l.cur)
	case '}':
		tok = token.New(token.RBRACE, l.cur)
	case ',':
		tok = token.New(token.COMMA, l.cur)
	case '=':
		if l.peek() == '=' {
			prev := l.cur
			l.readChar()
			tok.Type = token.EQ
			tok.Value = string(prev) + string(l.cur)
		} else {
			tok = token.New(token.ASSIGN, l.cur)
		}
	case '+':
		tok = token.New(token.PLUS, l.cur)
	case '-':
		tok = token.New(token.MINUS, l.cur)
	case '!':
		if l.peek() == '=' {
			prev := l.cur
			l.readChar()
			tok.Type = token.NOT_EQ
			tok.Value = string(prev) + string(l.cur)
		} else {
			tok = token.New(token.BANG, l.cur)
		}
	case '/':
		tok = token.New(token.SLASH, l.cur)
	case '*':
		tok = token.New(token.ASTERISK, l.cur)
	case '<':
		tok = token.New(token.LT, l.cur)
	case '>':
		tok = token.New(token.GT, l.cur)
	case 0:
		tok.Value = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.cur) {
			tok.Value = l.readIdentifier()
			tok.Type = token.Lookup(tok.Value)
			return tok
		} else if isDigit(l.cur) {
			tok.Value = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.New(token.ILLEGAL, l.cur)
		}
	}

	l.readChar()
	return tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
