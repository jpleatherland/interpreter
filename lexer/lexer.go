package lexer

import (
	"github.com/jpleatherland/interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	currChar     byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currChar = 0
	} else {
		l.currChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.currChar {
	case '=':
		tok = newToken(token.ASSIGN, l.currChar)
	case ';':
		tok = newToken(token.SEMICOLON, l.currChar)
	case '(':
		tok = newToken(token.LPAREN, l.currChar)
	case ')':
		tok = newToken(token.RPAREN, l.currChar)
	case ',':
		tok = newToken(token.COMMA, l.currChar)
	case '+':
		tok = newToken(token.PLUS, l.currChar)
	case '{':
		tok = newToken(token.LBRACE, l.currChar)
	case '}':
		tok = newToken(token.RBRACE, l.currChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.currChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.currChar){
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}else {
			tok = newToken(token.ILLEGAL, l.currChar)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, currChar byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(currChar)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.currChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace(){
	for l.currChar == ' ' || l.currChar == '\t' || l.currChar == '\n' || l.currChar == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.currChar){
		l.readChar()
	}
	return l.input[position: l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}



