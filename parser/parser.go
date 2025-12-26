package parser

import (
	"github.com/self-sasi/monkey-interpreter/ast"
	"github.com/self-sasi/monkey-interpreter/lexer"
	"github.com/self-sasi/monkey-interpreter/token"
)

// Parser transforms a stream of tokens produced by the lexer
// into an Abstract Syntax Tree (AST).
type Parser struct {
	lex       *lexer.Lexer // source of tokens
	curToken  token.Token  // current token under examination
	peekToken token.Token  // next token (one-token lookahead)
}

// Creates and initializes a new Parser.
func New(lex *lexer.Lexer) *Parser {
	parserPointer := &Parser{lex: lex}

	// read two tokens, so curToken and peekToken are both set
	parserPointer.nextToken()
	parserPointer.nextToken()

	return parserPointer
}

// Advances the parser to the next token.
// The current token becomes the previous peek token,
// and a new peek token is read from the lexer.
func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lex.NextToken()
}

func (parser *Parser) ParseProgram() *ast.Program {
	return nil
}
