package lexer

import (
	"fmt"

	"github.com/iashyam/bclang/token"
)

type Lexer struct {
	input      string
	readPos    int
	currentPos int
	ch         byte
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func New(inp string) *Lexer {
	l := &Lexer{input: inp}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.currentPos = l.readPos
	l.readPos++
}

func (l *Lexer) readIdentifier() string {
	start := l.currentPos
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.currentPos]
}

func (l *Lexer) readNumber() string {
	start := l.currentPos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.currentPos]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			idf := l.readIdentifier()
			fmt.Print(idf)
			tok.Type = token.LookUpIdent(idf)
			tok.Literal = idf
			return tok
		} else if isDigit(l.ch) {
			idf := l.readNumber()
			fmt.Print(idf)
			tok.Type = token.INT
			tok.Literal = idf
			return tok
		} else {
			tok.Literal = ""
			tok.Type = token.ILLEGAL
		}
	}
	l.readChar()
	return tok
}
