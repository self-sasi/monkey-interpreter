package lexer

import "github.com/self-sasi/monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to the current char)
	readPosition int  // current reading position in input (after current char)
	char         byte // current char under examination
}

func New(input string) *Lexer {
	newLexer := &Lexer{input: input}
	newLexer.readChar()
	return newLexer
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.char = 0 // ascii code for "NUL" character
	} else {
		lex.char = lex.input[lex.readPosition]
	}
	lex.position = lex.readPosition
	lex.readPosition += 1
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	lex.eatWhitespace()

	switch lex.char {
	case '=':
		tok = newToken(token.ASSIGN, lex.char)
	case ';':
		tok = newToken(token.SEMICOLON, lex.char)
	case '(':
		tok = newToken(token.LPAREN, lex.char)
	case ')':
		tok = newToken(token.RPAREN, lex.char)
	case ',':
		tok = newToken(token.COMMA, lex.char)
	case '+':
		tok = newToken(token.PLUS, lex.char)
	case '{':
		tok = newToken(token.LBRACE, lex.char)
	case '}':
		tok = newToken(token.RBRACE, lex.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lex.char) {
			tok.Literal = lex.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(lex.char) {
			tok.Type = token.INT
			tok.Literal = lex.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lex.char)
		}
	}

	lex.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (lex *Lexer) readIdentifier() string {
	position := lex.position
	for isLetter(lex.char) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lex *Lexer) readNumber() string {
	position := lex.position
	for isDigit(lex.char) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

func (lex *Lexer) eatWhitespace() {
	for lex.char == ' ' || lex.char == '\t' || lex.char == '\n' || lex.char == '\r' {
		lex.readChar()
	}
}
