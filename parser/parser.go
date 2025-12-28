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
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.curToken.Type != token.EOF {
		statement := parser.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		parser.nextToken()
	}

	return program
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.curToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	default:
		return nil
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	letStatement := &ast.LetStatement{Token: parser.curToken}

	if !parser.expectPeek(token.IDENT) {
		return nil
	}

	letStatement.Name = &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}

	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: skipping the expressions until encounter a semicolon
	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return letStatement
}

func (parser *Parser) curTokenIs(tok token.TokenType) bool {
	return parser.curToken.Type == tok
}

func (parser *Parser) peekTokenIs(tok token.TokenType) bool {
	return parser.peekToken.Type == tok
}

func (parser *Parser) expectPeek(tok token.TokenType) bool {
	if parser.peekTokenIs(tok) {
		parser.nextToken()
		return true
	} else {
		return false
	}
}
