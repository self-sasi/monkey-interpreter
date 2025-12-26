package lexer

import "github.com/self-sasi/monkey-interpreter/token"

// Represents a lexer that maintains state for lexical analysis.
type Lexer struct {
	input        string
	position     int  // current position in input (points to the current char)
	readPosition int  // current reading position in input (after current char)
	char         byte // current char under examination
}

// Creates a new [Lexer] and returns the pointer to the struct.
func New(input string) *Lexer {
	newLexer := &Lexer{input: input}
	newLexer.readChar()
	return newLexer
}

// Makes the [Lexer] read a char, i.e, increment [Lexer.position] to the next
// index and store the corresponding character in [Lexer.char].
//
// For example:
//
//	someLexer := Lexer{
//			input: "xyz"
//			position: 0
//			readPosition: 1
//			char: "x"
//	}
//
//	someLexer.readChar() // position == 1, char == 'y'
func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.char = 0 // ascii code for "NUL" character
	} else {
		lex.char = lex.input[lex.readPosition]
	}
	lex.position = lex.readPosition
	lex.readPosition += 1
}

// Makes the [Lexer] peek the next char, i.e, return the char that exists at
// [Lexer.char]'s [Lexer.readPosition] index. This function comes in handy
// when checking whether a token is single-character length, or double. For
// example, differentiating between "=" and "==".
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// Makes the [Lexer] identify the next token, i.e, read the next char/chars &
// classify it/them into a [token.Token] by identifying [token.TokenType].
// Returns the [token.Token] which stores the literal value and token type.
func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	lex.eatWhitespace() // clear any white space

	switch lex.char {
	case '=':
		if lex.peekChar() == '=' {
			prevChar := lex.char
			lex.readChar() // increment lex.position as it is a dual-char token
			tok = newTwoCharToken(token.EQ, prevChar, lex.char)
		} else {
			tok = newToken(token.ASSIGN, lex.char)
		}
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
	case '!':
		if lex.peekChar() == '=' {
			prevChar := lex.char
			lex.readChar() // increment lex.position as it is a dual-char token
			tok = newTwoCharToken(token.NOT_EQ, prevChar, lex.char)
		} else {
			tok = newToken(token.BANG, lex.char)
		}
	case '-':
		tok = newToken(token.MINUS, lex.char)
	case '*':
		tok = newToken(token.ASTERISK, lex.char)
	case '/':
		tok = newToken(token.SLASH, lex.char)
	case '>':
		tok = newToken(token.GT, lex.char)
	case '<':
		tok = newToken(token.LT, lex.char)
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

// Consumes consecutive letter characters starting at the current
// [Lexer.position] and returns the corresponding identifier literal.
func (lex *Lexer) readIdentifier() string {
	position := lex.position
	for isLetter(lex.char) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

// Consumes consecutive digit characters starting at the current
// [Lexer.position] and returns the corresponding numeric literal.
func (lex *Lexer) readNumber() string {
	position := lex.position
	for isDigit(lex.char) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

// Helper that creates a new token given the tokenType and ch (char).
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Helper that creates a new token with a two-char-length Literal value given the tokenType
// and first and second characters.
func newTwoCharToken(tokenType token.TokenType, firstCh byte, secondCh byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(firstCh) + string(secondCh)}
}

// Helper that determines if the given char is a letter.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Helper that determines if the given char is a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Helper that eats whitespace.
func (lex *Lexer) eatWhitespace() {
	for lex.char == ' ' || lex.char == '\t' || lex.char == '\n' || lex.char == '\r' {
		lex.readChar()
	}
}
